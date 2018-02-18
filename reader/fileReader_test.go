package reader

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldThrowExceptionWhenFileNotFound(t *testing.T) {
	ignoreNotFound := false
	filename := "./some_name"
	var reader Reader = &FileReader{filename, ignoreNotFound}
	prop, err := reader.Read();
	assert.Nil(t, prop)
	assert.NotNil(t, err)
}

func Test_ShouldIgnoreFileNotFoundIfRequested(t *testing.T) {
	ignoreNotFound := true;
	filename := "./some_name"
	var reader Reader = &FileReader{filename, ignoreNotFound}

	prop, err := reader.Read();
	assert.NotNil(t, prop)
	assert.Nil(t, err)

}

func Test_ShouldReadProperiesFromFile(t *testing.T) {
	ignoreNotFound := false;
	filename := "../test/files/test1.properties";
	var reader Reader = &FileReader{filename, ignoreNotFound}
	prop, err := reader.Read();
	assert.NotNil(t, prop)
	assert.Nil(t, err)

	_, found := prop["keyOne"]
	assert.True(t, found);
	assert.Equal(t, "firstvalue", prop["keyOne"].Value);
	assert.Equal(t, "second VALUE", prop["keyTwo"].Value);
}

func Test_ShouldReadEmptyProperiesFile(t *testing.T) {
	ignoreNotFound := false;
	filename := "../test/files/empty.properties";
	var reader Reader = &FileReader{filename, ignoreNotFound}
	prop, err := reader.Read();
	assert.NotNil(t, prop)
	assert.Nil(t, err)
	assert.True(t, len(prop) == 0);
}
