package reader

// Property represents a property key-value pair
type Property struct {
	Value      string
	Resolvable bool
}

// Reader creates a properies map
type Reader interface {
	Read() (map[string]Property, error)
}
