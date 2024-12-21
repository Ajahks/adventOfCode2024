package day1

import (
	"adventOfCode2024/util"
	"fmt"
)

func runPart2Problem() {
	fmt.Println("Problem 2 Start:")
	inputs, err := util.ReadInputListsVertical("input/day1/part1input")
	if err != nil {
		fmt.Println("Error parsing input: ", err)
		return
	}
	inputList1 := inputs[0]
	inputList2 := inputs[1]

	similarityRating := findSimilarityRating(inputList1, inputList2)
	fmt.Println("Solution: ", similarityRating)
}

func findSimilarityRating(inputList1 []int, inputList2 []int) int {
	list2IdToOccurrenceMap := listToOccurrenceMap(inputList2)
	similarityRating := 0
	for _, id := range inputList1 {
		similarityRating += id * list2IdToOccurrenceMap[id]
	}
	return similarityRating
}

func listToOccurrenceMap(idList []int) map[int]int {
	idToOccurrenceMap := make(map[int]int)
	for _, id := range idList {
		idToOccurrenceMap[id]++
	}

	return idToOccurrenceMap
}
