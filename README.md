# urlp
A small Go package for manipulating URL query parameters



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

## Benchmark

See [Benchmark.md](./Benchmark.md).