package util

func MergeSort(list []int) []int {
	if len(list) <= 1 {
		return list
	}
	if len(list) == 2 && list[0] < list[1] {
		return list
	}

	firstHalf := list[0 : len(list)/2]
	secondHalf := list[len(list)/2:]

	sortedFirstHalf := MergeSort(firstHalf)
	sortedSecondHalf := MergeSort(secondHalf)

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
