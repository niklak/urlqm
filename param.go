// Package urlp implements an alternative approach to handle url query parameters.
package urlp

import (
	"net/url"
	"strings"
)

// QueryParam represents a key-value pair in a URL query string.
type QueryParam struct {
	Key, Value string
}

// EncodeParams takes a slice of QueryParam and returns the encoded query string.
func EncodeParams(params []QueryParam) string {
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

// ParseParams takes a query string and returns a slice of QueryParam.
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
func ParseParams(query string) ([]QueryParam, error) {
	var err error
	params := make([]QueryParam, 0)
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

		params = append(params, QueryParam{Key: key, Value: value})
	}

	return params, err
}

// SortOrderParams sorts the QueryParam slice based on the provided order members while omitted params are placed as it was.
func SortOrderParams(paramsPtr *[]QueryParam, order ...string) {

	restParams := *paramsPtr
	ordered := make([]QueryParam, 0, len(restParams))

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

// PopParam removes and returns the value of a parameter from the query string.
func PopParam(query *string, key string) (value string, err error) {
	before, after, _ := strings.Cut(*query, key+"=")
	//if the given param wasn't found or it was without value
	if after == "" {
		return
	}

	var sep string

	before, sep = trimParamSeparator(before)
	buf := strings.Builder{}
	buf.WriteString(before)

	value, after = cutStringByAnySep(after, separators)
	value, err = url.QueryUnescape(value)
	if err != nil {
		return
	}

	if buf.Len() > 0 && after != "" {
		buf.WriteString(sep)
	}
	buf.WriteString(after)

	*query = buf.String()
	return
}

// PopParamValues removes and returns a slice of values for a parameter from the query string.
func PopParamValues(query *string, key string) (values []string, err error) {
	values = make([]string, 0)
	var after string = *query

	buf := strings.Builder{}

	for after != "" {
		var before string
		before, after, _ = strings.Cut(after, key+"=")
		if after == "" {
			break
		}

		before, sep := trimParamSeparator(before)

		buf.WriteString(before)
		var value string
		value, after = cutStringByAnySep(after, separators)
		value, err = url.QueryUnescape(value)
		if err != nil {
			return []string{}, err
		}
		values = append(values, value)
		if buf.Len() > 0 && after != "" {
			buf.WriteString(sep)
		}
		buf.WriteString(after)
	}

	*query = buf.String()

	return
}

// GetParam returns the value of a parameter from the query string.
func GetParam(query string, key string) (value string, err error) {
	_, after, _ := strings.Cut(query, key+"=")
	//if the given param wasn't found or it was without value
	if after == "" {
		return
	}

	value, _ = cutStringByAnySep(after, separators)
	value, err = url.QueryUnescape(value)
	return
}

// GetParamValues returns the slice of values for a parameter from the query string.
func GetParamValues(query, key string) (values []string, err error) {

	values = make([]string, 0)
	for query != "" {
		_, query, _ = strings.Cut(query, key+"=")
		if query == "" {
			break
		}
		value, _ := cutStringByAnySep(query, separators)
		if value, err = url.QueryUnescape(value); err != nil {
			return []string{}, err
		}
		values = append(values, value)
	}

	return
}

// TODO: probably add a container
// TODO: readme
// TODO: some benchmarks
