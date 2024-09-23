# urlp
A small Go package for manipulating URL query parameters

## Benchmark




```
goos: linux
goarch: amd64
pkg: github.com/niklak/urlp
cpu: AMD Ryzen 9 6900HX with Radeon Graphics
```

### Parse params from query string

```bash
go test -test.bench BenchmarkParseParams -run=^Bench -benchmem -benchtime 5s ./test
```

```
BenchmarkParseParamsStd-16               10128763    588.4 ns/op    160 B/op    8 allocs/op
BenchmarkParseParamsUrlP-16              11921043    502.9 ns/op    224 B/op    4 allocs/op
```

### Get a single parameter value
> **Situation**: you just need to extract a single parameter's value from the query string.
> Using standard library you forced to parse the whole query string, and then extract a single parameter value from it.
Keep in mind that almost all parameters have more than one value. 
So there is no need to decode every parameter to get a certain parameter value.

```bash
go test -test.bench BenchmarkGetQueryParamOne -run=^Bench -benchmem -benchtime 5s ./test
```

```
BenchmarkGetQueryParamOne-16             39516896    145.7 ns/op     32 B/op    1 allocs/op
BenchmarkGetQueryParamOneUrlP-16         11472595    515.3 ns/op    224 B/op    4 allocs/op
BenchmarkGetQueryParamOneStd-16           9955138    602.1 ns/op    160 B/op    8 allocs/op
```

### Get all parameters values

```bash
go test -test.bench BenchmarkGetQueryParamAll -run=^Bench -benchmem -benchtime 5s ./test
```

```
BenchmarkGetQueryParamAll-16             24137166    250.4 ns/op     64 B/op    3 allocs/op
BenchmarkGetQueryParamAllUrlP-16         10403386    579.6 ns/op    272 B/op    6 allocs/op
BenchmarkGetQueryParamAllStd-16           9946518    600.1 ns/op    160 B/op    8 allocs/op
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
BenchmarkEncodeParamsStd-16              9185296    647.0 ns/op    376 B/op     9 allocs/op
BenchmarkEncodeParamsUrlP-16            11621364    521.4 ns/op    312 B/op     8 allocs/op
BenchmarkEncodeParamsSortedUrlP-16       9320493    646.9 ns/op    416 B/op    11 allocs/op
```

### Add a new parameter

```bash
go test -test.bench BenchmarkAddParam -run=^Bench -benchmem -benchtime 5s ./test
```

```
BenchmarkAddParam-16                     51030302    115.5 ns/op    384 B/op    2 allocs/op
BenchmarkAddParamUrlP-16                 81422187    75.30 ns/op    320 B/op    1 allocs/op
BenchmarkAddParamStd-16                 172536259    39.46 ns/op     87 B/op    0 allocs/op
```

### Delete parameters

```bash
go test -test.bench BenchmarkDeleteParam -run=^Bench -benchmem -benchtime 5s ./test
```

```
BenchmarkDeleteParam-16                  81099346    69.07 ns/op    112 B/op    1 allocs/op
BenchmarkDeleteParamUrlP-16             437800039    13.75 ns/op      0 B/op    0 allocs/op
BenchmarkDeleteParamAll-16               35512520    170.9 ns/op    144 B/op    2 allocs/op
BenchmarkDeleteParamAllUrlP-16          439215859    13.61 ns/op      0 B/op    0 allocs/op
BenchmarkDeleteParamAllStd-16           718473314    8.609 ns/op      0 B/op    0 allocs/op
```


### Extract parameters

```bash
go test  -test.bench BenchmarkExtractParam  -run=^Bench  -benchmem  -benchtime 5s  ./test
```

```
BenchmarkExtractParam-16                 81509199    73.68 ns/op    112 B/op    1 allocs/op
BenchmarkExtractParamUrlP-16            383607861    15.49 ns/op      0 B/op    0 allocs/op
BenchmarkExtractParamAll-16              18252739    323.3 ns/op    208 B/op    5 allocs/op
BenchmarkExtractParamAllUrlP-16         410659760    14.57 ns/op      0 B/op    0 allocs/op
```

### Set a parameter

```bash
go test -test.bench BenchmarkSetParam -run=^Bench -benchmem -benchtime 5s ./test
```

```
BenchmarkSetParam-16                     25434712    238.9 ns/op    400 B/op    3 allocs/op
BenchmarkSetParamUrlP-16                 72594936    79.54 ns/op    320 B/op    1 allocs/op
BenchmarkSetParamStd-16                 202840244    29.68 ns/op     16 B/op    1 allocs/op
BenchmarkSetParamExisting-16             17748502    333.4 ns/op    232 B/op    5 allocs/op
BenchmarkSetParamExistingUrlP-16        351103875    17.36 ns/op      0 B/op    0 allocs/op
BenchmarkSetParamExistingStd-16         218821728    27.38 ns/op     16 B/op    1 allocs/op
```

### Parse query params and then set a new param

```bash
go test -test.bench BenchmarkParseSetParam -run=^Bench -benchmem -benchtime 5s ./test
```

```
BenchmarkParseSetParamUrlP-16            10121535    598.9 ns/op    544 B/op     5 allocs/op
BenchmarkParseSetParamStd-16              7931796    754.8 ns/op    576 B/op    11 allocs/op
BenchmarkParseSetParamExistingUrlP-16    11554497    513.6 ns/op    224 B/op     4 allocs/op
BenchmarkParseSetParamExistingStd-16      8157975    734.3 ns/op    576 B/op    11 allocs/op
```