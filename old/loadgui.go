package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
)

func loadgui() {
	i := &instance{license: "wad"}
	a := app.NewWithID("WadFix")
	a.Settings().SetTheme(theme.DarkTheme())
	home(a, i)

	if drv, ok := fyne.CurrentApp().Driver().(desktop.Driver); ok {
		w := drv.CreateSplashWindow()
		setupSplashWindow(w, i)
		w.ShowAndRun()
		w.CenterOnScreen()
		w.RequestFocus()
	}
	i.window.ShowAndRun()
	// tidyUp()
}
