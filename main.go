package main

import (
	"flag"
	"math/rand"
	"strings"
	"time"
)

const defaultWords = "black, silver, gray, white, maroon, red, purple, fuchsia, green, lime, olive, yellow, navy, blue, teal, aqua, aquamarine"

func main() {
	myConfig := config{}

	wordsPtr := flag.String("words", defaultWords, "Comma seperated value of the words to use.")

	flag.BoolVar(&myConfig.diagonals, "diag", true, "If diagonals should be used or not.")
	flag.BoolVar(&myConfig.backwards, "backward", true, "If backwards should be used or not.")
	flag.BoolVar(&myConfig.stopFill, "stopFill", false, "Stop filling unused cells with random chars")
	flag.BoolVar(&myConfig.capitalize, "cap", true, "If true capatlized, if false lowercase")

	flag.IntVar(&myConfig.width, "w", 15, "Width of the grid")
	flag.IntVar(&myConfig.height, "h", 15, "Height of the grid")

	flag.Parse()

	myConfig.words = parseWords(*wordsPtr)

	rand.Seed(time.Now().UnixNano())
	grid, placed, failed := Generate(myConfig)

	ConsolePrintGrid(grid, myConfig.capitalize, placed, failed)
}

func parseWords(newWords string) []string {
	words := strings.Split(newWords, ",")

	for i, s := range words {
		words[i] = strings.TrimSpace(s)
	}

	return words
}
