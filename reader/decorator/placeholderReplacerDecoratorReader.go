package decorator

import (
	"errors"
	"strings"

	"github.com/ufoscout/go-up/reader"
	"github.com/ufoscout/go-up/util"
)

type PlaceholderReplacerDecoratorReader struct {
	Reader                         reader.Reader
	StartDelimiter                 string
	EndDelimiter                   string
	DefaultValueSeparator          string
	IgnoreUnresolvablePlaceholders bool
}

func (f *PlaceholderReplacerDecoratorReader) Read() (map[string]reader.Property, error) {

	result, err := f.Reader.Read()

	if err != nil {
		return nil, err
	}

	output := map[string]reader.Property{}

	for k, v := range result {
		output[k] = v
	}

	valuesToBeReplacedMap := map[string]reader.Property{}
	valuesToBeReplaced := true
	valuesReplacedOnLastLoop := true

	for valuesReplacedOnLastLoop && valuesToBeReplaced {

		valuesToBeReplaced = false
		valuesReplacedOnLastLoop = false
		valuesToBeReplacedMap = map[string]reader.Property{}

		for key, value := range output {

			if value.Resolvable {

				tokens := util.AllTokensDistinct(value.Value, f.StartDelimiter, f.EndDelimiter, true)
				if len(tokens) > 0 {
					valuesToBeReplaced = true
					valuesToBeReplacedMap[key] = value
				}

				for _, token := range tokens {

					tokenValue, tokenFound := output[getBaseValue(token, f.DefaultValueSeparator)]

					if tokenFound {
						if !util.HasTokens(tokenValue.Value, f.StartDelimiter, f.EndDelimiter) {
							value.Value = strings.Replace(value.Value, f.StartDelimiter+token+f.EndDelimiter, tokenValue.Value, -1)
							output[key] = value
							valuesReplacedOnLastLoop = true
						}
					} else if hasDefaultValue(token, f.DefaultValueSeparator) {
						value.Value = getDefaultValue(token, f.DefaultValueSeparator)
						output[key] = value
						valuesReplacedOnLastLoop = true
					}
				}
			}
		}
	}

	if len(valuesToBeReplacedMap) > 0 && !f.IgnoreUnresolvablePlaceholders {

		message := "Unresolvable placeholders: \n"
		for key, value := range valuesToBeReplacedMap {
			message = message + "key: [" + key + "] value: [" + value.Value + "]\n"
		}

		return nil, errors.New(message)
	}

	return output, nil

}

func hasDefaultValue(token string, defaultValueSeparator string) bool {
	return strings.Index(token, defaultValueSeparator) >= 0
}

func getBaseValue(token string, defaultValueSeparator string) string {
	index := strings.Index(token, defaultValueSeparator)
	if index >= 0 {
		return token[:index]
	}
	return token
}

func getDefaultValue(token string, defaultValueSeparator string) string {
	index := strings.Index(token, defaultValueSeparator)
	if index >= 0 {
		return token[index+1:]
	}
	return ""
}
