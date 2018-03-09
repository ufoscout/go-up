package reader

import (
	"bufio"
	"os"
	"strings"
)

// FileReader reads properties from a file
type FileReader struct {
	Filename       string
	IgnoreNotFound bool
}

func (f *FileReader) Read() (map[string]Property, error) {

	config := map[string]Property{}

	file, err := os.Open(f.Filename)
	if err != nil {
		if f.IgnoreNotFound {
			return config, nil
		}
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				if !keyToBeIgnored(key) {
					value := ""
					if len(line) > equal {
						value = strings.TrimSpace(line[equal+1:])
					}
					config[key] = Property{value, true}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}

func keyToBeIgnored(key string) bool {
	return strings.HasPrefix(key, "#")
}
