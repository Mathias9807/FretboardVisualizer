package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

var SelectedKey = 0
var SelectedMode = 0

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)

	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle("Fretboard")

	widget := widgets.NewQWidget(nil, 0)
	widget.SetLayout(widgets.NewQVBoxLayout())
	window.SetCentralWidget(widget)

	topBar := widgets.NewQWidget(nil, 0)
	hBox := widgets.NewQHBoxLayout()
	topBar.SetLayout(hBox)
	widget.Layout().AddWidget(topBar)

	hBox.AddStretch(0)
	labelKey := widgets.NewQLabel2("Key", nil, 0)
	labelKey.SetAlignment(core.Qt__AlignRight | core.Qt__AlignCenter)
	topBar.Layout().AddWidget(labelKey)
	comboKey := widgets.NewQComboBox(nil)
	comboKey.AddItems([]string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"})
	comboKey.ConnectCurrentIndexChanged(func(index int) {
		SelectedKey = index
		window.Update()
	})

	topBar.Layout().AddWidget(comboKey)
	hBox.AddStretch(0)

	labelMode := widgets.NewQLabel2("Mode", nil, 0)
	labelMode.SetAlignment(core.Qt__AlignRight | core.Qt__AlignCenter)
	topBar.Layout().AddWidget(labelMode)
	comboMode := widgets.NewQComboBox(nil)
	comboMode.AddItems([]string{"Ionian (Major)", "Dorian", "Phrygian", "Lydian", "Mixolydian", "Aeolian (Minor)", "Locrian"})
	comboMode.ConnectCurrentIndexChanged(func(index int) {
		SelectedMode = index
		window.Update()
	})
	topBar.Layout().AddWidget(comboMode)
	hBox.AddStretch(0)

	fretboard := NewFretboard()
	widget.Layout().AddWidget(fretboard)

	window.Show()
	window.Layout().Invalidate()
	window.Hide()
	sCenter := app.QApplication_PTR().ScreenAt(window.Geometry().Center()).AvailableGeometry().Center()
	wSize := window.Geometry().Size()
	window.Move2(sCenter.X()-wSize.Width()/2, sCenter.Y()-wSize.Height()*2/3)
	window.Show()
	app.Exec()
}
