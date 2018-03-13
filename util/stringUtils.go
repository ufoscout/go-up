package util

import "strings"

/*
HasTokens returns true if the input contains at least one token
delimited by startDelimiter and endDelimiter
*/
func HasTokens(input string, startDelimiter string, endDelimiter string) bool {
	start := strings.Index(input, startDelimiter)
	end := strings.LastIndex(input, endDelimiter)
	return start >= 0 && end >= start
}

/*
FirstToken returns the first token delimited by the startDelimiter and endDelimiter.
 Example:
 startDelimiter = "${"
 endDelimiter = "}"

 - input = "abdc" -> empty {@link Optional}
 - input = "${abcd}" -> "abcd"
 - input = "${${${abcd}}}" -> "abcd"
 - input = "aaa${abcd}aaa${efgh}" -> "abcd"
*/
func FirstToken(input string, startDelimiter string, endDelimiter string) *string {
	if HasTokens(input, startDelimiter, endDelimiter) {
		for strings.Contains(input, startDelimiter) {
			start := strings.Index(input, startDelimiter)

			input = input[start+len(startDelimiter):]

			for strings.Contains(input, endDelimiter) {
				input = input[:strings.Index(input, endDelimiter)]
			}
		}
		return &input
	}
	return nil
}

/*
 AllTokens returns all tokens delimited by the startDelimiter and endDelimiter.
 Example:
 startDelimiter = "${"
 endDelimiter = "}"
  - input = "abcd" -> {}
  - input = "${abcd}" -> {"abcd"}
  - input = "${${${abcd}}}" -> {"abcd"}
  - input = "aaa${abcd}aaa${efgh}" -> {"abcd","efgh"}
*/
func AllTokens(input string, startDelimiter string, endDelimiter string) []string {
	tokens := []string{}

	token := FirstToken(input, startDelimiter, endDelimiter)
	for token != nil {
		tokenValue := *token
		tokens = append(tokens, tokenValue)
		index := strings.Index(input, tokenValue) + len(tokenValue) + len(endDelimiter)
		input = input[index:]
		token = FirstToken(input, startDelimiter, endDelimiter)
	}

	return tokens
}

/*
 AllTokensDistinct returne all tokens delimited by the startDelimiter and endDelimiter.
 If distinct is true, it removes duplicated tokens.
 Example:
 startDelimiter = "${"
 endDelimiter = "}"
 distinc = true
  - input = "abcd" -> {}
  - input = "${abcd}" -> {"abcd"}
  - input = "${${${abcd}${abcd}}}" -> {"abcd"}
  - input = "__${abcd}__${efgh}__${abcd}" -> {"abcd","efgh"}
*/
func AllTokensDistinct(input string, startDelimiter string, endDelimiter string, distinct bool) []string {
	tokens := AllTokens(input, startDelimiter, endDelimiter)
	if distinct {
		return unique(tokens)
	}
	return tokens
}

func unique(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
