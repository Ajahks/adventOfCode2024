package util

import (
    "bufio"
    "os"
    "strconv"
    "strings"
)

// ReadInputListsVertical - Reads input lists in format of vertical (columns) of lists
func ReadInputListsVertical(filepath string) ([][]int, error) {
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

// ReadInputListsHorizontal - Reads input lists in format of horizontal (rows) of lists
func ReadInputListsHorizontal(filepath string) ([][]int, error) {
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
        numList := make([]int, len(values))
        for i, value := range values {
            numList[i], _ = strconv.Atoi(value)
        }
        lists = append(lists, numList)
    }
    return lists, nil
}

func ReadInputIntoSingleString(filepath string) (string, error) {
    file, err := os.ReadFile(filepath)
    if err != nil {
        return "", err
    }

    return string(file), nil
}
