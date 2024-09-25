package main

import (
	"fmt"
	"net/url"

	"github.com/niklak/urlp"
)

func main() {
	var rawURL = "https://www.example.com/?q=a+testing+query&region=1&region=3&region=5" +
		"&some-hash=EgZjaHJvbWUyBggAEEUYOTIGCAEQABhAMgYIAhAAGEAyBggD" +
		"EAAYQDIGCAQQABhAMgYIBRAAGEAyBggGAYQDIGCAcQABhAMgYICBAAGEDSAQgxMjQxajBqMagCArACAQ&first=11&brightness=90%"

	// At this point url.Parse will not parse a query component
	// because it can be inefficient for urls with many parameters.
	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	fmt.Printf("URL: %s\n", u)

	// To get value of a `q` parameter using a std net/url package,
	// unfortunately, we have also to parse the other parameters.
	// So this isn't very efficient for large URLs with many parameters when you need just a few of them.
	// also this was a bad approach, because if there is an error, you will not know about it,
	// and corresponding parameter will not appear in the query map.
	query := u.Query()
	q := query.Get("q")
	fmt.Printf("q (std): %s\n", q)

	// So more explicit way to parse query parameters would be:
	query, err = url.ParseQuery(u.RawQuery)
	if err != nil {
		// got error here, now if you wish you can handle it.
		fmt.Printf("Error parsing query: %s\n", err)
	}
	// But both approaches will omit 'brightness' parameter, because it contains and invalid character.
	fmt.Printf("Parsed Query: %#v\n", query)

	// At this point we have an url with RawQuery and instead of parsing the whole query string,
	// we can search for the key and parse the corresponding value.

	q, err = urlp.GetQueryParam(u.RawQuery, "q")
	// Sometimes keys can be also encoded, so to find a value, you need search with encoded key:
	//q, err = urlp.GetQueryParam(u.RawQuery, url.QueryEscape(`%22q%22`))
	if err != nil {
		panic(err)
	}
	fmt.Printf("q (urlp): %s\n", q)

	// If you sure that there are multiple values behind the same key, you can use GetQueryParamAll function to get all values for the key:

	regionList, err := urlp.GetQueryParamAll(u.RawQuery, "region")
	if err != nil {
		// the error appear only if the function is unable to decode the value.
		panic(err)
	}

	fmt.Println("region list:", regionList)

	// Sometimes, you want want to collect urls, but they are too long, and they contain some not significant (for you) parameters.

	// Instead of this:
	// query.Del("some-hash")
	// u.RawQuery = query.Encode()

	// We can do this:
	urlp.DeleteQueryParam(&u.RawQuery, "some-hash")

	// If we need to delete all query parameters with same key, we can use DeleteQueryParamAll function:
	urlp.DeleteQueryParamAll(&u.RawQuery, "region")

	fmt.Println("URL after delete param:", u.String())

	// To add a new parameter to the query string we can use AddQueryParam function:
	// let's bring back our regions
	urlp.AddQueryParam(&u.RawQuery, "region", "1")
	// It's possible to add multiple values for the same key:
	urlp.AddQueryParam(&u.RawQuery, "region", "2", "3")

	// If we want to get a value and remove it from query immediately, we can use ExtractQueryParam function,
	// to take the first value with the key, and remove it from query string
	region, _ := urlp.ExtractQueryParam(&u.RawQuery, "region")
	fmt.Println("extracted region:", region)
	fmt.Printf("URL after extracting region parameter: %s\n", u.String())

	// It is also possible to extract all values with by with same key:
	regionList, _ = urlp.ExtractQueryParamAll(&u.RawQuery, "region")
	fmt.Println("extracted region list:", regionList)
	fmt.Printf("URL after extracting every region parameter: %s\n", u.String())

	// It is also possible to quickly check presence of the key, without parsing a whole query-string:

	regionOk := urlp.HasQueryParam(u.RawQuery, "region")

	if !regionOk {
		// urlp.SetQueryParam can be used to set a parameter instead of query.Set().
		// It will replace any existing parameters with same key.
		// It will also take place of the first appeared parameter with the same key.
		urlp.SetQueryParam(&u.RawQuery, "region", "12")
		fmt.Println("URL after setting a new region parameter:", u.String())
	}

}
