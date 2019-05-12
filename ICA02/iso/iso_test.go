package iso

import (
	"testing"
	//"unicode"

	"golang.org/x/text/encoding/charmap"
)

func TestGreetingExtendedASCII(t *testing.T) {
	//str := GreetingExtendedASCII()
	str := Greeting()

	for _, v := range str {
		_, ok := charmap.ISO8859_15.EncodeRune(v)

		if !ok {
			t.Errorf("Char (%q) er ikke et gyldig ascii extended tegn", v)
		}
		//if v > 255 {
		//	t.Errorf("Char (%q) er ikke et gyldig ascii tegn", v)
		//}
	}
}
