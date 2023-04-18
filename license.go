package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2/canvas"
	"github.com/hako/durafmt"
)

func stripRegex(in string) (int, int) {
	reg, _ := regexp.Compile("[^a-zA-Z0-9 \\ ]+")
	explode := strings.Split(reg.ReplaceAllString(in, ""), "br")

	ver, err := strconv.Atoi(explode[0])
	expire, err := strconv.Atoi(explode[1])
	if err != nil {
		fmt.Println(err)
	}
	return ver, expire
}

func checkLicense(i *instance, l *canvas.Text) {
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
		_, i.expires = stripRegex(string(body))
		l.Text = "[7.4] Time left: " + sectotime(i)
		l.Refresh()
	}
	defer resp.Body.Close()
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
		i.timeractive = true // Modify timeractive to be part of the instance struct
		return fmt.Sprint(ret)

	}
}
