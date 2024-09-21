# urlp
A small Go package for manipulating URL query parameters

## Benchmark


go test -test.bench BenchmarkParseParams -run=^Bench -benchmem -benchtime 5s ./test
go test -test.bench BenchmarkGetQueryParamOne -run=^Bench -benchmem -benchtime 5s ./test
go test -test.bench BenchmarkGetQueryParamAll -run=^Bench -benchmem -benchtime 5s ./test
go test -test.bench BenchmarkEncodeParams -run=^Bench -benchmem -benchtime 5s ./test


```
goos: linux
goarch: amd64
pkg: github.com/niklak/urlp
cpu: AMD Ryzen 9 6900HX with Radeon Graphics
```

### Parse params from query string

```
BenchmarkParseParamsStd-16      10128763               588.4 ns/op           160 B/op          8 allocs/op
BenchmarkParseParamsUrlP-16     11921043               502.9 ns/op           224 B/op          4 allocs/op
```

### Get a single parameter value

```
BenchmarkGetQueryParamOne-16            39516896               145.7 ns/op            32 B/op          1 allocs/op
BenchmarkGetQueryParamOneUrlP-16        11472595               515.3 ns/op           224 B/op          4 allocs/op
BenchmarkGetQueryParamOneStd-16          9955138               602.1 ns/op           160 B/op          8 allocs/op
```

### Get all parameters values

```
BenchmarkGetQueryParamAll-16            24137166               250.4 ns/op            64 B/op          3 allocs/op
BenchmarkGetQueryParamAllUrlP-16        10403386               579.6 ns/op           272 B/op          6 allocs/op
BenchmarkGetQueryParamAllStd-16          9946518               600.1 ns/op           160 B/op          8 allocs/op
```

### Encode params to query string

**Note:** standard library `net/url` breaks the order of parameters on parsing, and always sorts them by key when encoding.
So this is a hidden behavior that you should keep in mind.
It would be better to have a separate function for sorting params before encoding. I haven't observe forced sorting of query params outside go `net/url`.

```
BenchmarkEncodeParamsStd-16              9185296               647.0 ns/op           376 B/op          9 allocs/op
BenchmarkEncodeParamsUrlP-16            11621364               521.4 ns/op           312 B/op          8 allocs/op
BenchmarkEncodeParamsSortedUrlP-16       9320493               646.9 ns/op           416 B/op         11 allocs/op
```
