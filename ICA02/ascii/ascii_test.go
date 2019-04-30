package ascii

import (
	"testing"
	"unicode"
)

func TestGreetingASCII(t *testing.T) {
	str := GreetingASCII()

	for _, v := range str {
		if v > unicode.MaxASCII {
			t.Errorf("Char (%q) er ikke et gyldig ascii tegn", v)
		}
	}
}
