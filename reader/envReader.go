package reader

import (
	"os"
	"strings"
)

// EnvReader creates properties from the environment variables
type EnvReader struct {
	Prefix string
}

func (f *EnvReader) Read() (map[string]Property, error) {

	config := map[string]Property{}

	// fetch all env variables
	for _, element := range os.Environ() {
		variable := strings.SplitN(element, "=", 2)
		if strings.HasPrefix(variable[0], f.Prefix) {
			key := variable[0][len(f.Prefix):]
			config[key] = Property{variable[1], false}
		}
	}

	return config, nil
}
