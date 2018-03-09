package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldNotReturnNil(t *testing.T) {

	var reader Reader = NewProgrammaticReader()

	properties, error := reader.Read()

	assert.NotNil(t, properties)
	assert.Nil(t, error)
}

func Test_ShouldAddProperties(t *testing.T) {

	reader := NewProgrammaticReader()
	reader.Add("KEY", "VALUE")

	properties, error := reader.Read()

	assert.NotNil(t, properties)
	assert.Nil(t, error)

	assert.Equal(t, "VALUE", properties["KEY"].Value)
	assert.True(t, properties["KEY"].Resolvable)
}
