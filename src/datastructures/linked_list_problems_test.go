package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfSingleDigitNumbersWithoutCarryInLLsCanBeAdded(t *testing.T) {
	operand_1 := LinkedList{}
	operand_1.Append(Node{Data: 1})
	operand_2 := LinkedList{}
	operand_2.Append(Node{Data: 2})
	resutl := LinkedList{}
	resutl.Append(Node{Data: 3})
	assert.Equal(t, resutl, AddNumbersDigitWise(operand_1, operand_2))
}

func TestIfSingleDigitNumbersWithCarryInLLsCanBeAdded(t *testing.T) {
	operand_1 := LinkedList{}
	operand_1.Append(Node{Data: 9})
	operand_2 := LinkedList{}
	operand_2.Append(Node{Data: 8})
	resutl := LinkedList{}
	resutl.Append(Node{Data: 7})
	resutl.Append(Node{Data: 1})
	assert.Equal(t, resutl, AddNumbersDigitWise(operand_1, operand_2))
}

func TestIfTwoDigitNumbersWithoutCarryAreAddedCorrectly(t *testing.T) {
	operand_1 := LinkedList{}
	operand_1.Append(Node{Data: 6})
	operand_1.Append(Node{Data: 6})
	operand_2 := LinkedList{}
	operand_2.Append(Node{Data: 1})
	operand_2.Append(Node{Data: 1})
	resutl := LinkedList{}
	resutl.Append(Node{Data: 7})
	resutl.Append(Node{Data: 7})
	assert.Equal(t, resutl, AddNumbersDigitWise(operand_1, operand_2))
}

func TestIfTwoDigitNumbersWithCarryAreAddedCorrectly(t *testing.T) {
	operand_1 := LinkedList{}
	operand_1.Append(Node{Data: 6})
	operand_1.Append(Node{Data: 6})
	operand_2 := LinkedList{}
	operand_2.Append(Node{Data: 4})
	operand_2.Append(Node{Data: 1})
	resutl := LinkedList{}
	resutl.Append(Node{Data: 0})
	resutl.Append(Node{Data: 8})
	assert.Equal(t, resutl, AddNumbersDigitWise(operand_1, operand_2))
}

func TestIfThreeDigtisCanBeAdded(t *testing.T) {
	operand_1 := LinkedList{}
	operand_1.Append(Node{Data: 2})
	operand_1.Append(Node{Data: 4})
	operand_1.Append(Node{Data: 3})
	operand_2 := LinkedList{}
	operand_2.Append(Node{Data: 5})
	operand_2.Append(Node{Data: 6})
	operand_2.Append(Node{Data: 4})
	resutl := LinkedList{}
	resutl.Append(Node{Data: 7})
	resutl.Append(Node{Data: 0})
	resutl.Append(Node{Data: 8})
	assert.Equal(t, resutl, AddNumbersDigitWise(operand_1, operand_2))
}

func TestIfTwoNumbersWithDifferentDigitsCarryAreAddedCorrectly(t *testing.T) {
	operand_1 := LinkedList{}
	operand_1.Append(Node{Data: 9})
	operand_1.Append(Node{Data: 8})
	operand_2 := LinkedList{}
	operand_2.Append(Node{Data: 1})
	resutl := LinkedList{}
	resutl.Append(Node{Data: 0})
	resutl.Append(Node{Data: 9})
	assert.Equal(t, resutl, AddNumbersDigitWise(operand_1, operand_2))
}
