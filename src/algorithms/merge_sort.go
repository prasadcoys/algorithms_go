package algorithms

func MergeSort(entries []int) []int {
	doRecursiveMergeSort(&entries, 0, len(entries)-1)
	return entries
}

func doRecursiveMergeSort(entries *[]int, startIndex int, endIndex int) {
	if startIndex < endIndex {
		leftEndIndex := (startIndex + endIndex) / 2
		rightStartIndex := leftEndIndex + 1
		doRecursiveMergeSort(entries, startIndex, leftEndIndex)
		doRecursiveMergeSort(entries, rightStartIndex, endIndex)
	}

}
