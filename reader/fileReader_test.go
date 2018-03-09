package reader

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ShouldThrowExceptionWhenFileNotFound(t *testing.T) {
	ignoreNotFound := false
	filename := "./some_name"
	var reader Reader = &FileReader{filename, ignoreNotFound}
	prop, err := reader.Read()
	assert.Nil(t, prop)
	assert.NotNil(t, err)
}

func Test_ShouldIgnoreFileNotFoundIfRequested(t *testing.T) {
	ignoreNotFound := true
	filename := "./some_name"
	var reader Reader = &FileReader{filename, ignoreNotFound}

	prop, err := reader.Read()
	assert.NotNil(t, prop)
	assert.Nil(t, err)

}

func Test_ShouldReadProperiesFromFile(t *testing.T) {
	ignoreNotFound := false
	filename := "../test/files/test1.properties"
	var reader Reader = &FileReader{filename, ignoreNotFound}
	prop, err := reader.Read()
	assert.NotNil(t, prop)
	assert.Nil(t, err)

	_, found := prop["keyOne"]
	assert.True(t, found)
	assert.Equal(t, "firstvalue", prop["keyOne"].Value)
	assert.Equal(t, "second VALUE", prop["keyTwo"].Value)
}

func Test_ShouldIdentifyValuesWithEqualsInTheValue(t *testing.T) {

	ignoreNotFound := false
	filename := "../test/files/test1.properties"
	var reader Reader = &FileReader{filename, ignoreNotFound}
	prop, err := reader.Read()
	assert.NotNil(t, prop)
	assert.Nil(t, err)

	custom, customFound := prop["key.with_equal"]
	assert.True(t, customFound)
	assert.Equal(t, "key=value", custom.Value)

}

func Test_ShouldReadEmptyProperiesFile(t *testing.T) {
	ignoreNotFound := false
	filename := "../test/files/empty.properties"
	var reader Reader = &FileReader{filename, ignoreNotFound}
	prop, err := reader.Read()
	assert.NotNil(t, prop)
	assert.Nil(t, err)
	assert.True(t, len(prop) == 0)
}

func Test_ShouldIgnoreCommentsOnProperiesFile(t *testing.T) {
	ignoreNotFound := false
	filename := "../test/files/test1.properties"
	var reader Reader = &FileReader{filename, ignoreNotFound}
	prop, err := reader.Read()
	assert.NotNil(t, prop)
	assert.Nil(t, err)

	_, found := prop["#this is a comment"]
	assert.False(t, found)

	_, found = prop["# this is a comment too"]
	assert.False(t, found)

	_, found = prop["### this is a comment too"]
	assert.False(t, found)

	value, found := prop["this.is.not.a.comment"]
	assert.True(t, found)
	assert.Equal(t, "#", value.Value)

	value, found = prop["this.is.#.not.a.comment"]
	assert.True(t, found)
	assert.Equal(t, "value##", value.Value)
}
