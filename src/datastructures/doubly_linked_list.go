package datastructures

import (
	"errors"
	"strconv"
	"strings"
)

type DoubleLinkedNode struct {
	Data int
	Prev *DoubleLinkedNode
	Next *DoubleLinkedNode
}

type DoublyLinkedList struct {
	Head *DoubleLinkedNode
	Tail *DoubleLinkedNode
	size int
}

func CreateDoubleLinkNode(data int) DoubleLinkedNode {
	return DoubleLinkedNode{
		Data: data,
		Prev: nil,
		Next: nil,
	}
}

func CreateDoublyLinkedList() DoublyLinkedList {
	return DoublyLinkedList{
		Head: nil,
		Tail: nil,
	}
}

func (d *DoublyLinkedList) Append(data int) {
	node := CreateDoubleLinkNode(data)
	if d.Head == nil {
		d.Head = &node
		d.Tail = &node
		d.size++
		return
	} else {
		node.Prev = d.Tail
		d.Tail.Next = &node
		d.Tail = &node
		d.size++
	}
}

func (d *DoublyLinkedList) ToString() string {
	currentNode := d.Head
	printableString := ""
	for {
		if currentNode == nil {
			return strings.Trim(printableString, " ")
		}
		printableString += strconv.Itoa(currentNode.Data) + " "
		currentNode = currentNode.Next

	}
}

func (d *DoublyLinkedList) Size() int {
	return d.size
}

func (d *DoublyLinkedList) IsEmpty() bool {
	return d.Size() == 0
}

func (d *DoublyLinkedList) Clear() {
	d.Head = nil
	d.Tail = nil
	d.size = 0
}

func (d *DoublyLinkedList) AddAt(index int, data int) {
	nodeToBeInserted := CreateDoubleLinkNode(data)
	if index <= d.Size()/2 {
		//search from the head
		currentNode := d.Head
		currentPosition := 1
		for {
			if currentNode == nil {
				break
			}
			if currentPosition+1 == index {
				currentNode.Next.Prev = &nodeToBeInserted
				nodeToBeInserted.Next = currentNode.Next
				currentNode.Next = &nodeToBeInserted
				nodeToBeInserted.Prev = currentNode
				d.size++
			}
			currentNode = currentNode.Next
			currentPosition++
		}
	} else {
		// search from the end
		currentNode := d.Tail
		currentPosition := d.Size()
		for {
			if currentNode == nil {
				break
			}
			if currentPosition == index {
				previousNode := currentNode.Prev
				previousNode.Next = &nodeToBeInserted
				nodeToBeInserted.Prev = previousNode
				nodeToBeInserted.Next = currentNode
				currentNode.Prev = &nodeToBeInserted
				d.size++
			}
			currentNode = currentNode.Prev
			currentPosition--
		}
	}

}

func (d *DoublyLinkedList) PeekFirst() int {
	return d.Head.Data
}

func (d *DoublyLinkedList) PeekLast() int {
	return d.Tail.Data
}

func (d *DoublyLinkedList) RemoveAt(index int) error {

	if index <= d.Size()/2 {
		currentNode := d.Head
		currentPosition := 1
		for {
			if currentNode == nil {
				return errors.New("Index out of bounds")
			}
			if currentPosition+1 == index {
				currentNode.Next = currentNode.Next.Next
				if currentNode.Next.Next != nil {
					currentNode.Next.Next.Prev = currentNode
				}
				d.size--
				return nil
			}
			currentNode = currentNode.Next
			currentPosition++
		}
	} else {
		currentNode := d.Tail
		currentPosition := d.Size()
		for {
			if currentNode == nil {
				return errors.New("Index out of bounds")
			}
			if currentPosition == index {
				currentNode.Prev.Next = currentNode.Next
				currentNode.Next.Prev = currentNode.Prev
				d.size--
				return nil
			}
			currentNode = currentNode.Prev
			currentPosition--
		}
	}

}

func (d *DoublyLinkedList) RemoveFirst() error {
	if d.Head == nil {
		return errors.New("List not initialised")
	}
	d.Head = d.Head.Next
	d.Head.Prev = nil
	d.size--
	return nil
}

func (d *DoublyLinkedList) RemoveLast() error {
	if d.Tail == nil {
		return errors.New("List not initialised")
	}
	d.Tail = d.Tail.Prev
	d.Tail.Next = nil
	d.size--
	return nil
}

func (d *DoublyLinkedList) IndexOf(data int) int {
	if d.Head == nil {
		return -1
	}
	currentNode := d.Head
	currentIndex := 1
	for {
		if currentNode == nil {
			break
		}
		if currentNode.Data == data {
			return currentIndex
		}
		currentNode = currentNode.Next
		currentIndex++
	}

	return -1
}

func (d *DoublyLinkedList) Contains(data int) bool {
	return d.IndexOf(data) != -1
}
