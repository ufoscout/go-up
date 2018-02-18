package decorator

import (
	"github.com/ufoscout/go-up/reader"
	"github.com/ufoscout/go-up/util"
	"errors"
	"strings"
)

type PlaceholderReplacerDecoratorReader struct {
	Reader                         reader.Reader
	StartDelimiter                 string
	EndDelimiter                   string
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

					tokenValue, tokenFound := output[token]
					if tokenFound {
						if !util.HasTokens(tokenValue.Value, f.StartDelimiter, f.EndDelimiter) {
							value.Value = strings.Replace(value.Value, f.StartDelimiter+token+f.EndDelimiter, tokenValue.Value, -1)
							output[key] = value
							valuesReplacedOnLastLoop = true
						}
					}
				}
			}
		}
	}

	if valuesToBeReplaced && !f.IgnoreUnresolvablePlaceholders {

		message := "Unresolvable placeholders: \n";
		for key, value := range valuesToBeReplacedMap {
			message = message + "key: [" + key + "] value: [" + value.Value + "]\n";
		}

		return nil, errors.New(message)
	}

	return output, nil

}
