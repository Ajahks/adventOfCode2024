package day1

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
)

func RunDay1Problem() {
    fmt.Println("Advent of Code 2024 - Day 1 Problem 1")
    inputs, err := readInputLists("day1/input/puzzle1Input")
    if err != nil {
        fmt.Println("Error parsing input: ", err)
        return
    }
    inputList1 := inputs[0]
    inputList2 := inputs[1]

    totalDistance := findTotalDistance(inputList1, inputList2)
    fmt.Println("Solution:", totalDistance)
}

/**
 * Reads input lists in format of vertical (columns) of lists
 */
func readInputLists(filepath string) ([][]int, error) {
    file, err := os.Open(filepath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    lists := make([][]int, 0)
    for scanner.Scan() {
        text := scanner.Text()
        values := strings.Fields(text)

        // initialize lists if not already created
        for len(lists) < len(values) {
            lists = append(lists, make([]int, 0))
        }

        for i, value := range values {
            intValue, _ := strconv.Atoi(value)
            lists[i] = append(lists[i], intValue)
        }
    }
    return lists, nil
}

func findTotalDistance(list1 []int, list2 []int) int {
    sortedList1 := mergeSort(list1)
    sortedList2 := mergeSort(list2)

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
