package reader

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_ShouldReturnUnresolvableEntries(t *testing.T) {

	var reader Reader = &EnvReader{""}
	prop, err := reader.Read()
	assert.NotNil(t, prop)
	assert.Nil(t, err)

	for key, value := range prop {
		assert.True(t, len(key) > 0)
		//assert.True(t, len(value.Value) > 0)
		assert.False(t, value.Resolvable)
	}

}

func Test_ShouldReturnFilteredByPrefix(t *testing.T) {

	os.Setenv("CUSTOM", "100")
	os.Setenv("PREFIX_CUSTOM", "200")
	os.Setenv("OTHER_PREFIX_CUSTOM", "300")

	var reader Reader = &EnvReader{"PREFIX_"}
	prop, _ := reader.Read()

	custom, customFound := prop["CUSTOM"]
	assert.True(t, customFound)
	assert.Equal(t, "200", custom.Value)

	_, otherFound := prop["OTHER_PREFIX_CUSTOM"]
	assert.False(t, otherFound)
}

func Test_ShouldIdentifyValuesWithEqualsSign(t *testing.T) {

	os.Setenv("CUSTOM", "key=value")

	var reader Reader = &EnvReader{""}
	prop, _ := reader.Read()

	custom, customFound := prop["CUSTOM"]
	assert.True(t, customFound)
	assert.Equal(t, "key=value", custom.Value)

}
