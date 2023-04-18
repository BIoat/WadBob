package main

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type guiElements struct {
	remainingTimeLabel *widget.Label
	timeractive        bool
	licenseHeader      *canvas.Text
}

func createGUIElements() *guiElements {
	e := &guiElements{
		remainingTimeLabel: widget.NewLabel("Remaining Time - Unknown"),
	}
	return e
}

func updateTime(e *guiElements, i *instance) {
	if e.timeractive {
		formatted := "[ 7.4 ] Time left: " + sectotime(i)
		e.remainingTimeLabel.SetText(formatted)
	}
}
