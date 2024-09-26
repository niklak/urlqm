# urlp
A small Go package for manipulating URL query parameters


## Installation

```bash
go get github.com/niklak/urlp
```


## Motivation

For the past 6-7 years, I've worked extensively with `net/url`, and in my experience, 
there have been many cases where url.ParseQuery() behaved unexpectedly, leading to certain issues.

- After manipulating URL parameters, the original order of parameters in the string is always broken.  As a result, encoding in the standard `net/url` forcibly sorts the parameters.

- Although url.Values is represented as `map[string][]string`, the interface is designed to work with a single parameter. This makes sense since most parameters in the string are unique.

- `url.URL.Query()` can quietly "drop" an important parameter that it fails to parse. Of course, it's rare to encounter such URLs, but they do still exist. Therefore, it is better to use url.ParseQuery(), as this method will at least return an error.


>I won't lie, I was also interested in improving the performance of operations related to URL parameters, and according to benchmarks, I succeeded. Sometimes, the difference is quite significant. However, I would still describe the performance of the approach using net/url as very fast, and it's unlikely to have a significant impact on overall performance.

In an effort to improve this situation, I developed two different approaches. The first involves directly manipulating the parameter string. The second involves parsing the parameters into a slice instead of a map.

The first approach is suitable for a small number of operations, where the decoding and encoding time is less than the total time of the operations.

The second approach is better for a larger number of operations, or when you need to sort the parameters or set a specific order (not alphabetical).

Both the first and second approaches preserve the original order of the parameters.


## Usage

### Query string direct manipulation

#### Get a parameter value

```go

u, err := url.Parse("https://example.com?a=1&b=2")
if err != nil {
    panic(err)
}
// there is no need to decode the whole query string
val, err := GetQueryParam(u.RawQuery, "a")
if err != nil {
    // handle this error
    log.Println("Error:", err)
}
fmt.Println(val)
```

#### Get all parameter's values

```go
u, err := url.Parse("https://example.com?a=1&b=2&a=3&c=4")
if err != nil {
    panic(err)
}
values, err := GetQueryParamAll(u.RawQuery, "a")
if err != nil {
    // handle this error
    log.Println("Error:", err)
}
fmt.Println(values)

```

### Extract a parameter value

```go
u, err := url.Parse("https://example.com?a=1&b=2")
if err != nil {
    panic(err)
}
val, err := ExtractQueryParam(&u.RawQuery, "a")
if err != nil {
    // handle this error
    log.Println("Error:", err)
}
fmt.Println(val)
// url was modified
fmt.Println(u)
```

### Extract all parameter's values

```go
u, err := url.Parse("https://example.com?a=1&b=2&a=3&c=4")
if err != nil {
    panic(err)
}
values, err := ExtractQueryParamAll(&u.RawQuery, "a")
if err != nil {
    // handle this error
    log.Println("Error:", err)
}
fmt.Println(values)
// url was modified
fmt.Println(u)
```

### Add a query parameter

```go
u, err := url.Parse("https://example.com?a=1&b=2")
if err != nil {
    panic(err)
}
AddQueryParam(&u.RawQuery, "c", "3", "4")

fmt.Println(u)
```

### Set a query parameter

```go
u, err := url.Parse("https://example.com?a=1&b=2&a=3&b=4")
if err != nil {
    panic(err)
}
SetQueryParam(&u.RawQuery, "b", "5")
fmt.Println(u)
```

### Delete a single value of query parameter

```go
u, err := url.Parse("https://example.com?a=1&b=2&a=3&b=4")
if err != nil {
    panic(err)
}
DeleteQueryParam(&u.RawQuery, "a")

fmt.Println(u)
```

### Delete all parameters with the same name

```go
u, err := url.Parse("https://example.com?a=1&b=2&a=3&b=4")
if err != nil {
    panic(err)
}
DeleteQueryParamAll(&u.RawQuery, "a")

fmt.Println(u)
```

### Check if query string contains a parameter

```go
u, err := url.Parse("https://example.com?a=1&b=2")
if err != nil {
    panic(err)
}

fmt.Println("a:", HasQueryParam(u.RawQuery, "a"))
fmt.Println("c:", HasQueryParam(u.RawQuery, "c"))
```


## Benchmark

See [Benchmark.md](./Benchmark.md).