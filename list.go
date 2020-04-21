// Copyright © 2020 ichenq@outlook.com All rights reserved.
// Distributed under the terms and conditions of the BSD License.
// See accompanying files LICENSE.

package go_queue_benchmark

import (
	"container/list"
	"fmt"
	"strings"
)

type ListQueue struct {
	list *list.List
}

func NewListQueue() Queue {
	return &ListQueue{
		list: list.New(),
	}
}

func (q ListQueue) String() string {
	var sb = &strings.Builder{}
	for elem := q.list.Front(); elem != nil; elem = elem.Next() {
		sb.WriteString(fmt.Sprintf("%v ", elem.Value))
	}
	return sb.String()
}

func (q *ListQueue) Len() int {
	return q.list.Len()
}

func (q *ListQueue) Front() interface{} {
	if q.list.Len() > 0 {
		return q.list.Front().Value
	}
	return nil
}

func (q *ListQueue) Enqueue(v interface{}) bool {
	q.list.PushBack(v)
	return true
}

func (q *ListQueue) Dequeue() interface{} {
	if q.list.Len() > 0 {
		front := q.list.Front()
		q.list.Remove(front)
		return front.Value
	}
	return nil
}
