package reader

type ProgrammaticReader struct {
	properties map[string]Property
}

func (f *ProgrammaticReader) Read() (map[string]Property, error) {
	return f.properties, nil
}

func (f *ProgrammaticReader) Add(key string, value string) *ProgrammaticReader {
	f.AddProperty(key, Property{value, true})
	return f
}

func (f *ProgrammaticReader) AddProperty(key string, value Property) *ProgrammaticReader {
	f.properties[key] = value
	return f
}

func NewProgrammaticReader() *ProgrammaticReader {
	return &ProgrammaticReader{map[string]Property{}}
}