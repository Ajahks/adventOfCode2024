package day1

import (
    "adventOfCode2024/util"
    "fmt"
    "math"
)

func runPart1Problem() {
    fmt.Println("Problem 1 Start:")
    inputs, err := util.ReadInputLists("input/day1/part1input")
    if err != nil {
        fmt.Println("Error parsing input: ", err)
        return
    }
    inputList1 := inputs[0]
    inputList2 := inputs[1]

    totalDistance := findTotalDistance(inputList1, inputList2)
    fmt.Println("Solution:", totalDistance)
}

func findTotalDistance(list1 []int, list2 []int) int {
    sortedList1 := util.MergeSort(list1)
    sortedList2 := util.MergeSort(list2)

    totalDistance := 0
    i := 0
    j := 0
    for i < len(sortedList1) && j < len(sortedList2) {
        distance := math.Abs(float64(sortedList1[i] - sortedList2[j]))
        totalDistance += int(distance)
        i++
        j++
    }

    for i < len(sortedList1) {
        totalDistance += sortedList1[i]
        i++
    }
    for j < len(sortedList2) {
        totalDistance += sortedList2[j]
        j++
    }

    return totalDistance
}
