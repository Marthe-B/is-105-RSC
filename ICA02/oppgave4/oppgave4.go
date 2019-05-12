package main

import (
	"fmt"
	"os"

	"../unicode"
)

func main() {

	expression := os.Args[1]
	language := os.Args[2]

	fmt.Println(unicode.Translate(expression, language))

	unicode.UnicodeCodePointDemo()
}
