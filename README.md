# go-mobileapp

## Benchmark

```
$ go test -benchmem -bench .
Benchmark_parseRegexp-4   	  200000	      6541 ns/op	    2016 B/op	      19 allocs/op
Benchmark_parseSplit-4    	 1000000	      1801 ns/op	     528 B/op	      18 allocs/op
Benchmark_detectAtoi-4     	 5000000	       293 ns/op	     144 B/op	       3 allocs/op
Benchmark_detectScan-4     	  500000	      2632 ns/op	    5760 B/op	       8 allocs/op
Benchmark_detectBytes-4    	50000000	        23.6 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/dictav/go-mobileapp	7.535s
```
