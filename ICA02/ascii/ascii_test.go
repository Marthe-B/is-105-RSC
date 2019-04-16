package ascii

import (
	"fmt"
	"testing"
	"unicode"
)

func TestGreetingASCII(t *testing.T) {
	str := GreetingASCII()

	fmt.Printf("Testing: %s \n", str)
	for _, v := range str {
		if v > unicode.MaxASCII {
			t.Errorf("Char (%q) er ikke et gyldig ascii tegn", v)
		}
	}
}
