package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func confirmexit(w fyne.Window) {
	dialog.ShowConfirm("Exit", "Are you sure you want to exit?", func(response bool) {
		if response {
			w.Close()
		}
	}, w)
}
