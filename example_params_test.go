package urlp

import (
	"errors"
	"fmt"
	"net/url"
)

func ExampleParseParams() {
	u, err := url.Parse(`https://example.com/?a=1&q=100%+truth&b=2&brightness=90%`)
	if err != nil {
		fmt.Println("Error:", u)
	}
	params, err := ParseParams(u.RawQuery)
	if err != nil {
		// handling this error, keep in mind that params is not empty,
		// it contains all keys and values, both escaped and unescaped
		fmt.Println("Error:", err)
	}
	var e url.EscapeError
	if errors.As(err, &e) {
		fmt.Printf("Error as: %s\n", err)
	}
	fmt.Printf("%v\n", params)
}

func ExampleSortOrderParams() {
	u, err := url.Parse(`https://example.com/?q=100+truth&c=3&b=2&page=1&a=1`)
	if err != nil {
		fmt.Println("Error:", u)
	}
	params, err := ParseParams(u.RawQuery)
	if err != nil {
		panic(err)
	}
	SortOrderParams(&params, "a", "page", "q")

	fmt.Println("Query:", EncodeParams(params))
	// Output: Query: a=1&page=1&q=100+truth&c=3&b=2

}
