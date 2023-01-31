package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

const fillText = "eeeeeeeeeeeeettttttttttaaaaaaaaaooooooooiiiiiiiinnnnnnnsssssssrrrrrrrhhhhhhdddddlllluuucccmmmfffyyywwwgggppbbvvkxqjz"

type config struct {
	words      []string
	backwards  bool
	diagonals  bool
	stopFill   bool
	capitalize bool
	width      int
	height     int
}

func createNewGrid(width, height int) [][]rune {
	newGrid := make([][]rune, height)

	for i := range newGrid {
		newGrid[i] = make([]rune, width)
	}

	return newGrid
}

func addWordsToGrid(conf config) (
	grid [][]rune,
	placed []string,
	failed []string,
) {
	grid = createNewGrid(conf.width, conf.height)

	// Sort the words by length
	sort.SliceStable(conf.words, func(i, j int) bool {
		return len(conf.words[i]) > len(conf.words[j])
	})

	// Loop over the words
	for {
		var outcome bool

		outcome, grid = tryFindPlaceOnGrid(grid, conf.words[0], conf.diagonals, conf.backwards)

		if outcome {
			placed = append(placed, conf.words[0])
		} else {
			failed = append(failed, conf.words[0])
		}

		conf.words = conf.words[1:]

		if len(conf.words) == 0 {
			break
		}
	}

	return
}

func tryFindPlaceOnGrid(grid [][]rune, word string, diagonals, backwards bool) (bool, [][]rune) {
	localDirections := setupPossibleDirections(diagonals, backwards)

	for {
		positions := setupPossiblePositions(grid)

		// Loop over positions
		for {
			row := positions[0] / len(grid)
			col := positions[0] % len(grid[0])
			outcome := false

			outcome, grid = tryPutOnGrid(grid, word, row, col, localDirections[0])

			if outcome {
				return true, grid
			}

			positions = positions[1:]

			if len(positions) == 0 {
				break
			}
		}

		localDirections = localDirections[1:]

		if len(localDirections) == 0 {
			break
		}
	}

	return false, grid
}

func tryPutOnGrid(grid [][]rune, word string, row, col int, dir [2]int) (bool, [][]rune) {
	cloneGrid := make([][]rune, len(grid))
	for i := range grid {
		cloneGrid[i] = make([]rune, len(grid[i]))
		copy(cloneGrid[i], grid[i])
	}

	// Place the letters
	for {
		if cloneGrid[row][col] != rune(0) &&
			cloneGrid[row][col] != rune(word[0]) {

			return false, grid
		}

		cloneGrid[row][col] = rune(word[0])

		row += dir[0]
		col += dir[1]

		word = word[1:]

		if len(word) == 0 {
			break
		}

		if row < 0 || row >= len(grid) ||
			col < 0 || col >= len(grid[0]) {

			return false, grid
		}
	}

	return true, cloneGrid
}

func setupPossibleDirections(diagonals, backwards bool) [][2]int {
	var directions [][2]int

	if len(directions) == 0 {
		directions = append(directions, [2]int{1, 0}, [2]int{0, 1})

		if diagonals {
			directions = append(directions, [2]int{1, 1})
		}
		if backwards {
			directions = append(directions, [2]int{-1, 0}, [2]int{0, -1})
		}
		if diagonals && backwards {
			directions = append(directions, [2]int{-1, -1})
		}
	}

	rand.Shuffle(len(directions), func(i, j int) { directions[i], directions[j] = directions[j], directions[i] })

	return directions
}

func setupPossiblePositions(grid [][]rune) []int {
	var pos []int
	size := len(grid) * len(grid[0])

	for i := 0; i < size; i++ {
		pos = append(pos, i)
	}

	rand.Shuffle(len(pos), func(i, j int) { pos[i], pos[j] = pos[j], pos[i] })

	return pos
}

func fillGrid(grid [][]rune, stopFill bool) [][]rune {
	chars := []rune(fillText)
	rand.Shuffle(len(chars), func(i, j int) { chars[i], chars[j] = chars[j], chars[i] })

	for i, line := range grid {
		for j, char := range line {
			if char == rune(0) {
				if stopFill {
					grid[i][j] = rune(' ')
				} else {
					char := chars[0]

					chars = chars[1:]
					chars = append(chars, char)

					grid[i][j] = char
				}
			}
		}
	}

	return grid
}

func ConsolePrintGrid(grid [][]rune, capitalize bool, placed, failed []string) {
	var (
		output string
		buf    strings.Builder
	)

	for _, line := range grid {
		for _, char := range line {
			buf.WriteRune(char)
			buf.WriteString(" ")
		}
		buf.WriteByte('\n')
	}

	if capitalize {
		output = strings.ToUpper(buf.String())
	} else {
		output = strings.ToLower(buf.String())
	}

	fmt.Println(output)

	fmt.Println("Placed Words:")
	fmt.Println(placed)

	fmt.Println(" ")

	fmt.Println("Failed to place:")
	fmt.Println(failed)
}

func Generate(myConfig config) ([][]rune, []string, []string) {
	grid, placed, failed := addWordsToGrid(myConfig)

	return fillGrid(grid, myConfig.stopFill), placed, failed
}
