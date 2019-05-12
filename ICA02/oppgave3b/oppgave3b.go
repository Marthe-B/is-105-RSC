package main

import (
	"fmt"
	"strings"

	//"strings"

	"golang.org/x/text/encoding/charmap"

	"github.com/Marthe-B/is-105-RSC/ICA02/fileutils"
)

func main() {
	//byteslice := fileutils.FileToByteslice("../files/lang01.wl")

	//byteEdit1 := bytes.Replace(byteslice, []byte("\xF8"), []byte("\xC3\xB8"), -1)
	//byteEdit2 := bytes.Replace(byteEdit1, []byte("\xE5"), []byte("\xC3\xA5"), -1)
	//byteEdit3 := bytes.Replace(byteEdit2, []byte("\xE6"), []byte("\xC3\xA6"), -1)
	//fmt.Printf("%s", bytes.Replace(byteEdit, []byte("\xE5"), []byte("\xC3\xA5"), -1))

	//fmt.Printf("%+q", byteslice)

	//fmt.Print(fileutils.FileToByteslice("../files/lang01.wl"))
	//fmt.Printf("%s", fileutils.FileToByteslice("../files/lang03.wl"))

	b := fileutils.FileToByteslice("../files/lang01.wl")

	s := []string{}

	for _, x := range b {
		//rn := charmap.ISO8859_1.DecodeByte(x)
		rn := charmap.KOI8R.DecodeByte(x)
		//fmt.Printf("%c", rn)
		str := string(rn)  // convert rune to string
		s = append(s, str) // add x to the current string s

		//str := string(x)   // convert rune to string
		//s = append(s, str) // add x to the current string s
	}

	fmt.Printf("%s", (strings.Join(s, ""))) //returns the full string

}

func compareBytes() {
	byte01 := []byte("\xEF\xDA\xA3\xD2\xD3\xCB\x0A\xEF\xDA\xA3\xD2\xD3\xCB\xC1\x0A\xEF")
	byte02 := []byte("\xFE\xFD\x73\x6B\x61\x72\x0A\xFE\xFD\x73\x6B\x61\x72\x61\x6E\x61")
	byte03 := []byte("\xF8\x79\x65\x73\x70\x65\x73\x69\x61\x6C\x69\x73\x74\x65\x6E\x0A")

	fmt.Printf("%c\n", byte01)
	fmt.Printf("%c\n", byte02)
	fmt.Printf("%c\n", byte03)
}
