# go-queue-benchmark

Benchmarking different go queue implementation(array, linked list, circular buffer, linked array)

see proposal [here](https://github.com/golang/go/issues/27935)

### Raw result
```
goos: windows
goarch: amd64
pkg: github.com/ichenq/go-queue-benchmark
BenchmarkArrayQueue/case0-8             1000000000               2.74 ns/op            0 B/op          0 allocs/op
BenchmarkArrayQueue/case1-8             30000000                50.8 ns/op            16 B/op          1 allocs/op
BenchmarkArrayQueue/case10-8             3000000               444 ns/op             392 B/op         11 allocs/op
BenchmarkArrayQueue/case100-8             500000              3626 ns/op            3768 B/op        101 allocs/op
BenchmarkArrayQueue/case1000-8             50000             33490 ns/op           39256 B/op       1001 allocs/op
BenchmarkArrayQueue/case10000-8             3000            424571 ns/op          684793 B/op      10008 allocs/op
BenchmarkArrayQueue/case100000-8             200           9778957 ns/op         8709392 B/op     100007 allocs/op
BenchmarkListQueue/case0-8              1000000000               2.73 ns/op            0 B/op          0 allocs/op
BenchmarkListQueue/case1-8              20000000                62.4 ns/op            48 B/op          1 allocs/op
BenchmarkListQueue/case10-8              2000000               755 ns/op             552 B/op         19 allocs/op
BenchmarkListQueue/case100-8              200000              7868 ns/op            5592 B/op        199 allocs/op
BenchmarkListQueue/case1000-8              20000             74502 ns/op           55992 B/op       1999 allocs/op
BenchmarkListQueue/case10000-8              2000            797867 ns/op          559995 B/op      19999 allocs/op
BenchmarkListQueue/case100000-8              100          14940021 ns/op         5600011 B/op     199999 allocs/op
BenchmarkCircularBufferQueue/case0-8            1000000000               3.10 ns/op            0 B/op          0 allocs/op
BenchmarkCircularBufferQueue/case1-8            50000000                29.2 ns/op             0 B/op          0 allocs/op
BenchmarkCircularBufferQueue/case10-8            3000000               442 ns/op              72 B/op          9 allocs/op
BenchmarkCircularBufferQueue/case100-8            200000              6186 ns/op            3864 B/op        101 allocs/op
BenchmarkCircularBufferQueue/case1000-8            30000             62319 ns/op           54072 B/op       1007 allocs/op
BenchmarkCircularBufferQueue/case10000-8            2000            677252 ns/op          582995 B/op      10015 allocs/op
BenchmarkCircularBufferQueue/case100000-8            200           9346010 ns/op         8264010 B/op     100025 allocs/op
BenchmarkLinkedArrayQueue/case0-8               1000000000               2.68 ns/op            0 B/op          0 allocs/op
BenchmarkLinkedArrayQueue/case1-8               100000000               22.1 ns/op            16 B/op          0 allocs/op
BenchmarkLinkedArrayQueue/case10-8               5000000               368 ns/op             234 B/op          9 allocs/op
BenchmarkLinkedArrayQueue/case100-8               500000              3530 ns/op            2417 B/op        100 allocs/op
BenchmarkLinkedArrayQueue/case1000-8               50000             34588 ns/op           24242 B/op       1014 allocs/op
BenchmarkLinkedArrayQueue/case10000-8               5000            376400 ns/op          242492 B/op      10155 allocs/op
BenchmarkLinkedArrayQueue/case100000-8               300           3942790 ns/op         2425007 B/op     101561 allocs/op
```