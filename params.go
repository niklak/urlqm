// Package urlp implements an alternative approach to handle url query parameters.
package urlp

import (
	"net/url"
	"sort"
	"strings"
)

// Param represents a key-value pair in a URL query string.
type Param struct {
	Key, Value string
}

// EncodeParams takes a slice of Param and returns the encoded query string.
func EncodeParams(params []Param) string {
	if params == nil {
		return ""
	}
	var buf strings.Builder

	for _, param := range params {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(url.QueryEscape(param.Key))
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(param.Value))
	}
	return buf.String()
}

// ParseParams takes a query string and returns a slice of Param.
// Unlike `url.ParseQuery`, this function collects unescaped keys and values if it fails to unescape them.
// Also it collects errors from `url.QueryUnescape`. It can be checked with [errors.As] or `err != nil`.
// If error is not `nil`, it contains all occurred errors.
//
// ## Example:
//
//	params, err := urlp.ParseParams(`a=1&q=100%+truth&b=2&brightness=90%`)
//	if err != nil {
//	    fmt.Println("Error:", err)
//	    // handling this error, keep in mind that params is not empty, it contains all keys and values, both escaped and unescaped
//	}
//	var e url.EscapeError
//	if errors.As(err, &e) {
//	    fmt.Printf("Error: %v\n", e)
//	}
//	fmt.Printf("%+v\n", params) // outputs: [{a 1} {q 100%+truth} {b 2} {brightness 90%}]
func ParseParams(query string) ([]Param, error) {
	var err error
	params := make([]Param, 0)
	for query != "" {
		var key, value string
		key, query = cutStringByAnySep(query, separators)
		if key == "" {
			continue
		}
		key, value, _ = strings.Cut(key, "=")

		if key1, err1 := url.QueryUnescape(key); err1 == nil {
			key = key1
		} else {
			err = errorMerge(err, err1)
		}

		if value1, err1 := url.QueryUnescape(value); err1 == nil {
			value = value1
		} else {
			err = errorMerge(err, err1)
		}

		params = append(params, Param{Key: key, Value: value})
	}

	return params, err
}

// SortOrderParams sorts the Param slice based on the provided order members while omitted params are placed as it was.
func SortOrderParams(paramsPtr *[]Param, order ...string) {

	restParams := *paramsPtr
	ordered := make([]Param, 0, len(restParams))

	for _, k := range order {
		for i := 0; i < len(restParams); i++ {
			if k == restParams[i].Key {
				ordered = append(ordered, restParams[i])
				restParams = append(restParams[:i], restParams[i+1:]...)
			}
		}
	}

	ordered = append(ordered, restParams...)
	*paramsPtr = ordered
}

// SortParams sorts the slice of Param by key in ascending order.
func SortParams(params []Param) {
	sort.Slice(params, func(i, j int) bool {
		return params[i].Key < params[j].Key
	})
}

// TODO: probably add a container
// TODO: readme
// TODO: some benchmarks
