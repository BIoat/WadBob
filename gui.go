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

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
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

func loadgui() {
	url := "https://wadbot.lol/WadBot/check.php"
	a := app.New()
	w := a.NewWindow("WadBot Utilities")
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
      //placeholder
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
	w.ShowAndRun()

	tidyUp()
}

func tidyUp() {
	fmt.Println("GUI exit")
}
