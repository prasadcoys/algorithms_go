package datastructures

func MergeSort(entries []int) []int {

	DoRecursiveMergeSort(&entries, 0, len(entries)-1, 0)
	return entries
}

// export
func DoRecursiveMergeSort(entries *[]int, startIndex int, endIndex int, d int) int {
	leftEndIndex := (startIndex + endIndex) / 2
	rightStartIndex := leftEndIndex + 1
	if startIndex < endIndex {
		d = d + DoRecursiveMergeSort(entries, startIndex, leftEndIndex, 0)
		d = d + DoRecursiveMergeSort(entries, rightStartIndex, endIndex, 0)

	}
	d = d + mergeTwoHalves(entries, startIndex, endIndex, 0)
	return d
}

func mergeTwoHalves(entries *[]int, startIndex int, endIndex int, d int) int {
	if startIndex == endIndex {
		return d
	}
	tempArray := []int{}
	leftEndIndex := (startIndex + endIndex) / 2
	rightStartIndex := leftEndIndex + 1
	i := startIndex
	j := rightStartIndex
	for {
		if i > leftEndIndex || j > endIndex {
			break
		}
		if (*entries)[i] > (*entries)[j] {
			tempArray = append(tempArray, (*entries)[j])
			j = j + 1
			d = d + leftEndIndex + 1 - i

		} else {
			tempArray = append(tempArray, (*entries)[i])
			i = i + 1
		}
	}
	if i <= leftEndIndex {
		tempArray = append(tempArray, (*entries)[i:leftEndIndex+1]...)
	} else {
		tempArray = append(tempArray, (*entries)[j:endIndex+1]...)
	}
	i = startIndex
	for _, entry := range tempArray {
		(*entries)[i] = entry
		i = i + 1
	}
	return d
}
