package main

import (
	"fmt"
	"os"

	"../fileutils"
)

func lineshift(text []byte) {
	carriageReturn := 0
	linefeed := 0
	for i := 0; i < len(text); i++ {
		if text[i] == ("\x0A")[0] {
			linefeed++
		} else if text[i] == ("\x0D")[0] {
			carriageReturn++
		}

	}

	fmt.Println("Carriage returns: ", carriageReturn)
	fmt.Println("Linefeeds: ", linefeed)
}

func main() {
	arg1 := os.Args[1]
	fmt.Println(arg1)

	text := fileutils.FileToByteslice(arg1)

	//fmt.Printf("%q \n\n", text)

	lineshift(text)
}
