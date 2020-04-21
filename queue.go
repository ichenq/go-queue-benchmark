// Copyright © 2020 ichenq@outlook.com All rights reserved.
// Distributed under the terms and conditions of the BSD License.
// See accompanying files LICENSE.

package go_queue_benchmark

const MinQueueBufferCapacity = 16

// Queue is a FIFO queue
type Queue interface {
	Len() int
	Front() interface{}
	Enqueue(interface{}) bool
	Dequeue() interface{}
}

// Deque is a double-ended queue
type Deque interface {
	Len() int
	Front() interface{}
	Back() interface{}
	PushBack(interface{})
	PushFront(interface{})
	PopFront() interface{}
	PopBack() interface{}
}