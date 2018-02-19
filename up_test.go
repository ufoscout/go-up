package go_up

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

}

func Test_ShouldReturnOptionalBool(t *testing.T) {
	properties := NewGoUp()

	properties.Add("key.one", "false")
	properties.Add("key.two", "1000000")

	prop, _ := properties.build()

	assert.Equal(t, "false", prop.GetString("key.one"))
	assert.Equal(t, false, prop.GetBool("key.one"))
	assert.Equal(t, false, prop.GetBool("key.two"))
}

func Test_ShouldReturnDefaultBool(t *testing.T) {
	properties := NewGoUp()

	properties.Add("key.one", "true")

	prop, _ := properties.build()

	assert.Equal(t, true, prop.GetBoolOrDefault("key.one", false))
	assert.Equal(t, true, prop.GetBoolOrDefault("key.two", true))

}

func Test_ShouldReturnZeroValueWhenParsingInvalidBool(t *testing.T) {
	properties := NewGoUp()

	properties.Add("key.one", "not a number")

	prop, _ := properties.build()

	var def bool
	assert.Equal(t, def, prop.GetBool("key.one"))
}

func Test_ShouldReturnDefaultWhenParsingInvalidBool(t *testing.T) {
	properties := NewGoUp()

	properties.Add("key.one", "not a number")

	prop, _ := properties.build()

	assert.Equal(t, true, prop.GetBoolOrDefault("key.one", true))

}

func Test_ShouldReturnEmptyOptionalString(t *testing.T) {
	properties := NewGoUp()

	properties.Add("key.one", "value.one")
	properties.Add("key.two", "value.two")

	prop, _ := properties.build()

	assert.False(t, prop.Exists("key.three"))
}

func Test_ShouldReturnOptionalString(t *testing.T) {
	properties := NewGoUp()

	properties.Add("key.one", "value.one")
	properties.Add("key.two", "value.two")

	prop, _ := properties.build()

	assert.Equal(t, "value.one", prop.GetString("key.one"))
	assert.Equal(t, "value.two", prop.GetString("key.two"))
}

func Test_ShouldReturnDefaultString(t *testing.T) {
	properties := NewGoUp()

	properties.Add("key.one", "value.one")

	prop, _ := properties.build()

	assert.Equal(t, "value.one", prop.GetStringOrDefault("key.one", "default"))
	assert.Equal(t, "default", prop.GetStringOrDefault("key.two", "default"))
}

func Test_ShouldReturnEmptyOptionalInteger(t *testing.T) {
	properties := NewGoUp()

	properties.Add("key.one", "value.one")
	properties.Add("key.two", "value.two")

	prop, _ := properties.build()

	assert.False(t, prop.Exists("key.three"))
}

func Test_ShouldReturnOptionalInteger(t *testing.T) {
	properties := NewGoUp()

	properties.Add("key.one", "123")
	properties.Add("key.two", "1000000")

	prop, _ := properties.build()

	assert.Equal(t, "123", prop.GetString("key.one"))
	assert.Equal(t, 123, prop.GetInt("key.one"))
	assert.Equal(t, 1000000, prop.GetInt("key.two"))
}

func Test_ShouldReturnDefaultInteger(t *testing.T) {
	properties := NewGoUp()

	properties.Add("key.one", "1")

	prop, _ := properties.build()

	assert.Equal(t, 1, prop.GetIntOrDefault("key.one", 10))
	assert.Equal(t, 10, prop.GetIntOrDefault("key.two", 10))

}

func Test_ShouldReturnZeroValueWhenParsingInvalidInt(t *testing.T) {
	properties := NewGoUp()

	properties.Add("key.one", "not a number")

	prop, _ := properties.build()

	assert.Equal(t, 0, prop.GetInt("key.one"))
}

func Test_ShouldReturnDefaultWhenParsingInvalidInt(t *testing.T) {
	properties := NewGoUp()

	properties.Add("key.one", "not a number")

	prop, _ := properties.build()

	assert.Equal(t, 10, prop.GetIntOrDefault("key.one", 10))

}

func Test_ShouldReturnOptionalFloat(t *testing.T) {
	properties := NewGoUp()

	properties.Add("key.one", "123")
	properties.Add("key.two", "1000000")

	prop, _ := properties.build()

	assert.Equal(t, float64(123), prop.GetFloat64("key.one"))
	assert.Equal(t, float64(1000000), prop.GetFloat64("key.two"))
}

func Test_ShouldReturnDefaultfloat(t *testing.T) {
	properties := NewGoUp()

	properties.Add("key.one", "1")

	prop, _ := properties.build()

	assert.Equal(t, float64(1), prop.GetFloat64OrDefault("key.one", float64(10)))
	assert.Equal(t, float64(10), prop.GetFloat64OrDefault("key.two", float64(10)))

}

func Test_ShouldReturnDefaultWhenParsingWrongFloat(t *testing.T) {
	properties := NewGoUp()

	properties.Add("key.one", "not a number")

	prop, _ := properties.build()

	assert.Equal(t, float64(10), prop.GetFloat64OrDefault("key.one", float64(10)))

}


func Test_ShouldReturnValueToArray(t *testing.T) {
	properties := NewGoUp()
	properties.Add("key.one", "111,AAAAA,BBB");

	prop,_ := properties.build()

	values := prop.GetStringSlice("key.one", ",");
	assert.Equal(t, 3, len(values));
	assert.Equal(t, "111", values[0]);
	assert.Equal(t, "AAAAA", values[1]);
	assert.Equal(t, "BBB", values[2]);
	assert.Equal(t, 0, len(prop.GetStringSlice("key.three", ",")));
}
