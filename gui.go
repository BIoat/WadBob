package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unsafe"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/hako/durafmt"
)

var (
	expire      int
	ver         int
	timeractive bool
)

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func sectotime(n int) string {
	if n == 0 {
		return "Invalid"
	} else if n == 1 {
		return "Unused"
	} else if n == 99999 {
		return "Lifetime"
	} else if &n == nil {
		return "No key"
	} else {
		ret, _ := durafmt.ParseString(fmt.Sprintf("%ds", n))

		expire -= 1
		timeractive = true
		return fmt.Sprint(ret)

	}
}

func stripRegex(in string) (int, int) {
	reg, _ := regexp.Compile("[^a-zA-Z0-9 \\ ]+")
	explode := strings.Split(reg.ReplaceAllString(in, ""), "br")

	ver, err := strconv.Atoi(explode[0])
	expire, err := strconv.Atoi(explode[1])
	if err != nil {
		fmt.Println("Error during conversion")
	}
	return ver, expire
}

func updateTime(timer *widget.Label) {
	if timeractive {
		formatted := "[ 7.4 ] Time left: " + sectotime(expire)
		timer.SetText(formatted)
	}
}

func login(a fyne.App) fyne.Window {
	// md := getnews()
	url := "https://wadbot.lol/WadBot/check.php"
	w := a.NewWindow("WadBot Utilities")

	// news := widget.NewRichTextFromMarkdown(md)
	// news.Wrapping = fyne.TextWrapWord
	l2 := canvas.NewText("WadBot Utilities",
		color.RGBA{R: 247, G: 173, B: 0, A: 230})
	l2.TextSize = 26
	// get key from user
	in := widget.NewEntry()
	in.PlaceHolder = "Enter Your License"

	l := widget.NewLabel("Remaining Time - Unknown")
	btn := widget.NewButton("Login", func() {
		key := in.Text
		req, _ := http.NewRequest("GET", url+"?key="+key, nil)
		req.Header.Add("Accept", "application/json")
		client := &http.Client{}
		resp, _ := client.Do(req)
		if resp.StatusCode != 200 {
			l.Text = "Bad Connection"
			l.Refresh()
		} else {
			body, _ := ioutil.ReadAll(resp.Body)
			ver, expire = stripRegex(string(body))
			// placeholder
			l.Text = "[7.4] Time left: " + sectotime(expire)
			l.Refresh()

			println(ver)
		}
		defer resp.Body.Close()
	})
	go func() {
		for range time.Tick(time.Second) {
			updateTime(l)
		}
	}()

	w.SetContent(container.NewVBox(l2, in, l, btn))
	w.CenterOnScreen()
	return w
}

func loadgui() {
	a := app.NewWithID("WadFix")
	fyne.CurrentApp().SetIcon(resourceIconIco)
	a.Settings().SetTheme(theme.DarkTheme())
	loginwin := login(a)
	if drv, ok := fyne.CurrentApp().Driver().(desktop.Driver); ok {
		w := drv.CreateSplashWindow()
		w.SetTitle("WadFix")
		img := canvas.NewImageFromResource(resourceSplashPng)
		img.FillMode = canvas.ImageFillOriginal
		img.ScaleMode = canvas.ImageScaleFastest
		startbutton := widget.NewButtonWithIcon("Start", resourceIconIco, func() {
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
					w.Close()
				} else {
					// go pauseresumetune()
				}
			}, w)

			// w.Close()
		})
		aboutbutton := widget.NewButtonWithIcon("About", theme.InfoIcon(), func() {
			dialog.ShowInformation("WadFix", ".:: WadFix | WadBot Utilities | WadBob ::.\n\nhttps://github.com/bioat", w)
		})
		closebutton := widget.NewButtonWithIcon("Exit", theme.NavigateBackIcon(), func() {
			go pauseresumetune()
			dialog.ShowConfirm("Quit", "Sure?", func(b bool) {
				if b {
					a.Quit()
				} else {
					go pauseresumetune()
				}
			}, w)
		})
		startbutton.Alignment = widget.ButtonAlignLeading
		startbutton.Move(fyne.NewPos(50, 50))

		w.SetContent(
			container.NewVBox(
				img,
				container.NewHBox(
					startbutton,
					aboutbutton,
					closebutton,
				),
			),
		)
		w.SetOnClosed(func() {
			loginwin.Show()
		})

		go func() {
			for i := 1; i <= 75; i++ {
				time.Sleep(time.Millisecond)
				w.CenterOnScreen()
			}
		}()
		w.ShowAndRun()
		w.CenterOnScreen()
	}

	tidyUp()
}

func tidyUp() {
	stoptune()

	fmt.Println("GUI exit")
}
