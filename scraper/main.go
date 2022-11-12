package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	println("Enter a one word topic : ")
	var topic string
	fmt.Scan(&topic)

	newScanner := colly.NewCollector()
	newScanner.OnHTML("span.dt", func(h *colly.HTMLElement) {
		println(h.ChildText("span[class=dtText]"))
	})
	newScanner.Visit("https://www.merriam-webster.com/dictionary/" + topic)

	newScanner1 := colly.NewCollector()
	newScanner1.OnHTML("div. bRMDJf islir", func(h *colly.HTMLElement) {
		println(h.ChildText("img[data-deffered=1]"))
	})
	newScanner1.Visit("https://www.google.com/search?q=" + topic + "&tbm=isch")
}
