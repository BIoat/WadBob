package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func initialSetup(w fyne.Window, i *instance) fyne.Window {
	w.CenterOnScreen()

	license := widget.NewEntry()
	license.SetPlaceHolder("######")
	dialog.ShowForm(".:: Initial Setup ::.", "Go", "Cancel", []*widget.FormItem{
		{
			Text:   "License: ",
			Widget: license,
		},
		{
			Text: "Already installed?",
			Widget: widget.NewCheck("Existing Config", func(b bool) {
				if b {
					temp := widget.NewModalPopUp(canvas.NewText("TODO: Scan for wadbot bin", color.RGBA{1, 255, 1, 255}), fyne.CurrentApp().Driver().CanvasForObject(license))

					go func() {
						temp.Show()
						time.Sleep(time.Second)
						temp.Hide()
					}()
				}
			}),
		},
	}, func(b bool) {
		if b {
			go pauseresumetune()
			i.license = license.Text
			w.Close()
		} else {
			// todo
		}
	}, w)
	return w
}
