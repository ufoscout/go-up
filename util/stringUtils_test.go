package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var startDelimiter string = "${"
var endDelimiter string = "}"

func TestStringsWithTokens(t *testing.T) {

	assert.True(t, HasTokens("hello ${hi}", startDelimiter, endDelimiter))
	assert.True(t, HasTokens("hello ${}", startDelimiter, endDelimiter))
	assert.True(t, HasTokens("${hello}", startDelimiter, endDelimiter))
	assert.True(t, HasTokens("${hello}}", startDelimiter, endDelimiter))
	assert.True(t, HasTokens("${${hello}", startDelimiter, endDelimiter))
	assert.True(t, HasTokens("${${hello}}", startDelimiter, endDelimiter))
	assert.True(t, HasTokens("start ${hello}", startDelimiter, endDelimiter))
	assert.True(t, HasTokens("${hello} end", startDelimiter, endDelimiter))

	assert.False(t, HasTokens("${hello", startDelimiter, endDelimiter))
	assert.False(t, HasTokens("$hello}", startDelimiter, endDelimiter))
	assert.False(t, HasTokens("{hello}", startDelimiter, endDelimiter))
	assert.False(t, HasTokens("}hello ${", startDelimiter, endDelimiter))
}

func TestShouldReturnTheFirstToken(t *testing.T) {

	assert.Equal(t, "hello", *FirstToken("${hello}", startDelimiter, endDelimiter))
	assert.Equal(t, "hello1", *FirstToken("${hello1}, ${hello2}", startDelimiter, endDelimiter))
	assert.Equal(t, "hello2", *FirstToken("hello1 ${hello2}", startDelimiter, endDelimiter))

	assert.Equal(t, "abcd", *FirstToken("${${${abcd}}}", startDelimiter, endDelimiter))
	assert.Equal(t, "abcd", *FirstToken("aaa${abcd}aaa${efgh}", startDelimiter, endDelimiter))

	assert.Nil(t, FirstToken("hello", startDelimiter, endDelimiter))
}

func Test_ShouldReturnEmptyArray(t *testing.T) {
	assert.Equal(t, 0, len(AllTokens("", "{", "}")))
}

func Test_ShouldReturnEmpty2(t *testing.T) {
	input := "asfasdfasfdasfdasf_${_asdasd"
	assert.True(t, len(AllTokens(input, startDelimiter, endDelimiter)) == 0)
}

func Test_ShouldReturnEmpty3(t *testing.T) {
	input := "asfasdfasfdasfdasf_}_asdasd"
	assert.True(t, len(AllTokens(input, startDelimiter, endDelimiter)) == 0)
}

func Test_ShouldReturnSimpleToken(t *testing.T) {
	input := "${TOKEN}"
	tokens := AllTokens(input, startDelimiter, endDelimiter)

	assert.Equal(t, 1, len(tokens))
	assert.Equal(t, "TOKEN", tokens[0])
}

func Test_ShouldReturnTokenAtTheEnd(t *testing.T) {
	input := "asfasdfasfdasfdasf_${TOKEN}"
	tokens := AllTokens(input, startDelimiter, endDelimiter)

	assert.Equal(t, 1, len(tokens))
	assert.Equal(t, "TOKEN", tokens[0])
}

func Test_ShouldReturnTokenAtTheBeginning(t *testing.T) {
	input := "${TOKEN}asfasdfasfd_${_asfdasf"
	tokens := AllTokens(input, startDelimiter, endDelimiter)

	assert.Equal(t, 1, len(tokens))
	assert.Equal(t, "TOKEN", tokens[0])
}

func Test_ShouldReturnTokenInTheMiddle(t *testing.T) {
	input := "asfasdfasf${TOKEN}dasfdasf"
	tokens := AllTokens(input, startDelimiter, endDelimiter)

	assert.Equal(t, 1, len(tokens))
	assert.Equal(t, "TOKEN", tokens[0])
}

func Test_ShouldReturnTokenInTheMiddle2(t *testing.T) {
	input := "asfas}dfasf${TOKEN}dasfdasf"
	tokens := AllTokens(input, startDelimiter, endDelimiter)

	assert.Equal(t, 1, len(tokens))
	assert.Equal(t, "TOKEN", tokens[0])
}

