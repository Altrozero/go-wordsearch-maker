package output

import (
	"fmt"
	"strings"
)

func ConsolePrintGrid(grid [][]rune, capitalize bool, placed, failed []string) {
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
