package main

import "testing"

var linkedList LinkedList

func TestAddToHead(t *testing.T) {
	linkedList = LinkedList{}
	linkedList.Add(1)
}

func TestAddToTail(t *testing.T) {
	linkedList = LinkedList{}
	linkedList.Tail(3)
}

func TestAddAfterValue(t *testing.T) {
	linkedList = LinkedList{}
	linkedList.AddAfter(1, 7)
}

func TestListValues(t *testing.T) {
	linkedList = LinkedList{}
	linkedList.List()
}
