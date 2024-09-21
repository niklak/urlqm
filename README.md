# urlp
A small Go package for manipulating URL query parameters

## Benchmark


go test -test.bench BenchmarkParseParams -run=^Bench -benchmem -benchtime 10000x ./test
go test -test.bench BenchmarkGetQueryParamOne -run=^Bench -benchmem -benchtime 10000x ./test
go test -test.bench BenchmarkGetQueryParamAll -run=^Bench -benchmem -benchtime 10000x ./test


```
goos: linux
goarch: amd64
pkg: github.com/niklak/urlp
cpu: AMD Ryzen 9 6900HX with Radeon Graphics
```

### Parse params from query string

```
BenchmarkParseParamsStd-16         10000               746.3 ns/op           160 B/op          8 allocs/op
BenchmarkParseParamsUrlP-16        10000               681.8 ns/op           224 B/op          4 allocs/op
```

### Get a single parameter value

```
BenchmarkGetQueryParamOne-16               10000               148.7 ns/op            32 B/op          1 allocs/op
BenchmarkGetQueryParamOneUrlP-16           10000               593.0 ns/op           224 B/op          4 allocs/op
BenchmarkGetQueryParamOneStd-16            10000               626.5 ns/op           160 B/op          8 allocs/op
```

### Get all parameters values

```
BenchmarkGetQueryParamAll-16               10000               237.9 ns/op            64 B/op          3 allocs/op
BenchmarkGetQueryParamAllUrlP-16           10000               579.9 ns/op           272 B/op          6 allocs/op
BenchmarkGetQueryParamAllStd-16            10000               627.8 ns/op           160 B/op          8 allocs/op
```