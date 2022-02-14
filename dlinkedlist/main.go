package main

import "fmt"

// Node structure of a Doubly Linked List
type Node struct {
	value    int
	nextNode *Node
	prevNode *Node
}

// LinkedList structure to initialize the start
type LinkedList struct {
	head *Node
}

// FindBetweenValues traverses the list and finds the node that is between the 2 values
func (linkedList *LinkedList) FindBetweenValues(firstValue, secondValue int) *Node {
	var node *Node
	var found *Node
	for node = linkedList.head; node != nil; node = node.nextNode {
		if node.prevNode != nil && node.nextNode != nil {
			if node.prevNode.value == firstValue && node.nextNode.value == secondValue {
				found = node
				break
			}
		}
	}
	return found
}

// Add value to doubly linked list
func (linkedList *LinkedList) Add(value int) {
	var node = &Node{}
	node.value = value
	node.nextNode = nil
	// if head is not empty add to next
	if linkedList.head != nil {
		node.nextNode = linkedList.head
		linkedList.head.prevNode = node
	}
	// else make value as head
	linkedList.head = node
}

// Find searches for the value inside the list
func (linkedList *LinkedList) Find(value int) *Node {
	var node *Node
	var found *Node
	for node = linkedList.head; node != nil; node = node.nextNode {
		if node.value == value {
			found = node
			break
		}
	}
	return found
}

// AddAfter adds a value at a specific position of the list after a value
func (linkedList *LinkedList) AddAfter(aftVal, value int) {
	var node = &Node{}
	node.value = value
	node.nextNode = nil
	var tempNode *Node
	tempNode = linkedList.Find(aftVal)
	if tempNode != nil {
		// swap positions
		node.nextNode = tempNode.nextNode
		node.prevNode = tempNode
		tempNode.nextNode = node
	}
}

// FindLast finds the position of the last node in the list
func (linkedList *LinkedList) FindLast() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.head; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// AddEnd adds a value to the end of the list
func (linkedList *LinkedList) AddEnd(value int) {
	var node = &Node{}
	node.value = value
	node.nextNode = nil
	var lastNode *Node
	lastNode = linkedList.FindLast()
	if lastNode != nil {
		lastNode.nextNode = node
		node.prevNode = lastNode
	}
}

// List prints all the values in a Linked List
func (linkedList *LinkedList) List() {
	var node *Node
	// loops until current node is no longer equal to nil
	for node = linkedList.head; node != nil; node = node.nextNode {
		fmt.Println(node.value)
	}
}

// driver code
func main() {
	var linkedList LinkedList
	linkedList = LinkedList{}
	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.AddEnd(5)
	linkedList.AddAfter(1, 3)
	fmt.Println(linkedList.head.value)
	var node *Node
	node = linkedList.FindBetweenValues(1, 5)
	fmt.Println(node.value)
	fmt.Println("Listing all values")
	linkedList.List()
}
