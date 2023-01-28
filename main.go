package main

import (
	"fretvis/music"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

var CurrentScale music.Scale

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Fretboard")

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	topBar := widgets.NewQWidget(nil, 0)
	topBar.SetLayout(widgets.NewQHBoxLayout())
	widget.Layout().AddWidget(topBar)

	labelKey := widgets.NewQLabel2("Key", nil, 0)
	topBar.Layout().AddWidget(labelKey)
	comboKey := widgets.NewQComboBox(nil)
	comboKey.AddItems([]string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"})
	topBar.Layout().AddWidget(comboKey)

	labelMode := widgets.NewQLabel2("Mode", nil, 0)
	topBar.Layout().AddWidget(labelMode)
	comboMode := widgets.NewQComboBox(nil)
	comboMode.AddItems([]string{"Ionian (Major)", "Dorian", "Phrygian", "Lydian", "Mixolydian", "Aeolian (Minor)", "Locrian"})
	topBar.Layout().AddWidget(comboMode)

	// CurrentScale := music.NewScale("C", 0)

	fretboard := NewFretboard()
	widget.Layout().AddWidget(fretboard)

	sCenter := app.QApplication_PTR().PrimaryScreen().AvailableGeometry().Center()
	wCenter := window.Geometry().Center()
	window.Move2(sCenter.X()-wCenter.X()/2, sCenter.Y()-wCenter.Y()/2)
	window.Show()
	app.Exec()
}

type Fretboard struct {
	widgets.QWidget
	margin    int
	nStrings  int
	strMargin int
}

func NewFretboard() *Fretboard {
	f := &Fretboard{
		QWidget: *widgets.NewQWidget(nil, 0),
	}
	f.margin = 10
	f.nStrings = 6
	f.strMargin = 20
	contentW := 360
	contentH := (f.nStrings - 1) * f.strMargin
	f.SetFixedSize2(contentW+2*f.margin, contentH+2*f.margin)
	f.ConnectPaintEvent(f.draw)
	return f
}

func (f *Fretboard) draw(paintEvent *gui.QPaintEvent) {
	painter := gui.NewQPainter2(f)
	painter.SetRenderHint(gui.QPainter__Antialiasing, false)
	painter.Pen().SetCapStyle(core.Qt__FlatCap)
	painter.Pen().SetJoinStyle(core.Qt__MiterJoin)

	w, h := f.Width(), f.Height()
	m := f.margin
	nutM := 8
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

	painter.End()
}
