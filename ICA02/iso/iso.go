package iso

import (
	"fmt"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

/*
const asciiExtended = "\x80\x81\x82\x83\x84\x85\x86\x87\x88\x89\x8A\x8B\x8C" +
	"\x8D\x8E\x8F\x90\x91\x92\x93\x94\x95\x96\x97\x98\x99\x9A\x9B\x9C\x9D\x9E\x9F" +
	"\xA0\xA1\xA2\xA3\xA4\xA5\xA6\xA7\xA8\xA9\xAA\xAB\xAC\xAD\xAE\xAF\xB0\xB1\xB2" +
	"\xB3\xB4\xB5\xB6\xB7\xB8\xB9\xBA\xBB\xBC\xBD\xBE\xBF\xC0\xC1\xC2\xC3\xC4\xC5" +
	"\xC6\xC7\xC8\xC9\xCA\xCB\xCC\xCD\xCE\xCF\xD0\xD1\xD2\xD3\xD4\xD5\xD6\xD7\xD8" +
	"\xD9\xDA\xDB\xDC\xDD\xDE\xDF\xE0\xE1\xE2\xE3\xE4\xE5\xE6\xE7\xE8\xE9\xEA\xEB" +
	"\xEC\xED\xEE\xEF\xF0\xF1\xF2\xF3\xF4\xF5\xF6\xF7\xF8\xF9\xFA\xFB\xFC\xFD\xFE" +
	"\xFF"

func GetAsciiExtended() string {
	return asciiExtended
}
*/

func GetExtendedASCIIStringLiteral() string {
	bs := make([]byte, 128)
	for i := 128; i < 256; i++ {
		bs[i-128] = byte(i)
	}

	return string(bs)
}

// Oppgave 2a
// Lag selv en "string literal" med utvidet ASCII
// Blir det kun en slik "string literal" eller trenger man flere
// for å representere utvidet ASCII?

// IterateOverASCIIStringLiteral tar en "string literal" som INN-data
func IterateOverASCIIStringLiteral(sl string) {
	// Kode for Oppgave 2a
	for i := 0; i < len(sl); i++ {
		fmt.Printf("%X ", sl[i])
		fmt.Printf("%c ", sl[i])
		fmt.Printf("%b \n", sl[i])
		//fmt.Printf("%#q \n", sl[i])
	}
}

// GreetingExtendedASCII returnerer en tekst-streng i
// utvidet ASCII
// Kode for Oppgave 2c
func GreetingExtendedASCII() string {
	return "Salut, ça va °-) Ça coute €50" +
		"Salut, ça va °-) Κοστίζει €50"
}

// Greeting er midlertidig så jeg ikke fucker opp noe som allerede fungerer
func Greeting() string {

	// This byte slice contains the hexa representation of the string
	// based on the ISO 8859_15 chracatermap
	hex := []byte("\x22\x53\x61\x6c\x75\x74\x2c\x20\xc7\x61\x20\x76\x61" +
		"\x20\xb0\x2d\x29\x20\xc7\x61\x20\x63\x6f\x75\x74\x65\x20\xa4\x35\x30\x22")

	//THIS IS THE PART THAT ACTUALLY WORKS! HOLY SHIT

	s := []string{} // creates an empty string something

	for _, b := range hex {
		rn := charmap.ISO8859_15.DecodeByte(b)
		//fmt.Printf("%c", rn)
		x := string(rn)  // convert rune to string
		s = append(s, x) // add x to the current string s
	}

	return (strings.Join(s, "")) //returns the full string

}
