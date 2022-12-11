package claxon

// Using this structure allows the data and metadata to be kept
// separate (and leverages Go generics) but when it comes time to
// write the data out, it will inline the metadata.  This is
// especially useful if you want to nest hypermedia metadata
// in hierarchical data.
type Emb[T any] struct {
	Data    T
	Context Claxon
}

func Embed[T any](data T, metadata Claxon) Emb[T] {
	return Emb[T]{
		Data:    data,
		Context: metadata,
	}
}

func (e Emb[T]) MarshalJSON() ([]byte, error) {
	return InlineMarshal(e.Data, e.Context)
}
