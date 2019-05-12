package main

import (
	"fmt"
	"os"

	"../fileutils"
)

func lineshift(text []byte) {
	carriageReturn := 0 // \r
	linefeed := 0       // \n
	crlf := 0           // \r\n
	for i := 0; i < len(text); i++ {
		if text[i] == ("\x0A")[0] { // = \n
			linefeed++
		} else if text[i] == ("\x0D")[0] {
			if len(text) > i && text[i+1] == ("\x0A")[0] { // = \r\n
				crlf++
				i++
			} else { // = \r
				carriageReturn++
			}

		}

	}

	fmt.Println("Antall brukte newline tegn funnet i filen:")
	fmt.Println("CR (Commodore, Apple II family, classic Mac OS): ", carriageReturn)
	fmt.Println("LF (Unix, Linux, macOS, FreeBSD ...): ", linefeed)
	fmt.Println("CR LF (MS Windows, DOS, Symbian OS ...): ", crlf)
}

func main() {
	arg1 := os.Args[1]
	fmt.Println(arg1)

	text := fileutils.FileToByteslice(arg1)

	fmt.Printf("%q \n\n", text)

	lineshift(text)
}
