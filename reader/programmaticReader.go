package reader

type ProgrammaticReader struct {
	properties map[string]Property
}

func (f *ProgrammaticReader) Read() (map[string]Property, error) {
	return f.properties, nil
}

func (f *ProgrammaticReader) Add(key string, value string) {
	f.AddProperty(key, Property{value, true})
}

func (f *ProgrammaticReader) AddProperty(key string, value Property) {
	f.properties[key] = value
}

func NewProgrammaticReader() *ProgrammaticReader {
	return &ProgrammaticReader{map[string]Property{}}
}