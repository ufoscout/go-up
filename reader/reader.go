package reader

type Property struct {
	Value string
	Resolvable bool
}

type Reader interface {
	Read() (map[string]Property, error)
}
