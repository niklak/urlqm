## Benchmark

go version: go1.23.1 linux/amd64

```
goos: linux
goarch: amd64
pkg: github.com/niklak/urlqm
cpu: AMD Ryzen 9 6900HX with Radeon Graphics
```

### Parse params from query string

```bash
go test -test.bench BenchmarkParseParams -run=^Bench -benchmem -benchtime 5s ./test
```

```
BenchmarkParseParamsStd-16               7365163    818.6 ns/op    208 B/op    10 allocs/op
BenchmarkParseParamsUrlP-16              8461663    709.5 ns/op    288 B/op     5 allocs/op
```

### Parse and get a last parameter's value
> **Situation**: you just need to extract a single parameter's value from the query string.
> Using standard library you forced to parse the whole query string, and then extract a single parameter value from it.
Keep in mind that almost all parameters have more than one value. 
So there is no need to decode every parameter to get a certain parameter value.

```bash
go test -test.bench BenchmarkGetQueryParamOne -run=^Bench -benchmem -benchtime 5s ./test
```

```
BenchmarkGetQueryParamOne-16            52150454    111.0 ns/op      0 B/op     0 allocs/op
BenchmarkGetQueryParamOneUrlP-16         8411431    717.8 ns/op    288 B/op     5 allocs/op
BenchmarkGetQueryParamOneStd-16          7332841    816.0 ns/op    208 B/op    10 allocs/op
```

### Parse and get all parameters values

```bash
go test -test.bench BenchmarkGetQueryParamAll -run=^Bench -benchmem -benchtime 5s ./test
```

```
BenchmarkGetQueryParamAll-16            24838568    230.7 ns/op     64 B/op     3 allocs/op
BenchmarkGetQueryParamAllUrlP-16         7596369    783.1 ns/op    336 B/op     7 allocs/op
BenchmarkGetQueryParamAllStd-16          7403596    803.4 ns/op    208 B/op    10 allocs/op
```

### Encode params to query string

> [!NOTE]
> standard library `net/url` breaks the order of parameters on parsing, and always sorts them by key when encoding.
> So this is a hidden behavior that you should keep in mind.
> It would be better to have a separate function for sorting params before encoding. I haven't observe forced sorting of query params outside go `net/url`.

```bash
go test -test.bench BenchmarkEncodeParams -run=^Bench -benchmem -benchtime 5s ./test
```

```
BenchmarkEncodeParamsStd-16              5800682      1036 ns/op    840 B/op    12 allocs/op
BenchmarkEncodeParamsUrlP-16             6696148     878.1 ns/op    936 B/op     9 allocs/op
BenchmarkEncodeParamsSortedUrlP-16       5761309      1019 ns/op    864 B/op    14 allocs/op
```

### Add a new parameter

```bash
go test -test.bench BenchmarkAddParam -run=^Bench -benchmem -benchtime 5s ./test
```

```
BenchmarkAddParam-16                     62475698    80.87 ns/op    224 B/op    1 allocs/op
BenchmarkAddParamUrlP-16                 72796962    85.07 ns/op    384 B/op    1 allocs/op
BenchmarkAddParamStd-16                 129659911    40.65 ns/op     92 B/op    0 allocs/op
```

### Delete parameters

```bash
go test -test.bench BenchmarkDeleteParam -run=^Bench -benchmem -benchtime 5s ./test
```

```
BenchmarkDeleteParam-16                  47607829    118.3 ns/op    304 B/op    2 allocs/op
BenchmarkDeleteParamUrlP-16             430348905    13.86 ns/op      0 B/op    0 allocs/op
BenchmarkDeleteParamAll-16               25986994    228.0 ns/op    288 B/op    2 allocs/op
BenchmarkDeleteParamAllUrlP-16          434766319    13.81 ns/op      0 B/op    0 allocs/op
BenchmarkDeleteParamAllStd-16           708161394    8.454 ns/op      0 B/op    0 allocs/op
```


### Extract parameters

```bash
go test  -test.bench BenchmarkExtractParam  -run=^Bench  -benchmem  -benchtime 5s  ./test
```

```
BenchmarkExtractParam-16                 39426294    133.7 ns/op    304 B/op    2 allocs/op
BenchmarkExtractParamUrlP-16            392536394    15.14 ns/op      0 B/op    0 allocs/op
BenchmarkExtractParamAll-16              16287254    363.2 ns/op    352 B/op    5 allocs/op
BenchmarkExtractParamAllUrlP-16         396350113    14.90 ns/op      0 B/op    0 allocs/op
```

### Set a parameter

```bash
go test -test.bench BenchmarkSetParam -run=^Bench -benchmem -benchtime 5s ./test
```

```
BenchmarkSetParam-16                     21202060     277.8 ns/op    688 B/op    3 allocs/op
BenchmarkSetParamUrlP-16                 65880314     89.22 ns/op    384 B/op    1 allocs/op
BenchmarkSetParamStd-16                 200352841     29.96 ns/op     16 B/op    1 allocs/op
BenchmarkSetParamExisting-16             15044833     397.7 ns/op    688 B/op    4 allocs/op
BenchmarkSetParamExistingUrlP-16        339079548     17.20 ns/op      0 B/op    0 allocs/op
BenchmarkSetParamExistingStd-16         217899476     27.61 ns/op     16 B/op    1 allocs/op
```

### Parse query and then set a new param

```bash
go test -test.bench BenchmarkParseSetParam -run=^Bench -benchmem -benchtime 5s ./test
```
> Note: Added BenchmarkSetParam-16, which doesn't need to be parsed to set a param

```
BenchmarkSetParam-16                     21202060     277.8 ns/op    688 B/op     3 allocs/op
BenchmarkParseSetParamUrlP-16             7148001     827.4 ns/op    672 B/op     6 allocs/op
BenchmarkParseSetParamStd-16              6314110     944.2 ns/op    624 B/op    13 allocs/op
BenchmarkParseSetParamExistingUrlP-16     8209448     732.5 ns/op    288 B/op     5 allocs/op
BenchmarkParseSetParamExistingStd-16      6340438     949.1 ns/op    624 B/op    13 allocs/op
```

### Parse query and then check if the last param exists


```bash
go test -test.bench BenchmarkHasQueryParam -run=^Bench -benchmem -benchtime 5s ./test
```

```
BenchmarkHasQueryParam-16               154644997    39.17 ns/op      0 B/op     0 allocs/op
BenchmarkHasQueryParamUrlP-16             8408752    720.1 ns/op    288 B/op     5 allocs/op
BenchmarkHasQueryParamStd-16              7297198    816.9 ns/op    208 B/op    10 allocs/op
```