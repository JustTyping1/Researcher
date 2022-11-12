package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	println("Enter topic : ")
	var topic string
	fmt.Scan(&topic)

	newScanner := colly.NewCollector()
	newScanner.OnHTML("span.dt", func(h *colly.HTMLElement) {
		println(h.ChildText("span[class=dtText]"))
	})
	newScanner.Visit("https://www.merriam-webster.com/dictionary/" + topic)
}
