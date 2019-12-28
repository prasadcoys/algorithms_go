package datastructures

import (
	"errors"
)

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

func (l *LinkedList) AddAt(index int, data int) error {
	currentNode := l.Head
	nextNode := currentNode.Next
	currentPosition := 1
	if index < 0 {
		return errors.New("Index out of bounds")
	}
	for {
		if currentNode.Next == nil {
			return errors.New("Index out of bounds")
		}
		if currentPosition+1 == index {
			nodeToBeInserted := CreateNode(data)
			currentNode.Next = &nodeToBeInserted
			nodeToBeInserted.Next = nextNode
			return nil
		}

		currentNode = currentNode.Next
		nextNode = currentNode.Next
		currentPosition++
	}
	return nil
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

func (l *LinkedList) PeekFirst() (int, error) {
	if l.Head != nil {
		return l.Head.Data, nil
	}
	return -1, errors.New("Head not set")
}

func (l *LinkedList) RemoveAt(index int) error {
	if l.Head == nil {
		return errors.New("Empty list")
	}
	if index == 1 {
		l.Head = l.Head.Next
		return nil
	}
	currentNode := l.Head
	nextNode := l.Head.Next
	currentPosition := 1
	for {
		if currentPosition+1 == index {
			currentNode.Next = nextNode.Next
			nextNode = nil
			break
		}
		currentNode = currentNode.Next
		nextNode = currentNode.Next
		currentPosition++
	}
	return nil
}

func (l *LinkedList) RemoveFirst() error {

	return l.RemoveAt(1)

}

func (l *LinkedList) RemoveLast() error {
	return l.RemoveAt(l.Size())
}
