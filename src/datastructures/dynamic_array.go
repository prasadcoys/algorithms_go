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

func (d *DynamicArray) SetAt(index int, element int) {
	d.elements[index] = element
}

func (d *DynamicArray) Add(element int) {
	d.elements = append(d.elements, element)
}

func (d *DynamicArray) RemoveAt(index int) {
	copy(d.elements[index:], d.elements[index+1:])
	d.elements = d.elements[:len(d.elements)-1]
}
