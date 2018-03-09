package decorator

import (
	"strings"

	"github.com/ufoscout/go-up/reader"
)

// ToLowerCaseKeyDecoratorReader transforms all property keys to lowercase
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
