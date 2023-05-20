package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)



func home(a fyne.App, i *instance) fyne.Window {
	g := &guiElements{
		licenseHeader: canvas.NewText("WadBot Utilities [Checking License...]",
			color.RGBA{R: 247, G: 173, B: 0, A: 230}),
		remainingTimeLabel: widget.NewLabel("Remaining Time - Unknown"),
	}

	g.licenseHeader.TextSize = 26
	go checkLicense(i, g.licenseHeader)

	w := a.NewWindow("WadBot Utilities")
	w.SetContent(container.NewVBox(g.licenseHeader, g.remainingTimeLabel))
	w.CenterOnScreen()

	i.window = w
	return w
}
