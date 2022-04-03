package main

import (
	"fmt"
)

type SinglyLinkedListNode struct {
	data int32
	next *SinglyLinkedListNode
}

type DoublyLinkedListNode struct {
	data int32
	next *DoublyLinkedListNode
	prev *DoublyLinkedListNode
}

type SinglyLinkedList struct {
	head *SinglyLinkedListNode
	tail *SinglyLinkedListNode
}

func (singlyLinkedList *SinglyLinkedList) insertNodeIntoSinglyLinkedList(nodeData int32) {
	node := &SinglyLinkedListNode{
		next: nil,
		data: nodeData,
	}

	if singlyLinkedList.head == nil {
		singlyLinkedList.head = node
	} else {
		singlyLinkedList.tail.next = node
	}

	singlyLinkedList.tail = node
}

func printSinglyLinkedList(node *SinglyLinkedListNode, sep string) {
	for node != nil {
		fmt.Printf("%d", node.data)

		node = node.next

		if node != nil {
			fmt.Printf(sep)
		}
	}
}

func insertNodeAtPosition(head *SinglyLinkedListNode, data int32, position int32) *SinglyLinkedListNode {
	if head == nil {
		head = &SinglyLinkedListNode{data: data, next: nil}
		return head
	}

	if position == 1 {
		next := head.next
		head = &SinglyLinkedListNode{data: data, next: next}
		return head
	}

	node := head
	for i := int32(0); i < position-1; i++ {
		if node.next != nil {
			node = node.next
		}
	}

	next := node.next
	node.next = &SinglyLinkedListNode{data: data, next: next}

	return head
}

func sortedInsert(head *DoublyLinkedListNode, data int32) *DoublyLinkedListNode {
	// Write your code here
	if head == nil {
		head = &DoublyLinkedListNode{data: data, next: nil, prev: nil}
		return head
	}

	if data < head.data {
		newNode := &DoublyLinkedListNode{data: data, prev: nil, next: head}
		head.prev = newNode
		return newNode
	}

	node := head
	for node != nil {
		if data <= node.data {
			newNode := &DoublyLinkedListNode{data: data, prev: nil, next: node}

			nxtNode := node.next
			prevNode := node.prev

			nxtNode.prev = newNode
			prevNode.next = newNode

			break
		}
		if node.next == nil {
			node.next = &DoublyLinkedListNode{data: data, prev: node, next: nil}
			break
		}
		node = node.next
	}
	return head
}

func main() {
	head1 := &SinglyLinkedListNode{data: 2,
		next: &SinglyLinkedListNode{data: 8,
			next: &SinglyLinkedListNode{data: 4, next: &SinglyLinkedListNode{data: 7, next: nil}}}}

	head2 := &SinglyLinkedListNode{data: 5,
		next: &SinglyLinkedListNode{data: 3,
			next: &SinglyLinkedListNode{data: 2, next: &SinglyLinkedListNode{data: 4, next: nil}}}}

	//l2 := SinglyLinkedList{head: head1, tail: nil}
	//l1 := SinglyLinkedList{head: head2, tail: nil}

	printSinglyLinkedList(head1, "->")
	printSinglyLinkedList(head2, "->")
	//insertNodeAtPosition(head, 1, 2)

	n1 := DoublyLinkedListNode{prev: nil, data: 1}
	n2 := DoublyLinkedListNode{prev: &n1, data: 4, next: nil}
	n1.next = &n2
	n2.prev = &n1

	n3 := DoublyLinkedListNode{prev: &n2, data: 6, next: nil}
	n2.next = &n3
	n4 := DoublyLinkedListNode{prev: &n3, data: 7, next: nil}
	n3.next = &n4

	sortedInsert(&n1, 5)
}
