/***************************************

$ go test
$ go test -V	(each test name and runtime)
$ go test -v -run="French|Canal" 	(only test French and Canal)

****************************************/

package word

import (
	"testing"
)

func TestPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"asdfdsa", true},
		{"qwertytrewq", true},
		{"", true},
		{"a", true},
		{"aa", true},
		{"été", true},
		{"A man, a plan, a canal: Panama", true},
		{"palindrome", false},
		{"false", false}, // DON'T forget ","!!
	}

	for _, test := range tests {
		if x := IsPalindrome(test.input); x != test.want {
			t.Errorf("IsPalndrome(%q) = %v", test.input, x)
		}
	}
	if !IsPalindrome("asdfdsa") {
		t.Error(`IsPalndrome("asdfdsa") = false`)
	}

	if !IsPalindrome("qwertytrewq") {
		t.Error(`IsPalindrome("qwertytrewq") = false`)
	}
}

func TestNonPalindrome(t *testing.T) {
	if IsPalindrome("palindrome") {
		t.Error(`IsPalindrome("palindrome") = true`)
	}
}

func TestFrenchPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`Ispalindrome("été") = false`)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Error(`IsPalindrome("A man, a plan, a canal: Panama") = true`)
	}
}
