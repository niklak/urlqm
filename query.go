package urlp

import (
	"net/url"
	"strings"
)

// ExtractQueryParam removes and returns the value of a parameter from the query string.
// If the given key contains non-ASCII characters, it must be url-encoded before calling this function.
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
// If the given key contains non-ASCII characters, it must be url-encoded before calling this function.
func ExtractQueryParamAll(query *string, key string) (values []string, err error) {
	var after string = *query

	buf := strings.Builder{}
	found := false

	for after != "" {
		var before string
		before, after, _ = strings.Cut(after, key+"=")
		if after == "" {
			buf.WriteString(before)
			break
		}
		found = true
		before, sep := trimParamSeparator(before)
		buf.WriteString(before)

		var value string
		value, after = cutStringByAnySep(after, separators)
		value, err = url.QueryUnescape(value)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
		if buf.Len() > 0 && !strings.HasPrefix(after, key) {
			buf.WriteString(sep)
		}
	}

	if found {
		*query = buf.String()
	}

	return
}

// GetQueryParam returns the value of a parameter from the query string.
// If the given key contains non-ASCII characters, it must be url-encoded before calling this function.
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
// If the given key contains non-ASCII characters, it must be url-encoded before calling this function.
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
// This function accepts multiple values for a single param.
func AddQueryParam(query *string, key string, values ...string) {

	if key == "" {
		return
	}
	buf := strings.Builder{}

	if len(*query) > 0 {
		buf.WriteString(*query)
		buf.WriteByte('&')
	}
	writeParam(&buf, "&", key, values...)
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

		before, sep := trimParamSeparator(before)

		buf.WriteString(before)
		if after == "" {
			break
		}

		_, after = cutStringByAnySep(after, separators)

		if buf.Len() > 0 && after != "" {
			buf.WriteString(sep)
		}
		if !found {
			found = true
			writeParam(&buf, sep, key, value)
			if after != "" {

				buf.WriteString(sep)
			}
		}
	}

	if !found {
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		writeParam(&buf, "&", key, value)
	}

	*query = buf.String()
}

// DeleteQueryParam removes a parameter from the query string by key.
// If the given key contains non-ASCII characters, it must be url-encoded before calling this function.
func DeleteQueryParam(query *string, key string) {
	before, after, _ := strings.Cut(*query, key+"=")
	//if the given param wasn't found or it was without value
	if after == "" {
		return
	}

	var sep string
	before, sep = trimParamSeparator(before)
	buf := strings.Builder{}
	buf.WriteString(before)

	_, after = cutStringByAnySep(after, separators)

	if buf.Len() > 0 && after != "" {
		buf.WriteString(sep)
	}
	buf.WriteString(after)

	*query = buf.String()
}

// DeleteQueryParamAll removes all parameters from the query string by key.
// If the given key contains non-ASCII characters, it must be url-encoded before calling this function.
func DeleteQueryParamAll(query *string, key string) {
	var after string = *query

	buf := strings.Builder{}
	found := false

	for after != "" {
		var before string
		before, after, _ = strings.Cut(after, key+"=")
		if after == "" {
			buf.WriteString(before)
			break
		}
		found = true
		before, sep := trimParamSeparator(before)

		buf.WriteString(before)
		_, after = cutStringByAnySep(after, separators)

		if buf.Len() > 0 && !strings.HasPrefix(after, key) {
			buf.WriteString(sep)
		}

	}

	if found {
		*query = buf.String()
	}
}

// HasQueryParam returns true if the query string contains a parameter with the given key.
// If the given key contains non-ASCII characters, it must be url-encoded before calling this function.
func HasQueryParam(query string, key string) bool {
	return strings.Contains(query, key+"=")
}

func writeParam(buf *strings.Builder, sep, key string, values ...string) {
	buf.WriteString(url.QueryEscape(key))
	buf.WriteByte('=')
	if len(values) == 0 {
		return
	}

	buf.WriteString(url.QueryEscape(values[0]))

	if len(values) > 1 {

		for _, value := range values[1:] {
			buf.WriteString(sep)
			buf.WriteString(url.QueryEscape(key))
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(value))
		}
	}
}

// TODO: github actions
// TODO: pick up a package name
