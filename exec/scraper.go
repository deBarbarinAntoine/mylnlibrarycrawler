package main

import (
	"fmt"
	colly "github.com/gocolly/colly/v2"
	termsize "github.com/kopoli/go-terminal-size"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	utils "webscraping"
)

type Library struct {
	Name  string
	Books []LightNovel
}

type LightNovel struct {
	Url    string
	Name   string
	Author string
	Genre  []string
	Status string
}

type DisplayParameters struct {
	InvariableColor string
	LoadingBarColor string
	ArrowColor      string
	VariableColor   string
	AlertColor      string
}

var myDisplayParams = DisplayParameters{
	InvariableColor: utils.ColorCode(utils.Orange),
	LoadingBarColor: utils.ColorCode(utils.Teal),
	ArrowColor:      utils.ColorCode(utils.Aquamarine),
	VariableColor:   utils.ColorCode(utils.Coral),
	AlertColor:      utils.ColorCode(utils.Orangered),
}

var myLightNovels Library

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
		return myDisplayParams.InvariableColor + invariable + " " + myDisplayParams.VariableColor + variable + utils.CLEARCOLOR
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
			result = myDisplayParams.InvariableColor + invariable + utils.CLEARCOLOR + " " + myDisplayParams.VariableColor + variable + utils.CLEARCOLOR
		} else {
			result = myDisplayParams.InvariableColor + invariable + utils.CLEARCOLOR + " " + myDisplayParams.LoadingBarColor + progressBar + utils.CLEARCOLOR + myDisplayParams.ArrowColor + " > " + utils.CLEARCOLOR + myDisplayParams.VariableColor + variable + utils.CLEARCOLOR
		}

		colorRegExp := regexp.MustCompile("\\033\\[[0-9;]+m")
		nbColors := len(colorRegExp.FindAllString(result, -1))
		str := strings.Join(colorRegExp.Split(result, -1), "")

		if rest := width - len([]rune(str)); rest < 0 {
			result = string([]rune(result)[:len([]rune(result))+rest-nbColors])
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

func (lib *Library) getLightNovels() {
	for i := 1; i < 76; i++ {
		var newLN LightNovel
		c := colly.NewCollector()

		c.OnHTML(".novel-title", func(e *colly.HTMLElement) {
			newLN.Url = e.ChildAttr("a", "href")
			newLN.Name = e.ChildAttr("a", "title")
			lib.Books = append(lib.Books, newLN)
		})

		err := c.Visit("https://novel-next.com/sort/completed-novelnext?page=" + strconv.Itoa(i))
		if err != nil {
			log.Fatal(err)
		}
		clearCurrentLine()
		fmt.Print("\r", resize(i, 75, "Processing... ["+strconv.Itoa(i)+"/75]", newLN.Name))
	}
	fmt.Println()
}

func (lib *Library) fetchLN() {
	lib.getLightNovels()

	for i := range lib.Books {
		clearCurrentLine()
		fmt.Print("\r", resize(i+1, len(lib.Books), "Processing... ["+strconv.Itoa(i+1)+"/"+strconv.Itoa(len(lib.Books))+"]", lib.Books[i].Name))
		myLN := &lib.Books[i]
		myLN.getInfo()
	}
	fmt.Println()
	fmt.Printf("My LightNovels library: %#v", lib)
	fmt.Println()
}

func instructions() {
	fmt.Println(myDisplayParams.ArrowColor, "Welcome to myLnLibraryCrawler!", utils.CLEARCOLOR)
	fmt.Println(myDisplayParams.InvariableColor, "Miscellaneous commands (type it and press [Enter]):", utils.CLEARCOLOR)
	fmt.Println(myDisplayParams.AlertColor, "     'stop' -> exit the program", utils.CLEARCOLOR)
	fmt.Println(myDisplayParams.VariableColor, "     'skip' -> go to the next part of the program", utils.CLEARCOLOR)
	fmt.Println(myDisplayParams.InvariableColor, "  [Any other input will show the library's current content]", utils.CLEARCOLOR)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	lib := &myLightNovels
	instructions()
	go lib.fetchLN()
	for {
		var input string
		fmt.Scanln(&input)
		if input == "stop" {
			os.Exit(0)
		} else if input == "skip" {
			break
		} else {
			fmt.Println()
			fmt.Printf("My LightNovels library: %#v", myLightNovels)
			fmt.Println()
			instructions()
		}
	}
	wg.Wait()
}
