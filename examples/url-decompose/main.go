package main

import (
	"fmt"
	"net/url"

	"github.com/niklak/urlp"
)

func main() {
	var rawURL = "https://www.example.com/?q=a+pretty+long+query" +
		"&some-hash=EgZjaHJvbWUyBggAEEUYOTIGCAEQABhAMgYIAhAAGEAyBggD" +
		"EAAYQDIGCAQQABhAMgYIBRAAGEAyBggGAYQDIGCAcQABhAMgYICBAAGEDSAQgxMjQxajBqMagCArACAQ&first=11&brightness=90%"

	// At this point url.Parse will not parse a query component
	// because it can be inefficient for urls with many parameters.
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	fmt.Printf("URL: %s\n", u)

	// To get value of a q parameter using a std net/url package,
	// unfortunately we have also to parse the other queries.
	// So this isn't very efficient for large URLs with many parameters when you need just a few of them.
	query := u.Query()
	q := query.Get("q")
	fmt.Printf("q: %s\n", q)
	// also this was a bad approach, because if there is an error, you will not know about it,
	// and corresponding parameter will disappear from the map

	// So more explicit way to parse query parameters would be:
	query, err = url.ParseQuery(u.RawQuery)

	if err != nil {
		fmt.Printf("Error parsing query: %s\n", err)
	}

	// But both approaches will omit 'brightness' parameter, because it contains and invalid character.
	fmt.Printf("Parsed Query: %#v\n", query)

	// At this point we have and url with RawQuery and instead of parsing the whole query string,
	// we can parse just a small peace of string and take the value.

	q, err = urlp.GetQueryParam(u.RawQuery, "q")
	//OR if we deal with encoded keys:
	//q, err = urlp.GetQueryParam(u.RawQuery, url.QueryEscape(`%22q%22`))

	if err != nil {
		panic(err)
	}

	fmt.Printf("urlp q: %s\n", q)

}
