// Copyright © 2020 ichenq@outlook.com All rights reserved.
// Distributed under the terms and conditions of the BSD License.
// See accompanying files LICENSE.

package go_queue_benchmark

import (
	"fmt"
	"strings"
)

const defaultInternalSliceSize = 128 // maximum size of each internal slice

type LinkedArrayNode struct {
	array []interface{}
	next  *LinkedArrayNode
}

func NewLinkedArrayNode() *LinkedArrayNode {
	return &LinkedArrayNode{
		array: make([]interface{}, 0, defaultInternalSliceSize),
	}
}

type LinkedArrayQueue struct {
	head      *LinkedArrayNode // first node of the linked list
	tail      *LinkedArrayNode // last node of the linked list
	pos       int              // index pointing to the current first element in the queue
	size      int              // count of queue items
	sliceSize int              //
}

func NewLinkedArrayQueue(sliceSize int) Queue {
	if sliceSize <= 0 {
		sliceSize = defaultInternalSliceSize
	}
	var node = NewLinkedArrayNode()
	return &LinkedArrayQueue{
		head:      node,
		tail:      node,
		sliceSize: sliceSize,
	}
}

func (q LinkedArrayQueue) String() string {
	var sb = &strings.Builder{}
	var node = q.head
	if node != nil {
		for i := q.pos; i < len(node.array); i++ {
			sb.WriteString(fmt.Sprintf("%v ", node.array[i]))
		}
	}
	for ; node != nil; node = node.next {
		for i := 0; i < len(node.array); i++ {
			sb.WriteString(fmt.Sprintf("%v ", node.array[i]))
		}
	}
	return sb.String()
}

func (q *LinkedArrayQueue) Len() int {
	return q.size
}

func (q *LinkedArrayQueue) isEmpty() bool {
	return q.head == nil || q.pos >= len(q.head.array)
}

func (q *LinkedArrayQueue) Front() interface{} {
	if q.size > 0 && !q.isEmpty() {
		return q.head.array[q.pos]
	}
	return nil
}

func (q *LinkedArrayQueue) Enqueue(v interface{}) bool {
	if q.head == nil {
		var node = NewLinkedArrayNode()
		q.head = node
		q.tail = node
	} else if len(q.tail.array) >= q.sliceSize {
		var node = NewLinkedArrayNode()
		q.tail.next = node
		q.tail = node
	}
	q.tail.array = append(q.tail.array, v)
	q.size++
	return true
}

func (q *LinkedArrayQueue) Dequeue() interface{} {
	if q.size == 0 || q.isEmpty() {
		return nil
	}
	var v = q.head.array[q.pos]
	q.head.array[q.pos] = nil
	q.pos++
	q.size--
	if q.pos >= q.sliceSize {
		var node = q.head.next
		q.head.next = nil
		q.head = node
		q.pos = 0
	}
	return v
}
