package main

import (
	"flag"
	"math/rand"
	"strings"
	"time"
)

func main() {
	wordsPtr := flag.String("words",
		"black, silver, gray, white, maroon, red, purple, fuchsia, green, lime, olive, yellow, navy, blue, teal, aqua, aquamarine",
		"Comma seperated value of the words to use.")

	flag.BoolVar(&diagonals, "diag", true, "If diagonals should be used or not.")
	flag.BoolVar(&backwards, "backward", true, "If backwards should be used or not.")
	flag.BoolVar(&stopFill, "stopFill", false, "Stop filling unused cells with random chars")
	flag.BoolVar(&capitalize, "cap", true, "If true capatlized, if false lowercase")

	flag.IntVar(&width, "w", 15, "Width of the grid")
	flag.IntVar(&height, "h", 15, "Height of the grid")

	flag.Parse()

	parseWords(*wordsPtr)

	rand.Seed(time.Now().UnixNano())
	Generate()

	ConsolePrintGrid()
}

func parseWords(newWords string) {
	words = strings.Split(newWords, ",")

	for i, s := range words {
		words[i] = strings.TrimSpace(s)
	}
}
