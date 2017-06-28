# go-mobileapp

## Benchmark

```
$ go test -benchmem -bench .
BenchmarkRegexp-4   	  200000	      6541 ns/op	    2016 B/op	      19 allocs/op
BenchmarkSplit-4    	 1000000	      1801 ns/op	     528 B/op	      18 allocs/op
BenchmarkAtoi-4     	 5000000	       293 ns/op	     144 B/op	       3 allocs/op
BenchmarkScan-4     	  500000	      2632 ns/op	    5760 B/op	       8 allocs/op
BenchmarkBytes-4    	50000000	        23.6 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/dictav/go-mobileapp	7.535s
```
