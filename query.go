package urlp

import (
	"net/url"
	"strings"
)

// PopQueryParam removes and returns the value of a parameter from the query string.
func PopQueryParam(query *string, key string) (value string, err error) {
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

// PopQueryParamValues removes and returns a slice of values for a parameter from the query string.
func PopQueryParamValues(query *string, key string) (values []string, err error) {
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

// GetQueryParam returns the value of a parameter from the query string.
func GetQueryParam(query string, key string) (value string, err error) {
	_, after, _ := strings.Cut(query, key+"=")
	//if the given param wasn't found or it was without value
	if after == "" {
		return
	}

	value, _ = cutStringByAnySep(after, separators)
	value, err = url.QueryUnescape(value)
	return
}

// GetQueryParamValues returns the slice of values for a parameter from the query string.
func GetQueryParamValues(query, key string) (values []string, err error) {

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

// AddQueryParam adds a parameter to the query string.
func AddQueryParam(query *string, key string, value string) {

	if key == "" {
		return
	}
	buf := strings.Builder{}

	if len(*query) > 0 {
		buf.WriteString(*query)
		buf.WriteByte('&')
	}
	buf.WriteString(url.QueryEscape(key))
	buf.WriteByte('=')
	buf.WriteString(url.QueryEscape(value))
	*query = buf.String()
}
