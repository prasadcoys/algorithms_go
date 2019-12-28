package datastructures

type DoubleLinkedNode struct {
	Data int
	Prev *DoubleLinkedNode
	Next *DoubleLinkedNode
}

func CreateDoubleLinkNode(data int) DoubleLinkedNode {
	return DoubleLinkedNode{
		Data: data,
		Prev: nil,
		Next: nil,
	}
}
