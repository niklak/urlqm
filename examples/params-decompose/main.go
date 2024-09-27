package main

import (
	"fmt"
	"net/url"

	"github.com/niklak/urlqm"
)

func main() {

	var rawURL = "https://www.example.com/?q=a+testing+query&region=1&region=3&region=5" +
		"&some-hash=EgZjaHJvbWUyBggAEEUYOTIGCAEQABhAMgYIAhAAGEAyBggD" +
		"EAAYQDIGCAQQABhAMgYIBRAAGEAyBggGAYQDIGCAcQABhAMgYICBAAGEDSAQgxMjQxajBqMagCArACAQ&first=11&brightness=90%"

	u, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}
	fmt.Printf("URL: %s\n", u)

	query, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		// got error here, now if you wish you can handle it.
		fmt.Printf("Error parsing query: %s\n", err)
	}
	// broken parameter `brightness` is dropped
	fmt.Printf("Parsed Query: %#v\n", query)

	// broken parameter `brightness` is kept as is
	queryParams, err := urlqm.ParseQuery(u.RawQuery)

	if err != nil {
		fmt.Printf("Error parsing query params: %s\n", err)
	}

	// When you encode parameters with std net/url, your url will be never same again.
	// because it sorts query parameters just before returning an url-encoded string.
	u.RawQuery = query.Encode()
	fmt.Printf("URL after std query encoding -- changed order: %s\n", u)

	// `urlqm.Params.Encode` will keep the original order, and encodes bad parameters
	u.RawQuery = queryParams.Encode()
	fmt.Printf("URL after urlqm params encoding: %s\n", u)

	// With urlqm you have a choice to sort or not to sort query parameters.

	// Simple sort by key
	queryParams.Sort()
	u.RawQuery = queryParams.Encode()
	fmt.Printf("URL after urlqm params sorting and encoding: %s\n", u)

	//moving one parameter to the start
	queryParams.SetOrder("q")

	// Or setting explicit order
	queryParams.SetOrder("q", "region", "region", "region", "some-hash", "first", "brightness")
	u.RawQuery = queryParams.Encode()
	fmt.Printf("URL after urlqm params setting order and encoding: %s\n", u)

	// You can get parameter value
	q := queryParams.Get("q")

	fmt.Printf("q (urlqm params): %s\n", q)

	// Or get a list of parameter values
	regionList := queryParams.GetAll("region")

	fmt.Println("region list:", regionList)

	// A single parameter's value can be deleted with:
	queryParams.Delete("some-hash")
	// Every parameter with the key can be deleted with:
	queryParams.DeleteAll("region")

	u.RawQuery = queryParams.Encode()

	fmt.Println("URL after delete param:", u.String())

	// Add a single parameter to the end of the parameter list.
	queryParams.Add("region", "1")
	// Or add a parameter with multiple values
	queryParams.Add("region", "2", "3")

	// Take a single parameter's value by key and remove it from the parameter list
	region := queryParams.Extract("region")

	fmt.Println("extracted region:", region)

	u.RawQuery = queryParams.Encode()
	fmt.Printf("URL after extracting region parameter: %s\n", u.String())

	// Take all parameter's values by key and remove them from the parameter list
	regionList = queryParams.ExtractAll("region")
	fmt.Println("extracted region list:", regionList)

	u.RawQuery = queryParams.Encode()
	fmt.Printf("URL after extracting every region parameter: %s\n", u.String())

	// Checking presence of a parameter in the parameter list.
	regionOk := queryParams.Has("region")

	if !regionOk {
		// Setting a parameter.
		// It will replace any existing parameters with same key.
		// It will also take place of the first appeared parameter with the same key.
		queryParams.Set("region", "12")
		u.RawQuery = queryParams.Encode()
		fmt.Println("URL after setting a new region parameter:", u.String())
	}

}
