package algorithms

import "fmt"

func MergeSort(entries []int) []int {
	doRecursiveMergeSort(&entries, 0, len(entries)-1, 0)
	return entries
}

func doRecursiveMergeSort(entries *[]int, startIndex int, endIndex int, d int) int {
	leftEndIndex := (startIndex + endIndex) / 2
	rightStartIndex := leftEndIndex + 1
	if startIndex < endIndex {
		d += doRecursiveMergeSort(entries, startIndex, leftEndIndex, d)
		d += doRecursiveMergeSort(entries, rightStartIndex, endIndex, d)

	}
	fmt.Println(startIndex, ":", endIndex)
	mergeTwoHalves(entries, startIndex, endIndex)
	return d
}

func mergeTwoHalves(entries *[]int, startIndex int, endIndex int) {
	if startIndex == endIndex {
		return
	}
	tempArray := []int{}
	leftEndIndex := (startIndex + endIndex) / 2
	rightStartIndex := leftEndIndex + 1
	for {
		i := startIndex
		j := rightStartIndex
		for {
			if i > leftEndIndex || j > endIndex {
				fmt.Println("About to break")
				fmt.Println(tempArray)
				return
			}
			if (*entries)[i] > (*entries)[j] {
				fmt.Println("i greater")
				tempArray = append(tempArray, (*entries)[j])
				j++
			} else {
				fmt.Println("j greater")
				tempArray = append(tempArray, (*entries)[i])
				i++
			}
		}

		/* if i <= leftEndIndex {
			for _, entry := range (*entries)[i:leftEndIndex] {
				tempArray = append(tempArray, entry)
			}
		} else {

			for _, entry := range (*entries)[j:endIndex] {
				tempArray = append(tempArray, entry)
			}
		} */

		// fmt.Println(tempArray)

	}
}
