package reader

// ProgrammaticReader permits to define properties programmatically
type ProgrammaticReader struct {
	properties map[string]Property
}

func (f *ProgrammaticReader) Read() (map[string]Property, error) {
	return f.properties, nil
}

// Add a new property
func (f *ProgrammaticReader) Add(key string, value string) *ProgrammaticReader {
	f.AddProperty(key, Property{value, true})
	return f
}

// AddProperty adds a new property
func (f *ProgrammaticReader) AddProperty(key string, value Property) *ProgrammaticReader {
	f.properties[key] = value
	return f
}

// NewProgrammaticReader creats a new ProgrammaticReader instance
func NewProgrammaticReader() *ProgrammaticReader {
	return &ProgrammaticReader{map[string]Property{}}
}
