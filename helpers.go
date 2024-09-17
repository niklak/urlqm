package urlp

import "strings"

func cutStringByAnySep(s string, sep string) (before string, after string, found bool) {

	if i := strings.IndexAny(s, sep); i >= 0 {
		return s[:i], s[i+1:], true
	}
	return s, "", false
}
