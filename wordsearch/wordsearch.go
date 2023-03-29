package wordsearch

import (
	"math/rand"
	"sort"
)

const letterFrequency = "eeeeeeeeeeeeettttttttttaaaaaaaaaooooooooiiiiiiiinnnnnnnsssssssrrrrrrrhhhhhhdddddlllluuucccmmmfffyyywwwgggppbbvvkxqjz"

type Config struct {
	Words      []string
	Backwards  bool
	Diagonals  bool
	StopFill   bool
	Capitalize bool
	Width      int
	Height     int
}

func createNewGrid(width, height int) [][]rune {
	newGrid := make([][]rune, height)

	for i := range newGrid {
		newGrid[i] = make([]rune, width)
	}

	return newGrid
}

func addWordsToGrid(width, height int, words []string, diagonals, backwards bool) ([][]rune, []string, []string) {
	grid := createNewGrid(width, height)
	placed := []string{}
	failed := []string{}

	// Sort the words by length
	sort.SliceStable(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})

	// Loop over the words
	for {
		outcome := false

		outcome, grid = tryFindPlaceOnGrid(grid, words[0], diagonals, backwards)

		if outcome {
			placed = append(placed, words[0])
		} else {
			failed = append(failed, words[0])
		}

		words = words[1:]

		if len(words) == 0 {
			break
		}
	}

	return grid, placed, failed
}

func tryFindPlaceOnGrid(grid [][]rune, word string, diagonals, backwards bool) (bool, [][]rune) {
	localDirections := setupPossibleDirections(diagonals, backwards)

	for {
		positions := setupPossiblePositions(grid)

		// Loop over positions
		for {
			row := positions[0] / len(grid[0])
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

		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
			return false, grid
		}
	}

	return true, cloneGrid
}

func setupPossibleDirections(diagonals, backwards bool) [][2]int {
	directions := [][2]int{}

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
	pos := []int{}
	size := len(grid) * len(grid[0])

	for i := 0; i < size; i++ {
		pos = append(pos, i)
	}

	rand.Shuffle(len(pos), func(i, j int) { pos[i], pos[j] = pos[j], pos[i] })

	return pos
}

func fillGrid(grid [][]rune, stopFill bool) [][]rune {
	chars := []rune(letterFrequency)
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

func Generate(myConfig Config) ([][]rune, []string, []string) {
	grid, placed, failed := addWordsToGrid(myConfig.Width, myConfig.Height,
		myConfig.Words, myConfig.Diagonals, myConfig.Backwards)

	return fillGrid(grid, myConfig.StopFill), placed, failed
}
