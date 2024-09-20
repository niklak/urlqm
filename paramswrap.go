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

// QueryParams takes a query string and returns a slice of [Param]. See [ParseParams].
func QueryParams(rawQuery string) (p Params, err error) {
	return ParseParams(rawQuery)
}

// TODO: add Get
// TODO: add GetAll
// TODO: add Pop
// TODO: add PopAll
