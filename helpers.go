package urlp

import (
	"strings"
)

var paramSep = "&"
var deprecatedParamSep = ";"
var separators = "&;"

func cutStringByAnySep(s string, seps string) (
	before string, after string) {

	if i := strings.IndexAny(s, seps); i >= 0 {
		return s[:i], s[i+1:]
	}
	return s, ""
}

func trimParamSeparator(s string) (string, string) {

	if strings.HasSuffix(s, paramSep) || strings.HasSuffix(s, deprecatedParamSep) {
		// length of suffix is always 1
		return s[:len(s)-1], s[len(s)-1:]
	}

	return s, ""
}
