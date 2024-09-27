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

> [!NOTE]
> Most of the query keys contain only ASCII characters, so it's fine to pass an unescaped key to a query function.
> But if the key may contain non-ASCII characters, you should escape them before passing them to functions.
> - GetQueryParam
> - GetQueryParamAll
> - ExtractQueryParam
> - ExtractQueryParamAll
> - DeleteQueryParam
> - DeleteQueryParamAll
> - HasQueryParam
>
> *So why this approach doesn't encode query keys by itself?* 
>
> Because it brings a little (very little) overhead even if the key doesn't contains non-ASCII characters.
> There is an option to add a trailing argument encKey bool, which will encode the query keys by itself. But I'm still not sure about it.

<details>
<summary>Get a parameter value</summary>

```go

rawURL := "https://example.com?a=1&b=2&%D0%BA%D0%BB%D1%8E%D1%87=%D0%B7%D0%BD%D0%B0%D1%87%D0%B5%D0%BD%D0%BD%D1%8F"
u, err := url.Parse(rawURL)
if err != nil {
    panic(err)
}
val, err := GetQueryParam(u.RawQuery, "a")
if err != nil {
    // handle this error
    fmt.Println("Error:", err)
}
fmt.Println(val)

// if the key contains non-ASCII characters, you must encode it before calling GetQueryParam
val, err = GetQueryParam(u.RawQuery, url.QueryEscape("ключ"))
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(val)
```

</details>


<details>
<summary>Get all parameter's values</summary>

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

</details>


<details>
<summary>Extract a parameter value (get and remove)</summary>

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

</details>


<details>
<summary>Extract all parameter's values</summary>

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

</details>

<details>
<summary>Add a query parameter</summary>


```go
u, err := url.Parse("https://example.com?a=1&b=2")
if err != nil {
    panic(err)
}
AddQueryParam(&u.RawQuery, "c", "3", "4")

fmt.Println(u)
```

</details>


<details>
<summary>Set a query parameter</summary>

```go
u, err := url.Parse("https://example.com?a=1&b=2&a=3&b=4")
if err != nil {
    panic(err)
}
SetQueryParam(&u.RawQuery, "b", "5")
fmt.Println(u)
```

</details>


<details>
<summary>Delete a single value of query parameter</summary>

```go
u, err := url.Parse("https://example.com?a=1&b=2&a=3&b=4")
if err != nil {
    panic(err)
}
DeleteQueryParam(&u.RawQuery, "a")

fmt.Println(u)
```

</details>


<details>
<summary>Delete all parameters with the same name</summary>

```go
u, err := url.Parse("https://example.com?a=1&b=2&a=3&b=4")
if err != nil {
    panic(err)
}
DeleteQueryParamAll(&u.RawQuery, "a")

fmt.Println(u)
```

</details>


<details>
<summary>Check if query string contains a parameter</summary>

```go
u, err := url.Parse("https://example.com?a=1&b=2")
if err != nil {
    panic(err)
}

fmt.Println("a:", HasQueryParam(u.RawQuery, "a"))
fmt.Println("c:", HasQueryParam(u.RawQuery, "c"))
```

</details>

### Manipulations with query parameter list

<details>
<summary>Parse a query string to the parameter list</summary>

```go
u, err := url.Parse(`https://example.com/?q=100%25+truth&a=1&b=2`)
if err != nil {
    panic(err)
}
params, err := ParseQuery(u.RawQuery)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Printf("%+v\n", params)
```

</details>

<details>
<summary>Encode a parameter list to the query string</summary>

```go
u, err := url.Parse(`https://example.com/`)
if err != nil {
    panic(err)
}
params := Params{{"q", "100% truth"}, {"a", "1"}, {"b", "2"}}

u.RawQuery = params.Encode()

