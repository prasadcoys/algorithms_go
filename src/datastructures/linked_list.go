package datastructures

type LinkedList struct {
	Head *Node
}

func (l *LinkedList) Append(node Node) {
	currentNode := l.Head
	if currentNode == nil {
		l.Head = &node
		return
	}
	for {
		if currentNode.Next == nil {
			currentNode.Next = &node
			node.Next = nil
			return
		}
		currentNode = currentNode.Next
	}
}

func (l *LinkedList) Size() int {
	if l.Head == nil {
		return 0
	}
	currentNode := l.Head
	numberOfElements := 1
	for {
		if currentNode.Next == nil {
			return numberOfElements
		}
		currentNode = currentNode.Next
		numberOfElements++
	}
}

func (l *LinkedList) IsEmpty() bool {
	if l.Size() == 0 {
		return true
	}
	return false
}

func (l *LinkedList) AddAt(index int, data int) {
	currentNode := l.Head
	nextNode := currentNode.Next
	currentPosition := 1
	for {
		if currentPosition+1 == index {
			nodeToBeInserted := CreateNode(data)
			currentNode.Next = &nodeToBeInserted
			nodeToBeInserted.Next = nextNode
			return
		}
		currentNode = currentNode.Next
		nextNode = currentNode.Next
	}
}

func (l *LinkedList) Clear() {
	l.Head = nil
}

func CreateNode(data int) Node {
	return Node{
		Data: data,
		Next: nil,
	}
}
