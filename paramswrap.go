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

// QueryParams takes a query string and returns a slice of [Param]. See [ParseParams].
func QueryParams(rawQuery string) (p Params, err error) {
	return ParseParams(rawQuery)
}
