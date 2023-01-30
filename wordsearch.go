package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

var (
	words      = []string{}
	backwards  = false
	diagonals  = false
	stopFill   = false
	capitalize = true
	width      = 15
	height     = 15
)

var (
	grid       = [][]rune{}
	directions = [][2]int{}
	placed     = []string{}
	failed     = []string{}
	chars      = []rune("abcdefghijklmnopqrstuvwxyz")
)

func createNewGrid() [][]rune {
	newGrid := make([][]rune, height)

	for i := range newGrid {
		newGrid[i] = make([]rune, width)
	}

	return newGrid
}

func addWordsToGrid() {
	// Sort the words by length
	sort.SliceStable(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})

	// Loop over the words
	for {
		localDirections := setupPossibleDirections()

	directions:
		for {
			positions := setupPossiblePositions()

			// Loop over positions
			for {
				row := positions[0] / len(grid)
				col := positions[0] % len(grid[0])

				if tryPlaceOnGird(words[0], row, col, localDirections[0]) {
					placed = append(placed, words[0])
					break directions
				}

				positions = positions[1:]

				if len(positions) == 0 {
					break
				}
			}

			localDirections = localDirections[1:]

			if len(localDirections) == 0 {
				failed = append(failed, words[0])

				break
			}
		}

		words = words[1:]

		if len(words) == 0 {
			break
		}
	}
}

func tryPlaceOnGird(word string, row int, col int, dir [2]int) bool {
	cloneGrid := make([][]rune, len(grid))
	for i := range grid {
		cloneGrid[i] = make([]rune, len(grid[i]))
		copy(cloneGrid[i], grid[i])
	}

	// Place the letters
	for {
		if cloneGrid[row][col] != rune(0) &&
			cloneGrid[row][col] != rune(word[0]) {
			return false
		}

		cloneGrid[row][col] = rune(word[0])

		row += dir[0]
		col += dir[1]

		word = word[1:]

		if len(word) == 0 {
			break
		}

		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
			return false
		}
	}

	grid = cloneGrid

	return true
}

func setupPossibleDirections() [][2]int {
	if len(directions) == 0 {
		directions = append(directions, [2]int{1, 0})
		directions = append(directions, [2]int{0, 1})

		if diagonals {
			directions = append(directions, [2]int{1, 1})
		}
		if backwards {
			directions = append(directions, [2]int{-1, 0})
			directions = append(directions, [2]int{0, -1})

			if diagonals {
				directions = append(directions, [2]int{-1, -1})
			}
		}
	}

	rand.Shuffle(len(directions), func(i, j int) { directions[i], directions[j] = directions[j], directions[i] })

	return directions
}

func setupPossiblePositions() []int {
	pos := []int{}
	size := len(grid) * len(grid[0])

	for i := 0; i < size; i++ {
		pos = append(pos, i)
	}

	rand.Shuffle(len(pos), func(i, j int) { pos[i], pos[j] = pos[j], pos[i] })

	return pos
}

func fillGrid() {
	rand.Shuffle(len(chars), func(i, j int) { chars[i], chars[j] = chars[j], chars[i] })

	for i, line := range grid {
		for j, char := range line {
			if char == rune(0) {
				if stopFill {
					grid[i][j] = rune(' ')
				} else {
					grid[i][j] = pullRandCharacter()
				}
			}
		}
	}
}

func ConsolePrintGrid() {
	output := ""

	for _, line := range grid {
		for _, char := range line {
			output = output + string(char) + " "
		}
		output = output + "\n"
	}

	if capitalize {
		output = strings.ToUpper(output)
	} else {
		output = strings.ToLower(output)
	}

	fmt.Println(output)

	fmt.Println("Placed Words:")
	fmt.Println(placed)

	fmt.Println(" ")

	fmt.Println("Failed to place:")
	fmt.Println(failed)
}

/*
Pull a random char from the possible chars and send to the back of the list
Currently better than random letters as some runs were coming out with a lot
of X and Z's
TODO: Replace with a more complete method
*/
func pullRandCharacter() rune {
	char := chars[0]

	chars = chars[1:]
	chars = append(chars, char)

	return char
}

func Generate() {
	grid = createNewGrid()

	addWordsToGrid()

	fillGrid()
}