fmt.Println(u)
```

</details>


<details>
<summary>Sort parameters by key</summary>

```go
u, err := url.Parse(`https://example.com/?q=100%25+truth&a=1&b=2`)
if err != nil {
    panic(err)
}
params, err := ParseQuery(u.RawQuery)
if err != nil {
    fmt.Println("Error:", err)
}
params.Sort()
u.RawQuery = params.Encode()
fmt.Println(u)
```

</details>


<details>
<summary>Set a specific order for parameters</summary>

```go
u, err := url.Parse(`https://example.com/?q=100%25+truth&a=1&b=2&d=4&c=3`)
if err != nil {
    panic(err)
}
params, err := ParseQuery(u.RawQuery)
if err != nil {
    fmt.Println("Error:", err)
}
params.SetOrder("q", "d", "c", "b", "a")
u.RawQuery = params.Encode()
fmt.Println(u)
```

</details>


<details>
<summary>Add a parameter (multiple values)</summary>

```go
u, err := url.Parse(`https://example.com/?q=100%25+truth&a=1&b=2`)
if err != nil {
    panic(err)
}
params, err := ParseQuery(u.RawQuery)
if err != nil {
    fmt.Println("Error:", err)
}
params.Add("c", "3", "4")
u.RawQuery = params.Encode()
fmt.Println(u)
```

</details>


<details>
<summary>Set a new parameter</summary>

```go
u, err := url.Parse(`https://example.com/?q=100%25+truth&a=1&b=2&c=3&b=4`)
if err != nil {
    panic(err)
}
params, err := ParseQuery(u.RawQuery)
if err != nil {
    fmt.Println("Error:", err)
}
params.Set("b", "5")
u.RawQuery = params.Encode()
fmt.Println(u)
```

</details>


<details>
<summary>Get a parameter's value</summary>

```go
u, err := url.Parse(`https://example.com/?q=100%25+truth&a=1&b=2`)
if err != nil {
    panic(err)
}
params, err := ParseQuery(u.RawQuery)
if err != nil {
    fmt.Println("Error:", err)
}
val := params.Get("b")
fmt.Println(val)
```
</details>


<details>
<summary>Get all parameter's values</summary>

```go
u, err := url.Parse(`https://example.com/?q=100%25+truth&a=1&b=2&a=3`)
if err != nil {
    panic(err)
}
params, err := ParseQuery(u.RawQuery)
if err != nil {
    fmt.Println("Error:", err)
}
values := params.GetAll("a")
fmt.Println(values)
```

</details>

<details>
<summary>Extract a single parameter from list by key (get and remove)</summary>

```go
u, err := url.Parse(`https://example.com/?q=100%25+truth&a=1&b=2&a=3`)
if err != nil {
    panic(err)
}
params, err := ParseQuery(u.RawQuery)
if err != nil {
    fmt.Println("Error:", err)
}
val := params.Extract("a")
u.RawQuery = params.Encode()

fmt.Println(val)
fmt.Println(u)
```

</details>

<details>
<summary>Extract all parameters from list by key</summary>

```go
u, err := url.Parse(`https://example.com/?q=100%25+truth&a=1&b=2&a=3`)
if err != nil {
    panic(err)
}
params, err := ParseQuery(u.RawQuery)
if err != nil {
    fmt.Println("Error:", err)
}
values := params.ExtractAll("a")
u.RawQuery = params.Encode()

fmt.Println(values)
fmt.Println(u)
```

</details>

<details>
<summary>Delete a parameter by key</summary>

```go
u, err := url.Parse(`https://example.com/?q=100%25+truth&a=1&b=2&a=3`)
if err != nil {
    panic(err)
}
params, err := ParseQuery(u.RawQuery)
if err != nil {
    fmt.Println("Error:", err)
}
params.Delete("a")
u.RawQuery = params.Encode()
fmt.Println(u)
```

</details>

<details>
<summary>Delete all parameters by key</summary>

```go
u, err := url.Parse(`https://example.com/?q=100%25+truth&a=1&b=2&a=3`)
if err != nil {
    panic(err)
}
params, err := ParseQuery(u.RawQuery)
if err != nil {
    fmt.Println("Error:", err)
}
params.DeleteAll("a")
u.RawQuery = params.Encode()
fmt.Println(u)

```

</details>


<details>
<summary>Check if parameter list has a parameter with a key</summary>

```go
u, err := url.Parse(`https://example.com/?q=100%25+truth&a=1&b=2`)
if err != nil {
    panic(err)
}
params, err := ParseQuery(u.RawQuery)
if err != nil {
    fmt.Println("Error:", err)
}
fmt.Println(params.Has("a"))
fmt.Println(params.Has("c"))
```

</details>

## Benchmark

See [Benchmark.md](./Benchmark.md).