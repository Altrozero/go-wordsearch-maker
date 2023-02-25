package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/Altrozero/go-wordsearch-maker/output"
	"github.com/Altrozero/go-wordsearch-maker/wordsearch"
)

func main() {
	myConfig := wordsearch.Config{}
	pngConfig := output.PngConfig{}
	png := false

	wordsPtr := flag.String("words",
		"black, silver, gray, white, maroon, red, purple, fuchsia, green, lime, olive, yellow, navy, blue, teal, aqua, aquamarine",
		"Comma seperated value of the words to use.")

	flag.BoolVar(&myConfig.Diagonals, "diag", true, "If diagonals should be used or not.")
	flag.BoolVar(&myConfig.Backwards, "backward", true, "If backwards should be used or not.")
	flag.BoolVar(&myConfig.StopFill, "stopFill", false, "Stop filling unused cells with random chars")
	flag.BoolVar(&myConfig.Capitalize, "cap", true, "If true capatlized, if false lowercase")

	flag.BoolVar(&png, "png", false, "If to save to png of not, default false. Must also specify the filepath with -file")
	flag.StringVar(&pngConfig.File, "file", "", "file to save the png to")
	flag.StringVar(&pngConfig.Title, "title", "Wordsearch", "The title of the wordsearch, default wordsearch")

	flag.IntVar(&myConfig.Width, "w", 15, "Width of the grid")
	flag.IntVar(&myConfig.Height, "h", 15, "Height of the grid")

	flag.Parse()

	myConfig.Words = parseWords(*wordsPtr)

	rand.Seed(time.Now().UnixNano())
	grid, placed, failed := wordsearch.Generate(myConfig)

	output.ConsolePrintGrid(grid, myConfig.Capitalize, placed, failed)

	if png {
		if pngConfig.File != "" {
			output.SaveToPNG(grid, myConfig.Capitalize, pngConfig, placed)
		} else {
			fmt.Println("PNG: No file specified")
		}
	}
}

func parseWords(newWords string) []string {
	words := strings.Split(newWords, ",")

	for i, s := range words {
		words[i] = strings.TrimSpace(s)
	}

	return words
}
