package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func confirm(prompt string, tries int) bool {
	reader := bufio.NewReader(os.Stdin)

	for tries > 0 {
		fmt.Printf("%s [y/n]: ", prompt)
		response, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		response = strings.ToLower(strings.TrimSpace(response))

		if len(response) == 0 {
			tries--
			continue
		}

		return response[0] == 'y'
	}

	return false
}

func main() {
	playtune(resourceTuneMp3.Content())
	loadgui()
}
