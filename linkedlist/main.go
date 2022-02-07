package main

import "fmt"

// Node is the basic structure of the Linked List
type Node struct {
	value    int
	nextNode *Node
}

// LinkedList is the data structure
type LinkedList struct {
	headNode *Node
}

// Add inserts a value to the Linked list
func (linkedList *LinkedList) Add(value int) {
	// initialize the Node
	var node = Node{}
	// set the value into node
	node.value = value

	// check if nextNode value is empty
	if node.nextNode != nil {
		// set the nextNode as the headNode
		node.nextNode = linkedList.headNode
	}
	// set headNode as the pointer to the new node created.
	linkedList.headNode = &node
}

// AddAfter adds a node after a specific node value
func (linkedList *LinkedList) AddAfter(nodeValue int, value int) {
	var node = &Node{}
	node.value = value
	node.nextNode = nil
	var newNode *Node
	newNode = linkedList.Find(nodeValue)
	if newNode != nil {
		node.nextNode = newNode.nextNode
		newNode.nextNode = node
	}
}

// List prints all the values in a Linked List
func (linkedList *LinkedList) List() {
	var node *Node
	// loops until current node is no longer equal to nil
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Println(node.value)
	}
}

// Last returns the last node in a Linked List
func (linkedList *LinkedList) Last() *Node {
	var node *Node
	var lastNode *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.nextNode == nil {
			lastNode = node
		}
	}
	return lastNode
}

// Tail adds the value to the tail of the linkedList
func (linkedList *LinkedList) Tail(value int) {
	var node = &Node{}
	node.value = value
	node.nextNode = nil
	var lastNode *Node
	lastNode = linkedList.Last()
	if lastNode != nil {
		lastNode.nextNode = node
	}
}

// Find returns the Node of a linked list with the value input
func (linkedList *LinkedList) Find(value int) *Node {
	var node *Node
	var found *Node
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		if node.value == value {
			found = node
			break
		}
	}
	return found
}

func main() {
	var linkedList LinkedList
	linkedList = LinkedList{}
	linkedList.Add(1)
	linkedList.Tail(3)
	linkedList.AddAfter(1, 7)
	linkedList.List()
}
