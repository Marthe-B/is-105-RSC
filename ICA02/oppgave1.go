package main

import (
	"github.com/Marthe-B/is105-ica02/ascii"
)

func main() {

	for i := 128; i < 255; i++ {

	}

	ascii.IterateOverASCIIStringLiteral(ascii.GetASCIIStringLiteral())
	ascii.IterateOverASCIIStringLiteral("Dette er en test")
}
