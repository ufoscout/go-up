package decorator

import (
	"github.com/ufoscout/go-up/reader"
	"github.com/ufoscout/go-up/util"
	"errors"
	"strings"
)

type ReplacerDecoratorReader struct {
	reader                         reader.Reader
	startDelimiter                 string
	endDelimiter                   string
	ignoreUnresolvablePlaceholders bool
}

func (f *ReplacerDecoratorReader) Read() (map[string]reader.Property, error) {

	result, err := f.reader.Read()

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

				tokens := util.AllTokensDistinct(value.Value, f.startDelimiter, f.endDelimiter, true)
				if len(tokens) > 0 {
					valuesToBeReplaced = true
					valuesToBeReplacedMap[key] = value
				}

				for _, token := range tokens {

					tokenValue, tokenFound := output[token]
					if tokenFound {
						if !util.HasTokens(tokenValue.Value, f.startDelimiter, f.endDelimiter) {
							value.Value = strings.Replace(value.Value, f.startDelimiter+token+f.endDelimiter, tokenValue.Value, -1)
							output[key] = value
							valuesReplacedOnLastLoop = true
						}
					}
				}
			}
		}
	}

	if valuesToBeReplaced && !f.ignoreUnresolvablePlaceholders {

		message := "Unresolvable placeholders: \n";
		for key, value := range valuesToBeReplacedMap {
			message = message + "key: [" + key + "] value: [" + value.Value + "]\n";
		}

		return nil, errors.New(message)
	}

	return output, nil

}
