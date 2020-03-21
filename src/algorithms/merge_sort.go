package algorithms

import "fmt"

func MergeSort(entries []int) []int {
	d := 0
	doRecursiveMergeSort(&entries, 0, len(entries)-1, &d)
	return entries
}

func doRecursiveMergeSort(entries *[]int, startIndex int, endIndex int, d *int) *int {
	leftEndIndex := (startIndex + endIndex) / 2
	rightStartIndex := leftEndIndex + 1
	if startIndex < endIndex {
		*d = *d + *doRecursiveMergeSort(entries, startIndex, leftEndIndex, d)
		*d = *d + *doRecursiveMergeSort(entries, rightStartIndex, endIndex, d)

	}
	*d = *d + *mergeTwoHalves(entries, startIndex, endIndex, d)
	return d
}

func mergeTwoHalves(entries *[]int, startIndex int, endIndex int, d *int) *int {
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
			fmt.Println("item of interest", j-i, (*entries)[j])
			j = j + 1
			*d = *d + 1

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
