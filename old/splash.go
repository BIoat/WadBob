package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func setupSplashWindow(w fyne.Window, i *instance) {
	w.SetTitle("WadFix")
	img := canvas.NewImageFromResource(resourceSplashPng)
	img.FillMode = canvas.ImageFillOriginal
	img.ScaleMode = canvas.ImageScaleFastest
	setupButton := widget.NewButtonWithIcon("Setup", resourceIconIco, func() { initialSetup(w, i) })
	aboutButton := widget.NewButtonWithIcon("About", theme.InfoIcon(), func() {
		dialog.ShowInformation("WadFix", ".:: WadFix | WadBot Utilities | WadBob ::.\n\nhttps://github.com/bioat", w)
	})
	closeButton := widget.NewButtonWithIcon("Exit", theme.NavigateBackIcon(), func() { confirmexit(w) })
	setupButton.Alignment = widget.ButtonAlignLeading
	setupButton.Move(fyne.NewPos(50, 50))

	w.SetContent(
		container.NewVBox(
			img,
			container.NewHBox(
				setupButton,
				aboutButton,
				closeButton,
			),
		),
	)
}
