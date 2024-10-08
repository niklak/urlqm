// Package urlqm implements an alternative approach to handle url query parameters.
package urlqm

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

	if len(params) == 0 {
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
func ParseParams(query string) ([]Param, error) {
	var err error

	if query == "" {
		return nil, err
	}

	estLen := strings.Count(query, "&") + strings.Count(query, ";") + 1
	params := make([]Param, 0, estLen)
	for query != "" {
		var key, value string
		key, query = cutStringByAnySep(query, separators)
		if key == "" {
			continue
		}
		key, value, _ = strings.Cut(key, "=")
		// TODO: alloc!
		if key1, err1 := url.QueryUnescape(key); err1 == nil {
			key = key1
		} else {
			err = errorMerge(err, err1)
		}

		// TODO: alloc!
		if value1, err1 := url.QueryUnescape(value); err1 == nil {
			value = value1
		} else {
			err = errorMerge(err, err1)
		}

		params = append(params, Param{Key: key, Value: value})
	}

	return params, err
}

// SortOrderParams sorts the Param slice
// based on the provided order members while omitted params are placed as it was.
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
