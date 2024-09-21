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

// Remove removes the first param with given key from the params slice and returns its value.
func (p *Params) Remove(key string) (value string) {
	for i := 0; i < len(*p); i++ {
		if (*p)[i].Key == key {
			value = (*p)[i].Value
			*p = append((*p)[:i], (*p)[i+1:]...)
			return value
		}
	}
	return ""
}

// RemoveAll removes all params with given key from the params slice and returns their values.
func (p *Params) RemoveAll(key string) (values []string) {
	for i := 0; i < len(*p); i++ {
		if (*p)[i].Key == key {
			values = append(values, (*p)[i].Value)
			*p = append((*p)[:i], (*p)[i+1:]...)
		}
	}
	return values
}

// QueryParams takes a query string and returns a slice of [Param]. See [ParseParams].
func QueryParams(rawQuery string) (p Params, err error) {
	return ParseParams(rawQuery)
}
