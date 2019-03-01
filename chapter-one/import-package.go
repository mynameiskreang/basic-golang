package main

import (
	"basic-golang/chapter-one/mylib"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main() {
	xy := mylib.PublicAdd(10, 20)
	fmt.Println("mylib.PublicAdd", xy)

	// Cant Access privateAdd
	//xy := mylib.privateAdd(10, 20)
	//fmt.Println("mylib.privateAdd", xy)

	try3party()

}

func try3party() {
	fmt.Println()
	res, err := http.Get("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, band, title)
	})
}