func Test_ShouldReturnTokenInTheMiddle3(t *testing.T) {
	input := "asfas${dfasf${TOKEN}dasfdasf"
	tokens := AllTokens(input, startDelimiter, endDelimiter)

	assert.Equal(t, 1, len(tokens))
	assert.Equal(t, "TOKEN", tokens[0])
}

func Test_ShouldReturnTokenInTheMiddle4(t *testing.T) {
	input := "asfas${dfasf${TOKEN}dasfd${asf"
	tokens := AllTokens(input, startDelimiter, endDelimiter)

	assert.Equal(t, 1, len(tokens))
	assert.Equal(t, "TOKEN", tokens[0])
}

func Test_ShouldReturnTokenInTheMiddle5(t *testing.T) {
	input := "asfas${dfasf${TOKEN}dasfd}asf"
	tokens := AllTokens(input, startDelimiter, endDelimiter)

	assert.Equal(t, 1, len(tokens))
	assert.Equal(t, "TOKEN", tokens[0])
}

func Test_ShouldReturnInnerToken(t *testing.T) {
	input := "asfas${dfasf${TOKEN}dasf}dasf"
	tokens := AllTokens(input, startDelimiter, endDelimiter)

	assert.Equal(t, 1, len(tokens))
	assert.Equal(t, "TOKEN", tokens[0])
}

func Test_ShouldReturnNestedToken(t *testing.T) {
	input := "${${${TOKEN}}}"
	tokens := AllTokens(input, startDelimiter, endDelimiter)

	assert.Equal(t, 1, len(tokens))
	assert.Equal(t, "TOKEN", tokens[0])
}

func Test_ShouldReturnAllTokens1(t *testing.T) {
	input := "___${____${TOKEN_1}__}___${TOKEN_2}__"
	tokens := AllTokens(input, startDelimiter, endDelimiter)

	assert.Equal(t, 2, len(tokens))
	assert.Equal(t, "TOKEN_1", tokens[0])
	assert.Equal(t, "TOKEN_2", tokens[1])
}

func Test_ShouldReturnAllTokens2(t *testing.T) {
	input := "___${____${TOKEN_1__}___${TOKEN_2}__"
	tokens := AllTokens(input, startDelimiter, endDelimiter)

	assert.Equal(t, 2, len(tokens))
	assert.Equal(t, "TOKEN_1__", tokens[0])
	assert.Equal(t, "TOKEN_2", tokens[1])
}

func Test_ShouldReturnAllTokens3(t *testing.T) {
	input := "___${____${TOKEN_2}_}___${TOKEN_2}__"
	tokens := AllTokens(input, startDelimiter, endDelimiter)

	assert.Equal(t, 2, len(tokens))
	assert.Equal(t, "TOKEN_2", tokens[0])
	assert.Equal(t, "TOKEN_2", tokens[1])
}

func Test_ShouldReturnAllTokens4(t *testing.T) {
	input := "___${____${TOKEN_1}_${TOKEN_2}_}___${TOKEN_2}_${${TOKEN_3}}_"
	tokens := AllTokensDistinct(input, startDelimiter, endDelimiter, false)

	assert.Equal(t, 4, len(tokens))
	assert.Equal(t, "TOKEN_1", tokens[0])
	assert.Equal(t, "TOKEN_2", tokens[1])
	assert.Equal(t, "TOKEN_2", tokens[2])
	assert.Equal(t, "TOKEN_3", tokens[3])
}

func Test_ShouldReturnUniqueTokens1(t *testing.T) {
	input := "___${____${TOKEN_2}_}___${TOKEN_2}__"
	tokens := AllTokensDistinct(input, startDelimiter, endDelimiter, true)

	assert.Equal(t, 1, len(tokens))
	assert.Equal(t, "TOKEN_2", tokens[0])
}

func Test_ShouldReturnUniqueTokens2(t *testing.T) {
	input := "___${____${TOKEN_1}_${TOKEN_2}_}___${TOKEN_2}_${${TOKEN_3}}_"
	tokens := AllTokensDistinct(input, startDelimiter, endDelimiter, true)

	assert.Equal(t, 3, len(tokens))
	assert.Equal(t, "TOKEN_1", tokens[0])
	assert.Equal(t, "TOKEN_2", tokens[1])
	assert.Equal(t, "TOKEN_3", tokens[2])
}
