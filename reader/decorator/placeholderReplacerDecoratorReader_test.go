package decorator

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ufoscout/go-up/reader"
)

func Test(t *testing.T) {
}

func Test_ShouldResolveSimpleKeys(t *testing.T) {
	properties := reader.NewProgrammaticReader()
	properties.Add("key.one", "${key.two}")
	properties.Add("key.two", "value.two")

	ignoreNotResolvable := false

	decorator := &PlaceholderReplacerDecoratorReader{properties, "${", "}", ":", ignoreNotResolvable}
	output, _ := decorator.Read()
	assert.NotNil(t, output)

	assert.Equal(t, 2, len(output))

	assert.Equal(t, "value.two", output["key.one"].Value)
	assert.Equal(t, "value.two", output["key.two"].Value)

}

func Test_ShouldNotResolveUnresolvableKeys(t *testing.T) {
	properties := reader.NewProgrammaticReader()
	properties.AddProperty("key.unresolvable", reader.Property{"${key.two}", false})
	properties.Add("key.one", "${key.two}")
	properties.Add("key.two", "value.two")

	ignoreNotResolvable := false
	decorator := &PlaceholderReplacerDecoratorReader{properties, "${", "}", ":", ignoreNotResolvable}
	output, _ := decorator.Read()

	assert.NotNil(t, output)

	assert.Equal(t, 3, len(output))

	assert.Equal(t, "${key.two}", output["key.unresolvable"].Value)
	assert.Equal(t, "value.two", output["key.one"].Value)
	assert.Equal(t, "value.two", output["key.two"].Value)

}

func Test_ShouldFailIfNotResolvablePlaceholders(t *testing.T) {
	properties := reader.NewProgrammaticReader()
	properties.Add("key.1", "${key.4}")
	properties.Add("key.2", "${key.1}")
	properties.Add("key.3", "${key.2}")
	properties.Add("key.4", "${key.3}")

	ignoreNotResolvable := false

	decorator := &PlaceholderReplacerDecoratorReader{properties, "${", "}", ":", ignoreNotResolvable}
	output, err := decorator.Read()

	assert.Nil(t, output)
	assert.NotNil(t, err)

	message := err.Error()

	assert.True(t, strings.Contains(message, "key: [key.1] value: [${key.4}]"))
	assert.True(t, strings.Contains(message, "key: [key.2] value: [${key.1}]"))
	assert.True(t, strings.Contains(message, "key: [key.3] value: [${key.2}]"))
	assert.True(t, strings.Contains(message, "key: [key.4] value: [${key.3}]"))

}

func Test_ShouldNotLoopOnSelfReferencingKeys(t *testing.T) {
	properties := reader.NewProgrammaticReader()
	properties.Add("key.one", "${key.one}")

	ignoreNotResolvable := true
	decorator := &PlaceholderReplacerDecoratorReader{properties, "${", "}", ":", ignoreNotResolvable}
	output, _ := decorator.Read()
	assert.NotNil(t, output)

	assert.Equal(t, 1, len(output))

	assert.Equal(t, "${key.one}", output["key.one"].Value)

}

func Test_ShouldNotResolveCircularReferencingKeys(t *testing.T) {
	properties := reader.NewProgrammaticReader()
	properties.Add("key.one", "${key.two}")
	properties.Add("key.two", "${key.one}")

	ignoreNotResolvable := true
	decorator := &PlaceholderReplacerDecoratorReader{properties, "${", "}", ":", ignoreNotResolvable}
	output, _ := decorator.Read()
	assert.NotNil(t, output)

	assert.Equal(t, 2, len(output))

	assert.Equal(t, "${key.two}", output["key.one"].Value)
	assert.Equal(t, "${key.one}", output["key.two"].Value)

}

