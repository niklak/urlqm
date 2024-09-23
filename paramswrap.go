package urlp

// Params is a wrapper around []Param
type Params []Param

// Encode transforms []Param into an url-encoded string and returns it.
func (p *Params) Encode() (encode string) {
	return EncodeParams(*p)
}

// Sort sorts the params by key in ascending order. Same as [SortParams].
func (p *Params) Sort() {
	SortParams(*p)
}

// SetOrder sets the order for the params. Same as [SortOrderParams].
func (p *Params) SetOrder(order ...string) {
	SortOrderParams((*[]Param)(p), order...)
}

// Add adds a new param to the params slice.
func (p *Params) Add(key, value string) {
	*p = append(*p, Param{key, value})
}

// Get returns the first value of a param with given key or an empty string if not found.
func (p *Params) Get(key string) string {
	for _, param := range *p {
		if param.Key == key {
			return param.Value
		}
	}
	return ""
}

// GetAll returns all values for the given key, if not found returns nil.
func (p *Params) GetAll(key string) []string {
	var values []string
	for _, param := range *p {
		if param.Key == key {
			values = append(values, param.Value)
		}
	}
	return values
}

// Extract removes the first param with given key from the params slice and returns its value.
func (p *Params) Extract(key string) (value string) {
	for i := 0; i < len(*p); i++ {
		if (*p)[i].Key == key {
			value = (*p)[i].Value
			*p = append((*p)[:i], (*p)[i+1:]...)
			return value
		}
	}
	return ""
}

// ExtractAll removes all params with given key from the params slice and returns their values.
func (p *Params) ExtractAll(key string) (values []string) {
	for i := 0; i < len(*p); i++ {
		if (*p)[i].Key == key {
			values = append(values, (*p)[i].Value)
			*p = append((*p)[:i], (*p)[i+1:]...)
		}
	}
	return values
}

// Set sets a param with given key and value.
// It replaces any existing param with the same key.
func (p *Params) Set(key, value string) {
	if len(key) == 0 {
		return
	}

	foundIdx := -1
	for i := 0; i < len(*p); i++ {
		if (*p)[i].Key == key {
			if foundIdx == -1 {
				// remember the first index we found.
				foundIdx = i
				continue
			}
			// in other case simply remove the param.
			*p = append((*p)[:i], (*p)[i+1:]...)
		}
	}

	if foundIdx > -1 {
		(*p)[foundIdx] = Param{Key: key, Value: value}
		return
	}
	*p = append(*p, Param{Key: key, Value: value})

}

// QueryParams takes a query string and returns a slice of [Param]. See [ParseParams].
func QueryParams(rawQuery string) (p Params, err error) {
	return ParseParams(rawQuery)
}

// Delete removes the first param with given key from the params slice and returns its value.
func (p *Params) Delete(key string) {
	for i := 0; i < len(*p); i++ {
		if (*p)[i].Key == key {
			*p = append((*p)[:i], (*p)[i+1:]...)
			return
		}
	}
}

// DeleteAll removes all params with given key from the params slice
func (p *Params) DeleteAll(key string) {
	for i := 0; i < len(*p); i++ {
		if (*p)[i].Key == key {
			*p = append((*p)[:i], (*p)[i+1:]...)
		}
	}
}

// Has returns true if the params slice contains a param with given key.
func (p Params) Has(key string) bool {
	for _, param := range p {
		if param.Key == key {
			return true
		}
	}
	return false
}
