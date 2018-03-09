package decorator

import (
	"github.com/stretchr/testify/assert"
	"github.com/ufoscout/go-up/reader"
	"testing"
)

func Test_ShouldReplaceStringsInTheKeys(t *testing.T) {

	properties := reader.NewProgrammaticReader()
	properties.AddProperty("KEY_one", reader.Property{"ENV_ONE", false})
	properties.Add("KEY_TWO_", "VALUE_TWO.")
	properties.Add("key.three", "value.three")

	decorator := &KeyStringReplacerDecoratorReader{properties, "_", "."}
	prop, _ := decorator.Read()

	assert.Equal(t, "ENV_ONE", prop["KEY.one"].Value)
	assert.Equal(t, "VALUE_TWO.", prop["KEY.TWO."].Value)
	assert.Equal(t, "value.three", prop["key.three"].Value)

}
