package main

import (
	"fretvis/music"
	"math"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type Fretboard struct {
	widgets.QWidget
	margin    int
	nutM      int
	nStrings  int
	nFrets    int
	strMargin int
	fretBLen  int
	notePen   *gui.QPen
	noteBrush *gui.QBrush
	noteFont  *gui.QFont
}

func NewFretboard() *Fretboard {
	f := &Fretboard{
		QWidget: *widgets.NewQWidget(nil, 0),
	}
	f.margin = 12
	f.nutM = 8
	f.nStrings = 6
	f.nFrets = 24
	f.strMargin = 25
	f.fretBLen = 650
	contentW := f.fretBLen + f.nutM
	contentH := (f.nStrings - 1) * f.strMargin
	f.SetFixedSize2(contentW+2*f.margin, contentH+2*f.margin)
	f.ConnectPaintEvent(f.draw)

	f.notePen = gui.NewQPen()
	f.notePen.SetWidth(2)
	f.noteBrush = gui.NewQBrush3(gui.NewQColor2(core.Qt__white), core.Qt__SolidPattern)
	f.noteFont = gui.NewQFont()
	f.noteFont.SetStyleHint(gui.QFont__Monospace, gui.QFont__StyleStrategy(gui.QFont__AbsoluteSpacing))
	f.noteFont.SetBold(true)
	f.noteFont.SetPointSize(10)
	return f
}

func (f *Fretboard) draw(paintEvent *gui.QPaintEvent) {
	painter := gui.NewQPainter2(f)
	painter.SetRenderHint(gui.QPainter__Antialiasing, false)
	painter.Pen().SetCapStyle(core.Qt__FlatCap)
	painter.Pen().SetJoinStyle(core.Qt__MiterJoin)

	w, h := f.Width(), f.Height()
	m := f.margin
	nutM := f.nutM
	nutStrokeW := 4
	strStrokeW := 2

	// Draw nut
	painter.Pen().SetWidth(nutStrokeW)
	painter.DrawRect2(m, m, nutM, h-2*m)
	// Bridge
	painter.DrawLine3(w-m, m, w-m, h-m)

	// Draw strings
	painter.Pen().SetWidth(strStrokeW)
	painter.Pen().SetCapStyle(core.Qt__SquareCap)
	for i := 0; i < f.nStrings; i++ {
		firstLine := 0
		if i == 0 {
			firstLine = nutStrokeW/2 - strStrokeW/2
		}
		lastLine := 0
		if i == f.nStrings-1 {
			lastLine = nutStrokeW/2 - strStrokeW/2
		}
		painter.DrawLine3(m+nutM, m+i*f.strMargin-firstLine+lastLine, w-m+1, m+i*f.strMargin-firstLine+lastLine)
	}

	// Draw frets
	for i := 0; i <= f.nFrets; i++ {
		fretDist := int(f.strDistOfFret(float64(i), f.fretBLen))
		painter.DrawLine3(m+nutM+fretDist, m, m+nutM+fretDist, h-m)
	}

	// Draw fret dots
	dotSize := 4
	painter.SetRenderHint(gui.QPainter__Antialiasing, true)
	painter.Brush().SetStyle(core.Qt__SolidPattern)
	fretDots := []float32{3, 5, 7, 9, 12, 15, 17, 19, 21, 24}
	for _, fretDot := range fretDots {
		fretDotOffs := (f.strDistOfFret(float64(fretDot-1.0), f.fretBLen-1) + f.strDistOfFret(float64(fretDot), f.fretBLen-1)) / 2.0
		if int(fretDot)%12 == 0 {
			painter.DrawEllipse3(m+nutM+int(fretDotOffs)-dotSize/2+1, m+int(1.5*float32(f.strMargin))-dotSize/2, dotSize, dotSize)
			painter.DrawEllipse3(m+nutM+int(fretDotOffs)-dotSize/2+1, m+int(3.5*float32(f.strMargin))-dotSize/2, dotSize, dotSize)
		} else {
			painter.DrawEllipse3(m+nutM+int(fretDotOffs)-dotSize/2+1, m+int(2.5*float32(f.strMargin))-dotSize/2, dotSize, dotSize)
		}
	}

	// Switch to settings for note drawing
	painter.SetPen(f.notePen)
	painter.SetBrush(f.noteBrush)
	painter.SetFont(f.noteFont)

	// Draw scale
	scale := music.GetScaleDegrees(music.Notes[SelectedKey], SelectedMode)
	for degree := 0; degree < 7; degree++ {
		if degree+1 == 1 {
			painter.Pen().SetColor(gui.NewQColor2(core.Qt__blue))
		} else if degree+1 == 3 {
			painter.Pen().SetColor(gui.NewQColor2(core.Qt__darkRed))
		} else if degree+1 == 5 {
			painter.Pen().SetColor(gui.NewQColor2(core.Qt__darkCyan))
		} else {
			painter.Pen().SetColor(gui.NewQColor2(core.Qt__gray))
		}

		f.drawNote(music.Notes[scale[degree]], painter)
	}

	painter.End()
}

// Calculate the distance up the fretboard a given fret is
func (f *Fretboard) strDistOfFret(fret float64, width int) float64 {
	fretsPerOctave := 20.0 // Should be 12 on a real guitar but alter slightly to make fret spacing more even
	maxFret := float64(f.nFrets)
	strLen := float64(width) / (1.0 - math.Pow(2.0, -maxFret/fretsPerOctave))
	return strLen - strLen/math.Pow(2.0, (fret/fretsPerOctave))
}

func (f *Fretboard) drawNote(note string, painter *gui.QPainter) {
	for str := 0; str < 6; str++ {
		noteInt, _ := music.GetNoteInt(note)
		for _, fret := range music.GetFretsOfNote(noteInt, str, 24) {
			f.drawNoteOnFret(note, str, fret, painter)
		}
	}
}

func (f *Fretboard) drawNoteOnFret(note string, str int, fret int, painter *gui.QPainter) {
	x := f.margin + f.nutM + int((f.strDistOfFret(float64(fret-1), f.fretBLen)+f.strDistOfFret(float64(fret), f.fretBLen))/2)
	y := f.margin + str*f.strMargin
	noteSize := 15

	if fret == 0 {
		x = f.margin + f.nutM/2
	}

	painter.DrawEllipse3(x-noteSize/2, y-noteSize/2, noteSize, noteSize)
	textRect := core.NewQRect4(x-noteSize/2-1, y-noteSize/2-1, noteSize+2, noteSize+4)
	painter.DrawText4(textRect, int(core.Qt__AlignCenter), string(note[0]), nil)
	if len(note) > 1 {
		prevSize := painter.Font().PointSize()
		painter.Font().SetPointSize(7)
		painter.DrawText4(textRect, int(core.Qt__AlignRight|core.Qt__AlignTop), string("#"), nil)
		painter.Font().SetPointSize(prevSize)
	}
}
