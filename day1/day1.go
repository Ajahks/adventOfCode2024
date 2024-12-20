package day1

import (
    "fmt"
    "math"
)

func RunDay1Problem() {
    inputList1 := []int{2, 1, 5, 6, 3, 12}
    inputList2 := []int{0, 1, 2}
    fmt.Println("Input list1:", inputList1)
    fmt.Println("Input list2:", inputList2)

    totalDistance := findTotalDistance(inputList1, inputList2)
    fmt.Println("Solution 1:", totalDistance)
}

func findTotalDistance(list1 []int, list2 []int) int {
    sortedList1 := mergeSort(list1)
    sortedList2 := mergeSort(list2)
    fmt.Println("Sorted list1:", sortedList1)
    fmt.Println("Sorted list2:", sortedList2)

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

func mergeSort(list []int) []int {
    if len(list) <= 1 {
        return list
    }
    if len(list) == 2 && list[0] < list[1] {
        return list
    }

    firstHalf := list[0 : len(list)/2]
    secondHalf := list[len(list)/2:]

    sortedFirstHalf := mergeSort(firstHalf)
    sortedSecondHalf := mergeSort(secondHalf)

    i := 0
    j := 0
    sortedList := make([]int, 0)
    for i < len(sortedFirstHalf) && j < len(sortedSecondHalf) {
        if sortedFirstHalf[i] < sortedSecondHalf[j] {
            sortedList = append(sortedList, sortedFirstHalf[i])
            i++
        } else {
            sortedList = append(sortedList, sortedSecondHalf[j])
            j++
        }
    }

    if i < len(sortedFirstHalf) {
        sortedList = append(sortedList, sortedFirstHalf[i:]...)
    }
    if j < len(sortedSecondHalf) {
        sortedList = append(sortedList, sortedSecondHalf[j:]...)
    }

    return sortedList
}
