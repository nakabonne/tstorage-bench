# tstorage-bench
Benchmark tests for [Tstorage](https://github.com/nakabonne/tstorage) project.

Two storage engines were compared: tstorage and [goleveldb](https://github.com/syndtr/goleveldb).

```
$ go test -benchtime=1s -benchmem -bench=. -cpuprofile=./cpu.out -memprofile=./mem.out .
goos: darwin
goarch: amd64
pkg: github.com/nakabonne/tstorage-bench
cpu: Intel(R) Core(TM) i5-7267U CPU @ 3.10GHz
BenchmarkStorage_InsertParallel-4   	 1000000	      1142 ns/op	     175 B/op	       2 allocs/op
BenchmarkLevelDB_InsertParallel-4   	  196279	      6190 ns/op	     350 B/op	       5 allocs/op
BenchmarkStorage_Insert-4           	 1687567	       732.4 ns/op	     169 B/op	       2 allocs/op
BenchmarkLevelDB_Insert-4           	  186916	      8853 ns/op	     388 B/op	       7 allocs/op
PASS
ok  	github.com/nakabonne/tstorage-bench	39.172s
```
