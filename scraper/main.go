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
	newScanner1.OnHTML("a[data-test-id=image-result-link]", func(x *colly.HTMLElement) {
		println(x.ChildAttr("img[class=image-result__image]", "src"))
	})
	newScanner1.Visit("https://www.ecosia.org/images?q=" + topic)
}
