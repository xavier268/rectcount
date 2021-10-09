# rectcount
Count all (non-degenerated) rectangles that can be found from a set of interger coordiantes points.

# Problem statement

Given a set of (distinct) points of integer coordinates, find a time efficient algorithm to count the number of rectangles.

# Performance 

Currently estimated O(n^2)log(n) computation time.
The log(n) comes from the need to use a hash map structure.

Benchmark and test results confirms O(n^2)Log(n) likelyhood ...

```
go test -bench=. -benchmem

goos: linux
goarch: amd64
pkg: github.com/xavier268/rectcount
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkCount50-8          9254            127181 ns/op          262860 B/op       1233 allocs/op
BenchmarkCount500-8           52          21023347 ns/op        31874901 B/op     124758 allocs/op
BenchmarkCount5000-8           1        2705828038 ns/op        2182755264 B/op 12497514 allocs/op
PASS
ok      github.com/xavier268/rectcount  5.572s
```