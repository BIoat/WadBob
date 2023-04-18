package main

import "fyne.io/fyne/v2"

type instance struct {
	license string
	version string
	path    string
	expires int
  window fyne.Window
  timeractive bool
}
