package urlp

import (
	"fmt"
	"log"
	"net/url"
)

func ExampleGetQueryParam() {
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
	// Output:
	// 1
	// значення
}

func ExampleGetQueryParamAll() {
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
	fmt.Println(u)
	// Output:
	// [1 3]
	// https://example.com?a=1&b=2&a=3&c=4
}

func ExampleExtractQueryParam() {
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
	fmt.Println(u)
	// Output:
	// 1
	// https://example.com?b=2
}

func ExampleExtractQueryParamAll() {
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
	fmt.Println(u)
	// Output:
	// [1 3]
	// https://example.com?b=2&c=4
}

func ExampleAddQueryParam() {
	u, err := url.Parse("https://example.com?a=1&b=2")
	if err != nil {
		panic(err)
	}
	AddQueryParam(&u.RawQuery, "c", "3", "4")

	fmt.Println(u)
	// Output: https://example.com?a=1&b=2&c=3&c=4
}

func ExampleSetQueryParam() {
	u, err := url.Parse("https://example.com?a=1&b=2&a=3&b=4")
	if err != nil {
		panic(err)
	}
	SetQueryParam(&u.RawQuery, "b", "5")

	fmt.Println(u)
	// Output: https://example.com?a=1&b=5&a=3
}

func ExampleDeleteQueryParam() {
	u, err := url.Parse("https://example.com?a=1&b=2&a=3&b=4")
	if err != nil {
		panic(err)
	}
	DeleteQueryParam(&u.RawQuery, "a")

	fmt.Println(u)
	// Output: https://example.com?b=2&a=3&b=4
}

func ExampleDeleteQueryParamAll() {
	u, err := url.Parse("https://example.com?a=1&b=2&a=3&b=4")
	if err != nil {
		panic(err)
	}
	DeleteQueryParamAll(&u.RawQuery, "a")

	fmt.Println(u)
	// Output: https://example.com?b=2&b=4
}

func ExampleHasQueryParam() {
	u, err := url.Parse("https://example.com?a=1&b=2")
	if err != nil {
		panic(err)
	}

	fmt.Println("a:", HasQueryParam(u.RawQuery, "a"))
	fmt.Println("c:", HasQueryParam(u.RawQuery, "c"))

	// Output:
	// a: true
	// c: false
}
