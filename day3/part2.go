package day3

import (
    "adventOfCode2024/util"
    "fmt"
)

func runPart2Problem() {
    fmt.Println("Problem 2 Start:")
    input, err := util.ReadInputIntoSingleString("input/day3/day3input")
    if err != nil {
        fmt.Println("Error parsing input: ", err)
        return
    }
    result := evalValidEnabledMuls(input)
    fmt.Println("Solution: ", result)
}

func evalValidEnabledMuls(inputString string) int {
    valMulSum := 0
    enabled := true
    i := 0
    for i < len(inputString)-7 {
        if enabled {
            if inputString[i:i+4] == "mul(" {
                i += 4
                result, endIndex, err := evalValidMulNoPrefix(inputString, i)
                if err == nil {
                    valMulSum += result
                }
                i = endIndex
                continue
            }
        }
        if inputString[i:i+4] == "do()" {
            enabled = true
            i += 4
            continue
        }
        if inputString[i:i+7] == "don't()" {
            enabled = false
            i += 7
            continue
        }

        i++
    }
    return valMulSum
}
