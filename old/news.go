package main

import (
	"io"
	"net/http"

	md "github.com/JohannesKaufmann/html-to-markdown"
)

func test() {
}

func getnews() (string){
	bodyString := ""
	url := ("https://view.matrix.org/room/!nrtkImjZNafxHXavjT:matrix.org/?anchor=$1weXEU045jrVvpRU_t9c6mlVL4-d8Do7_YIfIWhnvd8&offset=1")
  url = "https://sleep.codes"
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			println(err)
		}
		bodyString = string(bodyBytes)
	}
  // converter := md.NewConverter("view.matrix.org/", true, nil)
  converter := md.NewConverter("sleep.codes", true, nil)
	markdown, err := converter.ConvertString(bodyString)
	if err != nil {
    println(err)
	}
  println(markdown)
  return markdown
}
