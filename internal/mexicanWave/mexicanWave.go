package mexicanWave

import (
	"strings"
	"unicode"
)

// Wave returns a Mexican Wave as a slice of strings.
func Wave(input string) []string {
	lowerInput := strings.ToLower(input)
	var result []string

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
