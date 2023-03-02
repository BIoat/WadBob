package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	// "time"
)

func file_exists(path string) bool {
	f, _ := os.ReadFile(path)
	return f != nil
}

func confirm(s string, tries int) bool {
	r := bufio.NewReader(os.Stdin)
	for ; tries > 0; tries-- {
		fmt.Printf("%s [y/n]: ", s)
		res, err := r.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		// Empty input (i.e. "\n")
		if len(res) < 2 {
			continue
		}
		return strings.ToLower(strings.TrimSpace(res))[0] == 'y'
	}
	return false
}

func main() {
	go playtune(resourceTuneMp3.Content())
	loadgui()
	//
	// name := confirm("Check if LoL Config exists?", 3)
	// if name {
	// 	// placeholder
	// 	r := file_exists("C:/Riot Games/League of Legends/Config/PersistedSettings.json")
	// 	println(r)
	// 	time.Sleep(time.Second)
	// }
	//
	// println("quit")
	// os.Exit(0)
}
