package main

import (
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	file, err := os.Create("data.txt")
	if err != nil {
		println(err)
		return
	}
	defer file.Close()
	println("Enter a one word topic : ")
	var topic string
	fmt.Scan(&topic)

	file.WriteString("[" + topic + ",")

	newScanner := colly.NewCollector()
	newScanner.OnHTML("span.dt", func(h *colly.HTMLElement) {
		println(h.ChildText("span[class=dtText]"))
		file.WriteString("'" + h.ChildText("span[class=dtText]") + "',")
	})
	newScanner.Visit("https://www.merriam-webster.com/dictionary/" + topic)

	newScanner1 := colly.NewCollector()
	newScanner1.OnHTML("a[data-test-id=image-result-link]", func(x *colly.HTMLElement) {
		println(x.ChildAttr("img[class=image-result__image]", "src"))
		file.WriteString("'" + x.ChildAttr("img[class=image-result__image]", "src") + "',")
	})
	newScanner1.Visit("https://www.ecosia.org/images?q=" + topic)
	file.WriteString("'1']")
}
