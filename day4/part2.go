package day4

import (
	"adventOfCode2024/util"
	"fmt"
)

func runPart2Problem() {
	fmt.Println("Problem 2 Start:")
	inputs, err := util.ReadInputIntoMultipleStrings("input/day4/day4input.txt")
	if err != nil {
		fmt.Println("Error parsing input: ", err)
		return
	}

	result := findXMas(inputs)
	fmt.Println("X-MAS count: ", result)
}

func findXMas(wordSearch []string) int {
	totalCount := 0
	for rowIndex, row := range wordSearch {
		if rowIndex >= len(wordSearch) {
			break
		}

		for colIndex, _ := range row {
			if colIndex >= len(wordSearch[rowIndex]) {
				break
			}
			totalCount += checkXMases(colIndex, rowIndex, wordSearch)
		}
	}

	return totalCount
}

func checkXMases(colIndex int, rowIndex int, wordSearch []string) int {
	if colIndex < len(wordSearch[rowIndex])-2 && rowIndex < len(wordSearch)-2 {
		diag1 := string(wordSearch[rowIndex][colIndex]) + string(wordSearch[rowIndex+1][colIndex+1]) + string(wordSearch[rowIndex+2][colIndex+2])
		diag2 := string(wordSearch[rowIndex][colIndex+2]) + string(wordSearch[rowIndex+1][colIndex+1]) + string(wordSearch[rowIndex+2][colIndex])
		return countXMas(diag1, diag2)
	}
	return 0
}

func countXMas(diag1 string, diag2 string) int {
	if (diag1 == "MAS" || diag1 == "SAM") && (diag2 == "MAS" || diag2 == "SAM") {
		return 1
	}

	return 0
}
