package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
)

type Item struct {
	Chapter    string      `json:"chapter"`
	Paragraphs []Paragraph `json:"content"`
}

type Paragraph struct {
	Content string
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	defer timer("main")()
	c := colly.NewCollector(colly.Async(true))

	items := []Item{}

	c.OnHTML("li.chapter a", func(h *colly.HTMLElement) {
		link := h.Attr("href")
		c.Visit(h.Request.AbsoluteURL(link))
	})

	// c.OnHTML("li.next a", func(h *colly.HTMLElement) {
	// 	c.Visit(h.Request.AbsoluteURL(h.Attr("href")))
	// })

	c.OnHTML("body", func(h *colly.HTMLElement) {

		contents := []Paragraph{}

		h.ForEach("main p", func(i int, h *colly.HTMLElement) {
			p := Paragraph{
				Content: h.Text,
			}
			contents = append(contents, p)
		})

		i := Item{
			Chapter:    h.ChildText("header h1"),
			Paragraphs: contents,
			// Name:    h.ChildAttr("h3 a", "title"),
			// Price:   h.ChildText("p.price_color"),
		}

		items = append(items, i)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting ", r.URL)
	})

	// c.Visit("https://books.toscrape.com/catalogue/category/books_1/index.html")
	// c.Visit("https://books.toscrape.com/catalogue/category/books/travel_2/index.html")
	c.Visit("https://hypermedia.systems/book/contents/")
	c.Wait()

	data, err := json.MarshalIndent(items, " ", "")
	if err != nil {
		log.Fatal()
	}

	os.WriteFile("second-version.json", data, os.ModePerm)
	// fmt.Println(string(data))
}

// https://books.toscrape.com/catalogue/category/books_1/index.html
