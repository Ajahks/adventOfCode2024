package day2

import (
	"adventOfCode2024/util"
	"fmt"
	"math"
	"slices"
)

func runPart2Problem() {
	fmt.Println("Problem 2 Start:")
	inputs, err := util.ReadInputListsHorizontal("input/day2/day2input.txt")
	if err != nil {
		fmt.Println("Error parsing input: ", err)
		return
	}

	fmt.Println("Solution:", countSafeReportsAllow1Error(inputs))
}

func countSafeReportsAllow1Error(reports [][]int) int {
	safeReportsCount := 0
	for _, report := range reports {
		if isReportSafeWith1ErrorMax(report) {
			safeReportsCount++
		}
	}
	return safeReportsCount
}

func isReportSafeWith1ErrorMax(report []int) bool {
	isSafe, unsafeIndex := isReportSafe(report)

	if isSafe {
		return true
	}

	// edge case - see if it works if the first level is removed - first level is key in determining the order
	// post mortem, i should have separated the check for order and scale, might have had some optimizations there
	if unsafeIndex != 0 {
		reportWithFirstIndexRemoved := slices.Concat(report[1:])
		isSafeFirstRemoved, _ := isReportSafe(reportWithFirstIndexRemoved)
		if isSafeFirstRemoved {
			return true
		}
	}

	// check if the report is safe with either the left level or right level removed
	reportWithBadLeftIndexRemoved := slices.Concat(report[:unsafeIndex], report[unsafeIndex+1:])
	reportWithBadRightIndexRemoved := slices.Concat(report[:unsafeIndex+1], report[unsafeIndex+2:])
	isSafeLeftRemoved, _ := isReportSafe(reportWithBadLeftIndexRemoved)
	isSafeRightRemoved, _ := isReportSafe(reportWithBadRightIndexRemoved)
	if isSafeLeftRemoved || isSafeRightRemoved {
		return true
	}
	return false
}

// Basically my part 1 solution
// returns
// bool - whether the report is safe
// index of left unsafe level. -1 if safe
func isReportSafe(report []int) (bool, int) {
	i := 0
	j := 1
	if j >= len(report) {
		return true, -1
	}
	ascending := report[i] < report[j]
	safe := true
	leftUnsafeLevelIndex := -1

	for j < len(report) {
		if ascending && report[i] >= report[j] {
			safe = false
			leftUnsafeLevelIndex = i
			break
		} else if !ascending && report[i] <= report[j] {
			safe = false
			leftUnsafeLevelIndex = i
			break
		}
		if math.Abs(float64(report[i]-report[j])) > 3 {
			safe = false
			leftUnsafeLevelIndex = i
			break
		}
		i++
		j++
	}

	return safe, leftUnsafeLevelIndex
}
