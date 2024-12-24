package day4

import (
	"adventOfCode2024/util"
	"fmt"
)

func runPart1Problem() {
	fmt.Println("Problem 1 Start:")
	inputs, err := util.ReadInputIntoMultipleStrings("input/day4/day4input.txt")
	if err != nil {
		fmt.Println("Error parsing input: ", err)
		return
	}

	result := findXmas(inputs)
	fmt.Println("Xmas count: ", result)
}

func findXmas(wordSearch []string) int {
	totalCount := 0
	for rowIndex, row := range wordSearch {
		if rowIndex >= len(wordSearch) {
			break
		}

		for colIndex, _ := range row {
			if colIndex >= len(wordSearch[rowIndex]) {
				break
			}
			totalCount += checkXmases(colIndex, rowIndex, wordSearch)
		}
	}

	return totalCount
}

func checkXmases(colIndex int, rowIndex int, wordSearch []string) int {
	potentialXmas := []string{}
	if colIndex < len(wordSearch[rowIndex])-3 {
		row1 := wordSearch[rowIndex][colIndex : colIndex+4]
		potentialXmas = append(potentialXmas, row1)
	}
	if rowIndex < len(wordSearch)-3 {
		col1 := string(wordSearch[rowIndex][colIndex]) + string(wordSearch[rowIndex+1][colIndex]) + string(wordSearch[rowIndex+2][colIndex]) + string(wordSearch[rowIndex+3][colIndex])
		potentialXmas = append(potentialXmas, col1)
	}
	if colIndex < len(wordSearch[rowIndex])-3 && rowIndex < len(wordSearch)-3 {
		diag1 := string(wordSearch[rowIndex][colIndex]) + string(wordSearch[rowIndex+1][colIndex+1]) + string(wordSearch[rowIndex+2][colIndex+2]) + string(wordSearch[rowIndex+3][colIndex+3])
		diag2 := string(wordSearch[rowIndex][colIndex+3]) + string(wordSearch[rowIndex+1][colIndex+2]) + string(wordSearch[rowIndex+2][colIndex+1]) + string(wordSearch[rowIndex+3][colIndex])
		potentialXmas = append(potentialXmas, diag1)
		potentialXmas = append(potentialXmas, diag2)
	}

	return countXmas(potentialXmas)
}

func countXmas(potentialXmas []string) int {
	count := 0
	for _, str := range potentialXmas {
		if str == "XMAS" || str == "SAMX" {
			count++
		}
	}
	return count
}
