package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)


func getTagData(url string, tag string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("ERROR: Failed to crawl:", url)
		return
	}

	b := resp.Body
	defer b.Close() // close Body when the function completes

	z := html.NewTokenizer(b)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			// Check if the token is a tag
			isAnchor := t.Data == tag
			if isAnchor {
				tokenType := z.Next()
				//just make sure it's actually a text token
				if tokenType == html.TextToken {
				    fmt.Println(z.Token().Data)
				    break
				}
			}
		}
	}
}

func main() {
	url := os.Args[1]
	tag:= os.Args[2]

	getTagData(url,tag)
}
