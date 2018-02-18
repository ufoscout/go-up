package go_up

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {

}


func Test_ShouldReturnEmptyOptionalString(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "value.one");
properties.Add("key.two", "value.two");

prop,_ := properties.build()

assert.False(t, prop.Exists("key.three"));
}


func Test_ShouldReturnOptionalString(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "value.one");
properties.Add("key.two", "value.two");

prop,_ := properties.build()

assert.Equal(t, "value.one", prop.GetString("key.one"));
assert.Equal(t, "value.two", prop.GetString("key.two"));
}


func Test_ShouldReturnDefaultString(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "value.one");

prop,_ := properties.build()

assert.Equal(t, "value.one", prop.GetStringOrDefault("key.one", "default"));
assert.Equal(t, "default", prop.GetStringOrDefault("key.two", "default"));
}


func Test_ShouldReturnEmptyOptionalInteger(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "value.one");
properties.Add("key.two", "value.two");

prop,_ := properties.build()

assert.False(t, prop.Exists("key.three"));
}



func Test_ShouldReturnOptionalInteger(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "123");
properties.Add("key.two", "1000000");

prop,_ := properties.build()

assert.Equal(t, 123, prop.GetInt("key.one"));
assert.Equal(t, 1000000, prop.GetInt("key.two"));
}


func Test_ShouldReturnDefaultInteger(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "1");

prop,_ := properties.build()

assert.Equal(t, 1, prop.GetIntOrDefault("key.one", 10));
assert.Equal(t, 10, prop.GetIntOrDefault("key.two", 10));

}

func Test_ShouldReturnZeroValueForInt(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "not a number");

prop,_ := properties.build()

	assert.Equal(t, 0, prop.GetInt("key.one"))
}

/*
func Test_ShouldReturnEmptyOptionalDouble(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "value.one");
properties.Add("key.two", "value.two");

prop,_ := properties.build()

assert.False(t, prop.getDouble("key.three").isPresent());
}



func Test_ShouldReturnDefaultDouble(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "1");

prop,_ := properties.build()

assert.Equal(t, 1d, prop.getDouble("key.one", 1.1111d), 0.1d);
assert.Equal(t, 10d, prop.getDouble("key.two", 10d), 0.1d);

}

(expected=NumberFormatException.class)
func Test_ShouldThrowExceptionParsingWrongDouble(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "not a number");

prop,_ := properties.build()

prop.getDouble("key.one");

}


func Test_ShouldReturnEmptyOptionalFloat(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "value.one");
properties.Add("key.two", "value.two");

prop,_ := properties.build()

assert.False(t, prop.getFloat("key.three").isPresent());
}



func Test_ShouldReturnOptionalFloat(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "123");
properties.Add("key.two", "1000000");

prop,_ := properties.build()

assert.Equal(t, 123f, prop.getFloat("key.one").floatValue(), 0.1f);
assert.Equal(t, 1000000f, prop.getFloat("key.two").floatValue(), 0.1f);
}


func Test_ShouldReturnDefaultfloat(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "1");

prop,_ := properties.build()

assert.Equal(t, 1f, prop.getFloat("key.one", 10), 0.1f);
assert.Equal(t, 10f, prop.getFloat("key.two", 10), 0.1f);

}

(expected=NumberFormatException.class)
func Test_ShouldThrowExceptionParsingWrongFloat(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "not a number");

prop,_ := properties.build()

prop.getFloat("key.one");

}


func Test_ShouldReturnEmptyOptionalLong(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "value.one");
properties.Add("key.two", "value.two");

prop,_ := properties.build()

assert.False(t, prop.getLong("key.three").isPresent());
}



func Test_ShouldReturnOptionalLong(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "123");
properties.Add("key.two", "1000000");

prop,_ := properties.build()

assert.Equal(t, 123l, prop.getLong("key.one").longValue());
assert.Equal(t, 1000000l, prop.getLong("key.two").longValue());
}


func Test_ShouldReturnDefaultLong(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "1");

prop,_ := properties.build()

assert.Equal(t, 1l, prop.getLong("key.one", 10l));
assert.Equal(t, 10l, prop.getLong("key.two", 10l));

}

(expected=NumberFormatException.class)
func Test_ShouldThrowExceptionParsingWrongLong(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "not a number");

prop,_ := properties.build()

prop.getLong("key.one");

}


func Test_ShouldReturnEmptyOptionalEnum(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "value.one");
properties.Add("key.two", "value.two");

prop,_ := properties.build()

assert.False(t, prop.getEnum("key.three", NeedSomebodyToLove.class).isPresent());
}



func Test_ShouldReturnOptionalEnum(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "ME");
properties.Add("key.two", "THEM");

prop,_ := properties.build()

assert.Equal(t, NeedSomebodyToLove.ME, prop.getEnum("key.one", NeedSomebodyToLove.class));
assert.Equal(t, NeedSomebodyToLove.THEM, prop.getEnum("key.two", NeedSomebodyToLove.class));
}


func Test_ShouldReturnDefaultEnum(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "ME");

prop,_ := properties.build()

assert.Equal(t, NeedSomebodyToLove.ME, prop.getEnum("key.one", NeedSomebodyToLove.THEM));
assert.Equal(t, NeedSomebodyToLove.THEM, prop.getEnum("key.two", NeedSomebodyToLove.THEM));

}

(expected=IllegalArgumentException.class)
func Test_ShouldThrowExceptionParsingWrongEnumLong(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "not an enum");

prop,_ := properties.build()

prop.getEnum("key.one", NeedSomebodyToLove.class);

}


func Test_ShouldReturntheKeyApplyingTheMapFunction(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "111");

prop,_ := properties.build()

assert.Equal(t, 111, prop.GetString("key.one", Integer::valueOf).intValue());
assert.Equal(t, 222, prop.GetString("key.two", 222, Integer::valueOf).intValue());
assert.False(t, prop.GetString("key.three", Integer::valueOf).isPresent());
}


func Test_ShouldReturnValueToArray(t *testing.T) {
final Map<String, String> properties = new HashMap<>(t *testing.T);

properties.Add("key.one", "111,AAAAA,BBB");

prop,_ := properties.build()

final String[] values = prop.getArray("key.one");
assert.Equal(t, 3, values.length);
assert.Equal(t, "111", values[0]);
assert.Equal(t, "AAAAA", values[1]);
assert.Equal(t, "BBB", values[2]);
assert.Equal(t, 0, prop.getArray("key.three").length);
}


func Test_ShouldReturnValueToList(t *testing.T) {
final Map<String, String> properties = new HashMap<>(t *testing.T);

properties.Add("key.one", "111,AAAAA,BBB");

prop,_ := properties.build()

final List<String> values = prop.getList("key.one");
assert.Equal(t, 3, values.size());
assert.Equal(t, "111", values.get(0));
assert.Equal(t, "AAAAA", values.get(1));
assert.Equal(t, "BBB", values.get(2));
assert.Equal(t, 0, prop.getList("key.three").size());
}


func Test_ShouldReturnValueToListOfCustomObjects(t *testing.T) {
properties := NewGoUp()

properties.Add("key.one", "111,222,333");

prop,_ := properties.build()

final List<Integer> values = prop.getList("key.one", Integer::valueOf);
assert.Equal(t, 3, values.size());
assert.Equal(t, 111, values.get(0).intValue());
assert.Equal(t, 222, values.get(1).intValue());
assert.Equal(t, 333, values.get(2).intValue());
assert.Equal(t, 0, prop.getList("key.three", Integer::valueOf).size());
}

*/