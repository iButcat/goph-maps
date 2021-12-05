package internal

import (
	"fmt"
)

type node struct {
	value *Vertice
	next  *node
}

type queue struct {
	length int
	head   *node
	tail   *node
}

func (q *queue) enqueue(v *Vertice) {
	node := &node{value: v}

	if q.head == nil {
		q.head = node
		q.tail = node
		q.length++
		return
	}

	q.tail.next = node
	q.tail = node
	q.length++
}

func (q *queue) dequeue() *Vertice {
	node := q.head
	if node == nil {
		return nil
	}

	q.head = q.head.next

	if q.head == nil {
		q.tail = nil
	}

	return node.value
}

func (q *queue) isEmpty() bool {
	return q.tail == nil
}

func (q *queue) len() int {
	return q.length
}

func (q *queue) Print() {
	if q.isEmpty() {
		fmt.Println("queue is empty")
		return
	}
	current := q.head
	for current != nil {
		fmt.Println("node printing: ", current.value)
		current = current.next
	}
}
