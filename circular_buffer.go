// Copyright © 2020 ichenq@outlook.com All rights reserved.
// Distributed under the terms and conditions of the BSD License.
// See accompanying files LICENSE.

package go_queue_benchmark

import (
	"fmt"
	"strings"
)

// A circular buffer implementation from .net core
// https://referencesource.microsoft.com/#mscorlib/system/collections/queue.cs
type CircularBufferQueue struct {
	array []interface{}
	head  int
	tail  int
	size  int
}

func NewCircularBufferQueue(capacity int) Queue {
	if capacity <= 0 {
		capacity = MinQueueBufferCapacity
	}
	return &CircularBufferQueue{
		array: make([]interface{}, capacity),
	}
}

func (q CircularBufferQueue) String() string {
	var sb = &strings.Builder{}
	for i := q.head; i < q.tail; i++ {
		sb.WriteString(fmt.Sprintf("%v ", q.array[i]))
	}
	return sb.String()
}

func (q *CircularBufferQueue) Len() int {
	return q.size
}

func (q *CircularBufferQueue) Front() interface{} {
	if q.size > 0 {
		return q.array[q.head]
	}
	return nil
}

func (q *CircularBufferQueue) Enqueue(v interface{}) bool {
	if q.size == len(q.array) {
		q.resize()
	}
	q.array[q.tail] = v
	q.tail = (q.tail + 1) % len(q.array)
	q.size++
	return true
}

func (q *CircularBufferQueue) Dequeue() interface{} {
	if q.size == 0 {
		return nil
	}
	var v = q.array[q.head]
	q.array[q.head] = nil
	q.head = (q.head + 1) % len(q.array)
	q.size--
	q.shrink()
	return v
}

// resize down if the buffer 1/4 full
func (q *CircularBufferQueue) shrink() {
	if q.size > MinQueueBufferCapacity && q.size <= len(q.array)/4 {
		q.resize()
	}
}

func (q *CircularBufferQueue) resize() {
	var newcap = q.size
	if q.size < 1024 {
		newcap += q.size // double size
	} else {
		newcap += q.size / 2 // 1.5 x size
	}
	var newarray = make([]interface{}, newcap)
	if q.tail > q.head {
		copy(newarray, q.array[q.head:q.tail])
	} else {
		n := copy(newarray, q.array[q.head:])
		copy(newarray[n:], q.array[:q.tail])
	}
	q.array = newarray
	q.head = 0
	q.tail = q.size
	if q.size == newcap {
		q.tail = 0
	}
}
