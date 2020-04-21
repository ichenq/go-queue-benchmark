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
		t.Fatalf("expect empty queue")
	}
	for i := 1; i <= 100; i++ {
		que.Enqueue(i)
	}
	if que.Len() != 100 {
		t.Fatalf("expect queue size with 100")
	}
	for i := 1; i <= 50; i++ {
		var v = que.Dequeue()
		if v.(int) != i {
			t.Fatalf("expect %d, but got %v", i, v)
		}
	}
	if que.Len() != 50 {
		t.Fatalf("expect empty queue")
	}
	for i := 101; i <= 150; i++ {
		que.Enqueue(i)
	}
	for i := 51; i <= 150; i++ {
		var v = que.Dequeue()
		if v.(int) != i {
			t.Fatalf("expect %d, but got %v", i, v)
		}
	}
}

func TestArrayQueue(t *testing.T) {
	var que = NewArrayQueue(8)
	testQueueAPI(que, t)
}

func TestListQueue(t *testing.T) {
	var que = NewListQueue()
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
	var que = NewArrayQueue(8)
	runQueueBenchmarks(que, b)
}
