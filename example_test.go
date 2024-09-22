package urlp

import (
	"errors"
	"fmt"
	"net/url"
)

func ExampleParseParams() {
	params, err := ParseParams(`a=1&q=100%+truth&b=2&brightness=90%`)
	if err != nil {
		// handling this error, keep in mind that params is not empty,
		// it contains all keys and values, both escaped and unescaped
		fmt.Println("Error:", err)
	}
	var e url.EscapeError
	if errors.As(err, &e) {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Printf("%v\n", params)
}

func ExampleSortOrderParams() {
	query := "q=100+truth&c=3&b=2&page=1&a=1"
	params, err := ParseParams(query)
	if err != nil {
		panic(err)
	}
	SortOrderParams(&params, "a", "page", "q")

	fmt.Println("Query:", EncodeParams(params))
	// Output: Query: a=1&page=1&q=100+truth&c=3&b=2

}