func Test_ShouldRecursivelyResolveKeys(t *testing.T) {
	properties := reader.NewProgrammaticReader()
	properties.Add("key.1", "${key.2}")
	properties.Add("key.2", "${key.3} world!")
	properties.Add("key.3", "Hello")
	properties.Add("key.4", "${key.2}")

	ignoreNotResolvable := false
	decorator := &PlaceholderReplacerDecoratorReader{properties, "${", "}", ":", ignoreNotResolvable}
	output, _ := decorator.Read()
	assert.NotNil(t, output)

	assert.Equal(t, 4, len(output))

	assert.Equal(t, "Hello world!", output["key.1"].Value)
	assert.Equal(t, "Hello world!", output["key.2"].Value)
	assert.Equal(t, "Hello", output["key.3"].Value)
	assert.Equal(t, "Hello world!", output["key.4"].Value)

}

func Test_ShouldRecursivelyResolveDynamicKeys(t *testing.T) {
	properties := reader.NewProgrammaticReader()
	properties.Add("key.1", "${${key.2}}")
	properties.Add("key.2", "${key.3}")
	properties.Add("key.3", "key.4")
	properties.Add("key.4", "Hello world!")

	ignoreNotResolvable := false
	decorator := &PlaceholderReplacerDecoratorReader{properties, "${", "}", ":", ignoreNotResolvable}
	output, _ := decorator.Read()
	assert.NotNil(t, output)

	assert.Equal(t, 4, len(output))

	assert.Equal(t, "Hello world!", output["key.1"].Value)
	assert.Equal(t, "key.4", output["key.2"].Value)
	assert.Equal(t, "key.4", output["key.3"].Value)
	assert.Equal(t, "Hello world!", output["key.4"].Value)

}

func Test_ShouldDetectDefaultValue(t *testing.T) {
	properties := reader.NewProgrammaticReader()
	properties.Add("key.one", "${value1:defaultValue1}")
	properties.Add("key.two", "${value2:defaultValue2}")

	ignoreNotResolvable := true
	decorator := &PlaceholderReplacerDecoratorReader{properties, "${", "}", ":", ignoreNotResolvable}
	output, _ := decorator.Read()
	assert.NotNil(t, output)

	assert.Equal(t, 2, len(output))

	assert.Equal(t, "defaultValue1", output["key.one"].Value)
	assert.Equal(t, "defaultValue2", output["key.two"].Value)

}

func Test_ShouldDetectDefaultValueAndReplaceRecursively(t *testing.T) {
	properties := reader.NewProgrammaticReader()
	properties.Add("key.one", "${value1:defaultValue1}")
	properties.Add("key.two", "${key.one}")

	ignoreNotResolvable := true
	decorator := &PlaceholderReplacerDecoratorReader{properties, "${", "}", ":", ignoreNotResolvable}
	output, _ := decorator.Read()
	assert.NotNil(t, output)

	assert.Equal(t, 2, len(output))

	assert.Equal(t, "defaultValue1", output["key.one"].Value)
	assert.Equal(t, "defaultValue1", output["key.two"].Value)

}

func Test_ShouldDetectDefaultValueAndReplaceRecursively2(t *testing.T) {
	properties := reader.NewProgrammaticReader()
	properties.Add("key.one", "${value1:defaultValue1}")
	properties.Add("key.two", "${key.one:defaultValue2}")

	ignoreNotResolvable := true
	decorator := &PlaceholderReplacerDecoratorReader{properties, "${", "}", ":", ignoreNotResolvable}
	output, _ := decorator.Read()
	assert.NotNil(t, output)

	assert.Equal(t, 2, len(output))

	assert.Equal(t, "defaultValue1", output["key.one"].Value)
	assert.Equal(t, "defaultValue1", output["key.two"].Value)

}

func Test_ShouldDetectNestedDefaultValues(t *testing.T) {
	properties := reader.NewProgrammaticReader()
	properties.Add("key.one", "${${value1:${key.two}}")
	properties.Add("key.two", "${key.three}")
	properties.Add("key.three", "value")

	ignoreNotResolvable := true
	decorator := &PlaceholderReplacerDecoratorReader{properties, "${", "}", ":", ignoreNotResolvable}
	output, _ := decorator.Read()
	assert.NotNil(t, output)

	assert.Equal(t, 3, len(output))

	assert.Equal(t, "value", output["key.one"].Value)
	assert.Equal(t, "value", output["key.two"].Value)
	assert.Equal(t, "value", output["key.three"].Value)

}
