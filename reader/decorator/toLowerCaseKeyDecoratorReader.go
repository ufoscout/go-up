package decorator

import (
	"github.com/ufoscout/go-up/reader"
	"strings"
)

type ToLowerCaseKeyDecoratorReader struct {
	Reader reader.Reader
}

func (f *ToLowerCaseKeyDecoratorReader) Read() (map[string]reader.Property, error) {

	result, err := f.Reader.Read()

	if err != nil {
		return nil, err
	}

	output := map[string]reader.Property{}

	for k, v := range result {
		output[strings.ToLower(k)] = v
	}

	return output, nil

}
