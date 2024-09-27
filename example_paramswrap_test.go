package urlqm

import (
	"fmt"
	"net/url"
)

func ExampleParseQuery() {
	u, err := url.Parse(`https://example.com/?q=100%25+truth&a=1&b=2`)
	if err != nil {
		panic(err)
	}
	params, err := ParseQuery(u.RawQuery)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%+v\n", params)
	// Output:
	// [{Key:q Value:100% truth} {Key:a Value:1} {Key:b Value:2}]
}

func ExampleParams_Encode() {
	u, err := url.Parse(`https://example.com/`)
	if err != nil {
		panic(err)
	}
	params := Params{{"q", "100% truth"}, {"a", "1"}, {"b", "2"}}

	u.RawQuery = params.Encode()

	fmt.Println(u)
	// Output:
	// https://example.com/?q=100%25+truth&a=1&b=2
}

func ExampleParams_Sort() {
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
	// Output:
	// https://example.com/?a=1&b=2&q=100%25+truth
}

func ExampleParams_SetOrder() {
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
	// Output:
	// https://example.com/?q=100%25+truth&d=4&c=3&b=2&a=1
}

func ExampleParams_Add() {
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
	// Output:
	// https://example.com/?q=100%25+truth&a=1&b=2&c=3&c=4

}

func ExampleParams_Set() {
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
	// Output:
	// https://example.com/?q=100%25+truth&a=1&b=5&c=3

}

func ExampleParams_Get() {
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
	// Output:
	// 2
}

func ExampleParams_GetAll() {
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
	// Output:
	// [1 3]
}

func ExampleParams_Extract() {
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
	// Output:
	// 1
	// https://example.com/?q=100%25+truth&b=2&a=3
}

func ExampleParams_ExtractAll() {
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
	// Output:
	// [1 3]
	// https://example.com/?q=100%25+truth&b=2
}

func ExampleParams_Delete() {
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
	// Output:
	// https://example.com/?q=100%25+truth&b=2&a=3
}

func ExampleParams_DeleteAll() {
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
	// Output:
	// https://example.com/?q=100%25+truth&b=2
}

func ExampleParams_Has() {
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
	// Output:
	// true
	// false
}
