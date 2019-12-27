package datastructures

type DynamicArray struct {
	elements []int
}

// Size gives the current number of elements in the dynamic array
func (d *DynamicArray) Size() int {
	return len(d.elements)
}

func (d *DynamicArray) IsEmpty() bool {
	return d.Size() == 0
}

func (d *DynamicArray) Get(index int) int {
	return d.elements[index]
}
