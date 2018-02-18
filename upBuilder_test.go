package go_up

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"github.com/ufoscout/go-up/reader"
	"time"
)

func Test_environmentVariablesKeysShouldBeIncludedAndNormalized(t *testing.T) {

	os.Setenv("TEST_FULL_1", "100")
	os.Setenv("test.FUll_2", "200")

	up, _ := NewGoUp().
		AddReader(NewEnvReader("", true, true)).
		AddReader(NewEnvReader("", false, false)).
		build();
	assert.Equal(t, "100", up.GetString("TEST_FULL_1"))
	assert.Equal(t, "100", up.GetString("test.full.1"))

	assert.Equal(t, "200", up.GetString("test.FUll_2"))
	assert.Equal(t, "200", up.GetString("test.full.2"))

}

func Test_envVariablesShouldHaveHigherPriorityThanCustomProperties(t *testing.T) {

	os.Setenv("TEST_FULL_1", "100")
	os.Setenv("test.FUll_2", "200")

	customValue := "300";
	customKey2 := "400";
	up, _ := NewGoUp().
		AddReader(reader.NewProgrammaticReader().Add("TEST_FULL_1", customValue).Add(customKey2, customValue)).
		AddReader(NewEnvReader("", false, false)).
		build();
	assert.Equal(t, "100", up.GetString("TEST_FULL_1"));
	assert.Equal(t, customValue, up.GetString(customKey2))
}

func Test_shouldIgnoreFileNotFound(t *testing.T) {

	key := string(time.Now().UnixNano())
	up, _ := NewGoUp().
		AddReader(reader.NewProgrammaticReader().Add(key, "value")).
		AddFile("NOT VALID PATH", true).
		build();
	assert.NotNil(t, up);
	assert.True(t, up.Exists(key))

}

func Test_shouldFailIfFileNotFound(t *testing.T) {
	up, err := NewGoUp().
		AddFile("NOT VALID PATH", false).
		build();
	assert.NotNil(t, err);
	assert.Nil(t, up);
}

func Test_shouldConsiderFileAddPriority(t *testing.T) {
	up, err := NewGoUp().
		AddFile("./test/files/test1.properties", false).
		AddFile("./test/files/resource1.properties", false).
		AddFile("./test/files/inner/resource2.properties", false).
		build();
	assert.Nil(t, err)

	// from test1.properties
	assert.Equal(t, "firstvalue", up.GetString("keyOne"));

	// from resource1.properties AND resource2.properties
	assert.Equal(t, "resource2", up.GetString("name"));
}

func Test_shouldBePossibleToSetCustomPriority(t *testing.T) {

	up, _ := NewGoUp().
		AddReaderWithPriority(reader.NewProgrammaticReader().Add("one", "high"), HIGHEST_PRIORITY).
		AddReader(reader.NewProgrammaticReader().Add("one", "default")).
		build();
	assert.Equal(t, "high", up.GetString("one"));
}

func Test_shouldReplacePlaceHolders(t *testing.T) {

	key1 := "key1";
	value1 := "value1" + string(time.Now().UnixNano())

	up, _ := NewGoUp().
		Add(key1, value1).
		Add("key2", "${${key3}}__${key1}").
		Add("key3", "key1").
		build();
	assert.NotNil(t, up);
	assert.Equal(t, value1+"__"+value1, up.GetString("key2"));
}

func Test_shouldReplaceUsingCustomDelimiters(t *testing.T) {

	startDelimiter := "((";
	endDelimiter := "))";
	up, _ := NewGoUp().
		Delimiters(startDelimiter, endDelimiter).
		Add("key1", "value1").
		Add("key2", "((((key3))))__((key1))").
		Add("key3", "key1").
		build();
	assert.NotNil(t, up);
	assert.Equal(t, "value1__value1", up.GetString("key2"));
}

func Test_shouldIgnoreNotResolvedPlaceHolders(t *testing.T) {

	up, _ := NewGoUp().
		IgnoreUnresolvablePlaceholders(true).
		Add("key2", "${${key3}}__${key1}").
		Add("key3", "key1").
		build();
	assert.NotNil(t, up);
	assert.Equal(t, "${key1}__${key1}", up.GetString("key2"));
}

func Test_shouldFailIfNotResolvedPlaceHolders(t *testing.T) {
	up, err := NewGoUp().
		Add("key2", "${${key3}}__${key1}").
		Add("key3", "key1").
		build();
	assert.NotNil(t, err);
	assert.Nil(t, up);
}
