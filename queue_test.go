// Copyright © 2020 ichenq@outlook.com All rights reserved.
// Distributed under the terms and conditions of the BSD License.
// See accompanying files LICENSE.

package go_queue_benchmark

import (
	"fmt"
	"testing"
)

func testQueueAPI(que Queue, t *testing.T) {
	if que.Len() != 0 {
		t.Errorf("expect empty queue")
	}
	for i := 1; i <= 100; i++ {
		que.Enqueue(i)
	}
	if que.Len() != 100 {
		t.Errorf("expect queue size with 100")
	}
	for i := 1; i <= 50; i++ {
		var v = que.Front()
		if v == nil {
			t.Errorf("queue front should not be null, %v", que)
		}
		if v.(int) != i {
			t.Errorf("expect %d, but got %v", i, v)
		}
		que.Dequeue()
	}
	if que.Len() != 50 {
		t.Errorf("expect empty queue")
	}
	for i := 101; i <= 150; i++ {
		que.Enqueue(i)
	}
	for i := 51; i <= 150; i++ {
		var v = que.Dequeue()
		if v == nil {
			t.Errorf("queue front should not be null, %v", que)
		}
		if v.(int) != i {
			t.Errorf("expect %d, but got %v", i, v)
		}
	}
	// big alloc
	for i := 0; i < 3000000; i++ {
		que.Enqueue(i)
	}
}

func TestArrayQueue(t *testing.T) {
	var que = NewArrayQueue(0)
	testQueueAPI(que, t)
}

func TestListQueue(t *testing.T) {
	var que = NewListQueue()
	testQueueAPI(que, t)
}

func TestCircularBufferQueue(t *testing.T) {
	var que = NewCircularBufferQueue(0)
	testQueueAPI(que, t)
}

func TestLinkedArrayQueueQueue(t *testing.T) {
	var que = NewLinkedArrayQueue(0)
	testQueueAPI(que, t)
}

type testData struct {
	Count  int64
	Remove bool
}

var tests = []testData{
	{Count: 0},
	{Count: 1},
	{Count: 10},
	{Count: 100, Remove: true},
	{Count: 1000},                // 1k
	{Count: 10000, Remove: true}, //10k
	{Count: 100000},              // 100k
}

// Used to store temp values, avoiding any compiler optimizations.
var dummy interface{}

func runQueueBenchmarks(que Queue, b *testing.B) {
	for _, testCase := range tests {
		b.Run(fmt.Sprintf("case%d", testCase.Count), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				for i := int64(0); i < testCase.Count; i++ {
					que.Enqueue(i)
					if testCase.Remove && i > 0 && i%3 == 0 {
						dummy = que.Dequeue()
					}
				}
				for que.Len() > 0 {
					dummy = que.Dequeue()
				}
			}
		})
	}
}

func BenchmarkArrayQueue(b *testing.B) {
	var que = NewArrayQueue(0)
	runQueueBenchmarks(que, b)
}

func BenchmarkListQueue(b *testing.B) {
	var que = NewListQueue()
	runQueueBenchmarks(que, b)
}

func BenchmarkCircularBufferQueue(b *testing.B) {
	var que = NewCircularBufferQueue(0)
	runQueueBenchmarks(que, b)
}

func BenchmarkLinkedArrayQueue(b *testing.B) {
	var que = NewLinkedArrayQueue(0)
	runQueueBenchmarks(que, b)
}
