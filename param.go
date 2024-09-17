package urlp

import (
	"net/url"
	"strings"
)

type QueryParam struct {
	Key, Value string
}

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

func ParseParams(query string) (values []QueryParam, err error) {
	values = make([]QueryParam, 0)
	for query != "" {
		var key string
		key, query, _ = strings.Cut(query, "&")

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

func SortParams(paramsPtr *[]QueryParam, order []string) {

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

func PopParam(rawQuery *string, key string) (value string) {
	before, after, found := strings.Cut(*rawQuery, key+"=")
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

	*rawQuery = buf.String()
	return
}

func GetParam(rawQuery string, key string) (value string) {
	_, after, found := strings.Cut(rawQuery, key+"=")
	//if the given param wasn't found or it was without value
	if !found || after == "" {
		return
	}

	value, _, _ = strings.Cut(after, "&")
	return
}
