# tstorage-bench
Benchmark tests for [Tstorage](https://github.com/nakabonne/tstorage) project.

Two storage engines were compared: tstorage and [goleveldb](https://github.com/syndtr/goleveldb).

```
$ go test -benchtime=1s -benchmem -bench=. -cpuprofile=./cpu.out -memprofile=./mem.out .
goos: darwin
goarch: amd64
pkg: github.com/nakabonne/tstorage-bench
cpu: Intel(R) Core(TM) i5-7267U CPU @ 3.10GHz
BenchmarkStorage_InsertRows-4            2838434               371.4 ns/op           176 B/op
2 allocs/op
BenchmarkTstorage_Select-4               3639362               348.7 ns/op            56 B/op
2 allocs/op
PASS
ok      github.com/nakabonne/tstorage-bench     5.283s
```
