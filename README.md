# golang-chunks-benchmarks
Benchmarks for chunks algorithm

Added variant `Build` (based on `Chunks`) using `strings.Builder` instead of string slices. The benchmark results:

```
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/ardnew/golang-chunks-benchmarks
cpu: AMD Ryzen Threadripper 1950X 16-Core Processor 
BenchmarkChunks-32                           465           2393822 ns/op          316768 B/op       2688 allocs/op
BenchmarkSplitSubN-32                         56          19482647 ns/op        26870796 B/op      70765 allocs/op
BenchmarkChunkString-32                      356           3392278 ns/op          609264 B/op       2690 allocs/op
BenchmarkChunkStringImproved-32              351           3273287 ns/op          502336 B/op       2638 allocs/op
BenchmarkBuild-32                            651           1888669 ns/op          413992 B/op       5494 allocs/op
PASS
ok      github.com/ardnew/golang-chunks-benchmarks      6.970s
```
