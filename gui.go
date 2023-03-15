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

type instance struct {
	license string
	version string
	path    string
	expires int
}

var (
	ver         int
	timeractive bool
	a           fyne.App
)

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func sectotime(i *instance) string {
	n := i.expires

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

		i.expires -= 1
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
		fmt.Println(err)
	}
	//
	// ret := strings.Split(in, "|")
	// println(in)
	// ver, _ = strconv.Atoi(ret[0])
	// expire, _ := strconv.Atoi(ret[1])
	println(ver)
	println(expire)
	return ver, expire
}

func updateTime(timer *widget.Label, i *instance) {
	if timeractive {
		formatted := "[ 7.4 ] Time left: " + sectotime(i)
		timer.SetText(formatted)
	}
}

func checklicense(i *instance, l *widget.Label) {
	for len(i.license) < 20 {
		time.Sleep(time.Millisecond * 150)
	}

	url := "https://wadbot.lol/WadBot/check.php"
	req, _ := http.NewRequest("GET", url+"?key="+i.license, nil)
	req.Header.Add("Accept", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	if resp.StatusCode != 200 {
		l.Text = "Bad Connection"
		l.Refresh()
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		ver, i.expires = stripRegex(string(body))
		// placeholder
		l.Text = "[7.4] Time left: " + sectotime(i)
		l.Refresh()

	}
	defer resp.Body.Close()
}

func home(a fyne.App, i *instance) fyne.Window {
	// md := getnews()
	w := a.NewWindow("WadBot Utilities")

	// news := widget.NewRichTextFromMarkdown(md)
	// news.Wrapping = fyne.TextWrapWord
	l2 := canvas.NewText("WadBot Utilities [Checking License...]",
		color.RGBA{R: 247, G: 173, B: 0, A: 230})
	l2.TextSize = 26

	l := widget.NewLabel("Remaining Time - Unknown")
	go checklicense(i, l)

	go func() {
		for range time.Tick(time.Second) {
			l2.Text = "WadBot Utilities [" + i.license + "]"
			updateTime(l, i)
			l.Refresh()
			l2.Refresh()
		}
	}()

	w.SetContent(container.NewVBox(l2, l))
	w.CenterOnScreen()
	return w
}

func InititualSetup(w fyne.Window, i *instance) fyne.Window {
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

func confirmexit(w fyne.Window) fyne.Window {
	go pauseresumetune()
	dialog.ShowConfirm("Quit", "Sure?", func(b bool) {
		if b {
			a.Quit()
		} else {
			go pauseresumetune()
		}
	}, w)
	return w
}

func forcecenter(w fyne.Window) fyne.Window {
	for i := 1; i <= 75; i++ {
		time.Sleep(time.Millisecond)
		w.CenterOnScreen()
	}

	return w
}

func loadgui() {
	i := instance{license: "wad"}
	a := app.NewWithID("WadFix")
	fyne.CurrentApp().SetIcon(resourceIconIco)
	a.Settings().SetTheme(theme.DarkTheme())
	homewin := home(a, &i)
	if drv, ok := fyne.CurrentApp().Driver().(desktop.Driver); ok {
		w := drv.CreateSplashWindow()
		w.SetTitle("WadFix")
		img := canvas.NewImageFromResource(resourceSplashPng)
		img.FillMode = canvas.ImageFillOriginal
		img.ScaleMode = canvas.ImageScaleFastest
		setupbutton := widget.NewButtonWithIcon("Setup", resourceIconIco, func() { InititualSetup(w, &i) })
		aboutbutton := widget.NewButtonWithIcon("About", theme.InfoIcon(), func() {
			dialog.ShowInformation("WadFix", ".:: WadFix | WadBot Utilities | WadBob ::.\n\nhttps://github.com/bioat", w)
		})
		closebutton := widget.NewButtonWithIcon("Exit", theme.NavigateBackIcon(), func() { confirmexit(w) })
		setupbutton.Alignment = widget.ButtonAlignLeading
		setupbutton.Move(fyne.NewPos(50, 50))

		w.SetContent(
			container.NewVBox(
				img,
				container.NewHBox(
					setupbutton,
					aboutbutton,
					closebutton,
				),
			),
		)
		w.SetOnClosed(func() {
			homewin.Show()
		})

		go forcecenter(w)
		w.ShowAndRun()
		w.CenterOnScreen()
		w.RequestFocus()
	}

	tidyUp()
}

func tidyUp() {
	stoptune()

	fmt.Println("GUI exit")
}
