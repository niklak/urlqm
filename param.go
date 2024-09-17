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
func ParseParams(query string) (values []QueryParam, err error) {

	values = make([]QueryParam, 0)
	for query != "" {
		var key string
		key, query, _ = cutStringByAnySep(query, "&;")
		if key == "" {
			continue
		}
		key, value, _ := strings.Cut(key, "=")
		key, err1 := url.QueryUnescape(key)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		value, err1 = url.QueryUnescape(value)
		if err1 != nil {
			if err == nil {
				err = err1
			}
			continue
		}
		values = append(values, QueryParam{Key: key, Value: value})
	}
	return
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

// PopParam removes and returns the value of a parameter from query string.
func PopParam(query *string, key string) (value string) {
	before, after, found := strings.Cut(*query, key+"=")
	//if the given param wasn't found or it was without value
	if !found || after == "" {
		return
	}

	before = strings.TrimSuffix(before, "&")

	buf := strings.Builder{}
	buf.WriteString(before)

	value, after, _ = strings.Cut(after, "&")
	if buf.Len() > 0 && after != "" {
		buf.WriteString("&")
	}
	buf.WriteString(after)

	*query = buf.String()
	return
}

// GetParam returns the value of a parameter from query string.
func GetParam(query string, key string) (value string) {
	_, after, found := strings.Cut(query, key+"=")
	//if the given param wasn't found or it was without value
	if !found || after == "" {
		return
	}

	value, _, _ = strings.Cut(after, "&")
	return
}
