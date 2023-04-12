package mexicanWave

import (
	"strings"
	"unicode"
)

// Wave takes a word as input and returns a Mexican Wave of this word as a slice of strings.
func Wave(input string) []string {
	lowerInput := strings.ToLower(input)
	result := []string{} // When instantiating something with the var keyword Golang will give it the default value. The default value for
	// slices is nil. So in the test we were comparing nil with an empty slice of strings that's why we got the wrong result. So we need to
	// define as above in order to avoid the drawback of nil slice declaration.

	for i, char := range lowerInput {
		examChar := string(char)
		if isLetter(examChar) {
			tmpRes := lowerInput[:i] + strings.ToUpper(examChar) + lowerInput[i+1:]
			result = append(result, tmpRes)
		}
	}
	return result
}

// isLetter helps us define if a character is a letter or not.
func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
