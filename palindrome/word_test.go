/***************************************

$ go test
$ go test -V	(each test name and runtime)
$ go test -v -run="French|Canal" 	(only test French and Canal)

****************************************/

package word

import (
	"math/rand"
	"testing"
	"time"
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

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // random len(max: 24)
	runes := make([]rune, 25)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // random rune(max: \u-0999)
		runes[i] = r
		runes[n-i-1] = r
	}
	return string(runes)
}

func randomNonPalindrome(rng *rand.Rand) string {
	n := rand.Intn(25) + 2
	runes := make([]rune, 27)
	for i := 0; i < (n+1)/2; i++ {
		tmp := 0x62 + rng.Intn(0x79-0x62)
		runes[i] = rune(tmp)
		runes[n-i-1] = rune(tmp + 1)
	}
	return string(runes)
}

func TestRandomPalindrome(t *testing.T) {
	for i := 0; i < 1000; i++ {
		seed := time.Now().UTC().UTC().UnixNano()
		rng := rand.New(rand.NewSource(seed))
		// test palindrome
		s := randomPalindrome(rng)
		if !IsPalindrome(s) {
			t.Logf("Random seed: %d", seed)
			t.Errorf("IsPalindrome(%q) = false", s)
		}

		// test non-palindrome
		ss := randomNonPalindrome(rng)
		//t.Logf(ss)
		if IsPalindrome(ss) {
			t.Errorf("IsPalindrome(\"%s\") = true", ss)
		}
	}
}
