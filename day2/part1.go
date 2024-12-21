package day2

import (
	"adventOfCode2024/util"
	"fmt"
	"math"
)

func runPart1Problem() {
	fmt.Println("Problem 1 Start:")
	inputs, err := util.ReadInputListsHorizontal("input/day2/day2input.txt")
	if err != nil {
		fmt.Println("Error parsing input: ", err)
		return
	}

	fmt.Println("Solution:", countSafeReports(inputs))
}

func countSafeReports(reports [][]int) int {
	safeReportsCount := 0
	for _, report := range reports {
		i := 0
		j := 1
		if j >= len(report) {
			safeReportsCount++
			continue
		}
		ascending := report[i] < report[j]
		safe := true

		for j < len(report) {
			if ascending && report[i] >= report[j] {
				safe = false
				break
			} else if !ascending && report[i] <= report[j] {
				safe = false
				break
			}
			if math.Abs(float64(report[i]-report[j])) > 3 {
				safe = false
				break
			}
			i++
			j++
		}

		if safe {
			safeReportsCount++
		}
	}
	return safeReportsCount
}
