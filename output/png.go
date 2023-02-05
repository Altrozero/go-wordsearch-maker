package output

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"math"
	"os"
	"strings"

	"github.com/golang/freetype"
	"golang.org/x/image/font"
)

const dpi = 300.0

const fontFile = "fonts/RobotoMono-Regular.ttf"
const fontFileBold = "fonts/RobotoMono-Bold.ttf"
const titleFontSize = 30
const maxFontSize = 30
const maxSecondaryFontSize = 18
const minFontSize = 5
const fixedSpacedRatio = 0.6

// Size of an A4 Page and layout
var imgSize = square{0, 0, 2480, 3508}
var titleBox = square{300, 200, 1880, 100}
var gridBox = square{300, 450, 1880, 1880}
var wordsBox = square{300, 2500, 1880, 900}

const wordColumns = 3

type square struct {
	X, Y, Width, Height int
}

type PngConfig struct {
	File  string
	Title string
}

func SaveToPNG(grid [][]rune, capitalize bool, pngConfig PngConfig, placed []string) {
	img := setupImage()

	writeTitle(img, pngConfig.Title)

	str := createGridString(grid, capitalize)
	writeGrid(img, imgSize.Width, str)

	writeWords(img, placed)

	// Encode as PNG.
	f, err := os.Create(pngConfig.File)
	if err != nil {
		log.Println(err)
		return
	}
	png.Encode(f, img)
}

func setupImage() (img *image.RGBA) {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{imgSize.Width, imgSize.Height}

	img = image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Draw a white background
	draw.Draw(img, img.Bounds(), image.White, image.Point{}, draw.Src)

	return
}

func createGridString(grid [][]rune, capitalize bool) (output []string) {
	for _, line := range grid {
		lineStr := ""

		for _, char := range line {
			lineStr = lineStr + string(char) + "  "
		}

		if capitalize {
			lineStr = strings.ToUpper(lineStr)
		} else {
			lineStr = strings.ToLower(lineStr)
		}

		output = append(output, lineStr)
	}

	return output
}

func setupFont(fontName string, fontSize float64, img *image.RGBA) (*freetype.Context, error) {
	c := freetype.NewContext()

	fontBytes, err := os.ReadFile(fontName)

	if err != nil {
		return c, err
	}
	f, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return c, err
	}

	c.SetDPI(dpi)
	c.SetFont(f)
	c.SetFontSize(fontSize)
	c.SetSrc(image.Black)
	c.SetHinting(font.HintingNone)
	c.SetClip(img.Bounds())
	c.SetDst(img)

	return c, nil
}

func writeTitle(img *image.RGBA, text string) {
	fontSize := calculateFontSize(len(text), 1, 1, titleBox, titleFontSize)

	c, err := setupFont(fontFileBold, fontSize, img)

	if err != nil {
		fmt.Println("Failed to load the font.")
		return
	}

	x := xToCenter(len(text), titleBox.Width, fontSize)

	pt := freetype.Pt(titleBox.X+x, titleBox.Y+int(c.PointToFixed(titleFontSize)>>6))
	_, err = c.DrawString(text, pt)
	if err != nil {
		log.Println(err)
		return
	}
}

func writeGrid(img *image.RGBA, width int, grid []string) {
	fontSize := calculateFontSize(len(grid[0]), len(grid), 1.75, gridBox, maxFontSize)

	c, err := setupFont(fontFile, fontSize, img)

	if err != nil {
		fmt.Println("Failed to load the font.")
		return
	}

	x := xToCenter(len(grid[0]), gridBox.Width, fontSize)

	// Draw the text.
	pt := freetype.Pt(gridBox.X+x, gridBox.Y+int(c.PointToFixed(fontSize)>>6))

	for _, text := range grid {
		_, err := c.DrawString(text, pt)
		if err != nil {
			log.Println(err)
			return
		}
		pt.Y += c.PointToFixed(fontSize * 1.75)
	}
}

func writeWords(img *image.RGBA, words []string) {
	maxWordsInColumn := int(math.Ceil(float64(len(words)) / 3))
	columnWidth := wordsBox.Width / wordColumns
	singleColumnBox := square{0, 0, int(float64(columnWidth) * 0.9), wordsBox.Height}

	fontSize := calculateFontSize(len(words[0]), maxWordsInColumn, 1.25, singleColumnBox, maxSecondaryFontSize)

	c, err := setupFont(fontFile, fontSize, img)

	if err != nil {
		fmt.Println("Failed to load the font.")
		return
	}

	y := wordsBox.Y

	// Setup 3 columns
	col := 0
	for _, text := range words {
		x := wordsBox.X + ((col % wordColumns) * columnWidth)

		pt := freetype.Pt(x, y)

		_, err := c.DrawString(text, pt)
		if err != nil {
			log.Println(err)
			return
		}

		col++
		if col%3 == 0 {
			y += int(convertPtToPixels(fontSize) * 1.25)
		}
	}
}

func calculateFontSizeByWidth(characters int, maxWidth float64) (pt float64) {
	pxPerChar := ((maxWidth / float64(characters)) / fixedSpacedRatio)

	pt = pxPerChar / (dpi / 72)

	if pt > maxFontSize {
		return maxFontSize
	} else if pt < minFontSize {
		return minFontSize
	}

	return
}

func calculateFontSizeByHeight(characters int, lineSpacing float64, maxHeight float64) (pt float64) {
	pxPerChar := (maxHeight / float64(characters)) / lineSpacing

	pt = pxPerChar / (dpi / 72)

	if pt > maxFontSize {
		return maxFontSize
	} else if pt < minFontSize {
		return minFontSize
	}

	return
}

func calculateFontSize(colLen, rowLen int, lineSpacing float64, box square, maxSize float64) (pt float64) {
	pt = calculateFontSizeByWidth(colLen, float64(box.Width))
	ptByHeight := calculateFontSizeByHeight(rowLen, lineSpacing, float64(box.Height))

	if ptByHeight < pt {
		pt = ptByHeight
	}

	if pt > maxSize {
		pt = maxSize
	}

	return pt
}

func convertPtToPixels(pt float64) float64 {
	return (dpi / 72) * pt
}

func xToCenter(strLen, width int, pt float64) int {
	return (width / 2) - int((convertPtToPixels(pt)*float64(strLen)*fixedSpacedRatio)/2)
}
