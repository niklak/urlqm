package urlp

import (
	"net/url"
	"strings"
)

// ExtractQueryParam removes and returns the value of a parameter from the query string.
func ExtractQueryParam(query *string, key string) (value string, err error) {
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

// ExtractQueryParamAll removes and returns a slice of values for a parameter from the query string.
func ExtractQueryParamAll(query *string, key string) (values []string, err error) {
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
			return nil, err
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

// GetQueryParamAll returns the slice of values for a parameter from the query string.
func GetQueryParamAll(query, key string) (values []string, err error) {

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

// SetQueryParam sets a parameter in the query string.
func SetQueryParam(query *string, key string, value string) {
	if key == "" {
		return
	}

	found := false
	var after string = *query

	buf := strings.Builder{}

	for after != "" {
		var before string
		before, after, _ = strings.Cut(after, key+"=")
		if after == "" {
			buf.WriteString(before)
			break
		}

		before, sep := trimParamSeparator(before)
		buf.WriteString(before)

		_, after = cutStringByAnySep(after, separators)

		if buf.Len() > 0 {
			buf.WriteString(sep)
		}
		if !found {
			found = true
			buf.WriteString(url.QueryEscape(key))
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(value))
			if after != "" {
				buf.WriteString(sep)
			}
		}
	}

	if !found {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(url.QueryEscape(key))
		buf.WriteByte('=')
		buf.WriteString(url.QueryEscape(value))
	}

	*query = buf.String()
}
