package word

import (
	"unicode"
)

// IsPalindrome check if the word is palindrome
// exported function IsPalindrome should have comment or be unexported
func IsPalindrome(s string) bool {
	var letters []rune

	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}

	for i := 0; i < len(letters)/2; i++ {
		if letters[i] != letters[len(letters)-i-1] {
			return false
		}
	}
	return true
}
