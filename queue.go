// Copyright © 2020 ichenq@outlook.com All rights reserved.
// Distributed under the terms and conditions of the BSD License.
// See accompanying files LICENSE.

package go_queue_benchmark

type Queue interface {
	Len() int
	Front() interface{}
	Enqueue(interface{})
	Dequeue() interface{}
}
