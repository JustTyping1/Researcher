package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func scrape(topic string) {
	file, err := os.Create("data.txt")
	if err != nil {
		println(err)
		return
	}
	defer file.Close()
	file.WriteString(topic)
	file.WriteString("\n")
	newScanner := colly.NewCollector()
	newScanner.OnHTML("span.dt", func(h *colly.HTMLElement) {
		file.WriteString(h.ChildText("span[class=dtText]") + "\n")
	})
	newScanner.Visit("https://www.merriam-webster.com/dictionary/" + topic)

	newScanner1 := colly.NewCollector()
	newScanner1.OnHTML("a[data-test-id=image-result-link]", func(x *colly.HTMLElement) {
		file.WriteString(x.ChildAttr("img[class=image-result__image]", "src") + "\n")
	})
	newScanner1.Visit("https://www.ecosia.org/images?q=" + topic)
}

func main() {
	println("Enter a one word topic : ")
	var topic string
	fmt.Scan(&topic)
	scrape(topic)
}
