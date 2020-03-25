package datastructures

func AddNumbersDigitWise(operand_1 LinkedList, operand_2 LinkedList) LinkedList {
	result := LinkedList{}
	currentNode_1 := operand_1.Head
	currentNode_2 := operand_2.Head
	carryOver := 0
	for {
		if currentNode_1 == nil && currentNode_2 == nil {
			if carryOver > 0 {
				result.Append(Node{Data: carryOver})
			}
			break
		}
		sumOfCurrentDigits := 0
		if currentNode_1 != nil {
			sumOfCurrentDigits = sumOfCurrentDigits + currentNode_1.Data
		}
		if currentNode_2 != nil {
			sumOfCurrentDigits = sumOfCurrentDigits + currentNode_2.Data
		}
		sumOfCurrentDigits = sumOfCurrentDigits + carryOver
		carryOver = sumOfCurrentDigits / 10
		sumOfCurrentDigits = sumOfCurrentDigits % 10
		result.Append(Node{Data: sumOfCurrentDigits})
		if currentNode_1 != nil {
			currentNode_1 = currentNode_1.Next
		}
		if currentNode_2 != nil {
			currentNode_2 = currentNode_2.Next
		}

	}
	return result
}
