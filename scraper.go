package main

import (
	"fmt"
	colly "github.com/gocolly/colly/v2"
	termsize "github.com/kopoli/go-terminal-size"
	"log"
	"os"
	"strconv"
	"strings"
)

type LightNovel struct {
	Url    string
	Name   string
	Author string
	Genre  []string
	Status string
}

func clearCurrentLine() {
	fmt.Printf("\033[0;") // clear current line
	fmt.Printf("\033[2K\r%d", 0)
	fmt.Fprint(os.Stdout, "\033[y;0H")
	fmt.Fprint(os.Stdout, "\033[K")
	fmt.Print("\x1b[2k") // erase the current line
}

func loadingBar(current, max, space int) string {
	percentage := (float32(current) / float32(max)) * float32(space)
	progressBar := strings.Repeat("█", int(percentage)) + strings.Repeat("░", space-int(percentage))
	return progressBar
}

func resize(current, max int, invariable, variable string) string {
	var result string
	size, err := termsize.GetSize()
	if err != nil {
		return invariable + " " + variable
	} else {
		var progressBar string
		width := size.Width
		var barSize int
		if rest := width - (len(invariable) + 12); rest > 10 {
			barSize = int(float32(rest) * .6)
			progressBar = loadingBar(current, max, barSize)
		}
		lenMsg := len([]rune(fmt.Sprint(invariable, " ", progressBar, " > ", variable)))
		if len([]rune(variable)) <= -(width - lenMsg) {
			result = invariable + " " + variable
		}
		result = invariable + " " + progressBar + " > " + variable
		if rest := width - len([]rune(result)); rest < 0 {
			return string([]rune(result)[:width])
		}
		return result
	}
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
		clearCurrentLine()
		fmt.Print("\r", resize(i, 75, "Processing... ["+strconv.Itoa(i)+"/75]", newLN.Name))
	}
	fmt.Println()
	return lib
}

func main() {
	var myLightNovels []LightNovel

	myLightNovels = getLightNovels(myLightNovels)

	for i := range myLightNovels {
		clearCurrentLine()
		fmt.Print("\r", resize(i+1, len(myLightNovels), "Processing... ["+strconv.Itoa(i+1)+"/"+strconv.Itoa(len(myLightNovels))+"]", myLightNovels[i].Name))
		myLN := &myLightNovels[i]
		myLN.getInfo()
	}
	fmt.Println()
	fmt.Printf("My LightNovels library: %#v", myLightNovels)
	fmt.Println()
}
