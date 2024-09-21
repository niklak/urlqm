# urlp
A small Go package for manipulating URL query parameters

## Benchmark


go test -test.bench BenchmarkParseParams -run=^Bench -benchmem -benchtime 10000x ./...
go test -test.bench BenchmarkGetQueryParamOne -run=^Bench -benchmem -benchtime 10000x ./...
go test -test.bench BenchmarkGetQueryParamAll -run=^Bench -benchmem -benchtime 10000x ./...


```
goos: linux
goarch: amd64
pkg: github.com/niklak/urlp
cpu: AMD Ryzen 9 6900HX with Radeon Graphics
```

### Parse params from query string

```
BenchmarkParseParamsStd-16         10000               650.3 ns/op           160 B/op          8 allocs/op
BenchmarkParseParamsUrlP-16        10000               751.3 ns/op           544 B/op          7 allocs/op
```

### Get a single parameter value

```
BenchmarkGetQueryParamOne-16               10000               154.1 ns/op            32 B/op          1 allocs/op
BenchmarkGetQueryParamOneUrlP-16           10000               813.0 ns/op           544 B/op          7 allocs/op
BenchmarkGetQueryParamOneStd-16            10000               685.5 ns/op           160 B/op          8 allocs/op
```

### Get all parameters values

```
BenchmarkGetQueryParamAll-16               10000               243.1 ns/op            64 B/op          3 allocs/op
BenchmarkGetQueryParamAllUrlP-16           10000               848.8 ns/op           592 B/op          9 allocs/op
BenchmarkGetQueryParamAllStd-16            10000               672.4 ns/op           160 B/op          8 allocs/op
```