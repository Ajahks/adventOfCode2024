package day3

import (
    "adventOfCode2024/util"
    "errors"
    "fmt"
    "strconv"
)

func runPart1Problem() {
    fmt.Println("Problem 1 Start:")
    input, err := util.ReadInputIntoSingleString("input/day3/day3input")
    if err != nil {
        fmt.Println("Error parsing input: ", err)
        return
    }
    result := evalValidMuls(input)
    fmt.Println("Solution: ", result)
}

func evalValidMuls(inputString string) int {
    valMulSum := 0
    i := 0
    for i < len(inputString)-7 {
        if inputString[i:i+4] == "mul(" {
            i += 4
            result, endIndex, err := evalValidMulNoPrefix(inputString, i)
            if err == nil {
                valMulSum += result
            }
            i = endIndex
            continue
        }
        i++
    }
    return valMulSum
}

// returns the mul result if valid, the index of last character of the mul function, and error if not valid
func evalValidMulNoPrefix(inputString string, startIndex int) (int, int, error) {
    i := startIndex
    // Check first number
    firstNumber, endIndex, err := evaluateValidNumber(inputString, i)
    if err != nil {
        return 0, endIndex, err
    }
    i = endIndex

    // Check comma
    if inputString[i] != ',' {
        return 0, i, errors.New("no comma after first number")
    }
    i++

    // Check second number
    secondNumber, endIndex, err := evaluateValidNumber(inputString, i)
    if err != nil {
        return 0, endIndex, err
    }
    i = endIndex

    // Check end parenthesis
    var myVal = inputString[i]
    if myVal != ')' {
        return 0, i, errors.New("no parenthesis after second number")
    }
    i++

    // All good multiply and return
    return firstNumber * secondNumber, i, nil
}

// Check if the next line of text is a valid number of 1-3 digits
func evaluateValidNumber(inputString string, startIndex int) (int, int, error) {
    if startIndex+4 <= len(inputString) {
        _, err := strconv.Atoi(inputString[startIndex : startIndex+4])
        if err == nil {
            return 0, startIndex, errors.New("number too long")
        }
    }

    if startIndex+3 <= len(inputString) {
        value, err := strconv.Atoi(inputString[startIndex : startIndex+3])
        if err == nil {
            return value, startIndex + 3, nil
        }
    }

    if startIndex+2 <= len(inputString) {
        value, err := strconv.Atoi(inputString[startIndex : startIndex+2])
        if err == nil {
            return value, startIndex + 2, nil
        }
    }

    value, err := strconv.Atoi(string(inputString[startIndex]))
    if err == nil {
        return value, startIndex + 1, nil
    }

    return 0, startIndex, errors.New("invalid number")
}
