// Copyright © 2020 ichenq@outlook.com All rights reserved.
// Distributed under the terms and conditions of the BSD License.
// See accompanying files LICENSE.

package go_queue_benchmark

import (
	"fmt"
	"strings"
)

type ArrayQueue struct {
	buf []interface{}
}

func NewArrayQueue(capacity int) Queue {
	if capacity <= 0 {
		capacity = MinQueueBufferCapacity
	}
	return &ArrayQueue{
		buf: make([]interface{}, 0, capacity),
	}
}

func (q ArrayQueue) String() string {
	var sb = &strings.Builder{}
	for i := 0; i < len(q.buf); i++ {
		sb.WriteString(fmt.Sprintf("%v ", q.buf[i]))
	}
	return sb.String()
}

func (q *ArrayQueue) Len() int {
	return len(q.buf)
}

func (q *ArrayQueue) Front() interface{} {
	if len(q.buf) > 0 {
		return q.buf[0]
	}
	return nil
}

func (q *ArrayQueue) Enqueue(v interface{}) bool {
	q.buf = append(q.buf, v)
	return true
}

func (q *ArrayQueue) Dequeue() interface{} {
	if len(q.buf) > 0 {
		v := q.buf[0]
		q.buf = q.buf[1:]
		return v
	}
	return nil
}
