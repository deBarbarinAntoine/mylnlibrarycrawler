package main

import (
	"fmt"
	colly "github.com/gocolly/colly/v2"
	"log"
	"strconv"
)

type LightNovel struct {
	Url    string
	Name   string
	Author string
	Genre  []string
	Status string
}

func (ln *LightNovel) getInfo() {
	c := colly.NewCollector()
	c.OnHTML(".info-meta", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, h *colly.HTMLElement) {
			switch h.ChildText("h3") {
			case "Author:":
				ln.Author = h.ChildText("a")
			case "Genre:":
				ln.Genre = append(ln.Genre, h.ChildTexts("a")...)
			case "Status:":
				ln.Status = h.ChildText("a")
			}
		})
	})
	err := c.Visit(ln.Url)
	if err != nil {
		log.Fatal(err)
	}
}

func getLightNovels(lib []LightNovel) []LightNovel {
	for i := 1; i < 76; i++ {
		var newLN LightNovel
		c := colly.NewCollector()

		c.OnHTML(".novel-title", func(e *colly.HTMLElement) {
			newLN.Url = e.ChildAttr("a", "href")
			newLN.Name = e.ChildAttr("a", "title")
			lib = append(lib, newLN)
		})

		err := c.Visit("https://novel-next.com/sort/completed-novelnext?page=" + strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}
	}
	return lib
}

func main() {
	var myLightNovels []LightNovel

	myLightNovels = getLightNovels(myLightNovels)

	for i := range myLightNovels {
		myLN := &myLightNovels[i]
		myLN.getInfo()
	}
	fmt.Println()
	fmt.Printf("My LightNovels library: %#v", myLightNovels)
	fmt.Println()
}
