package main

import (
	"../ascii"
)

func main() {

	for i := 128; i < 255; i++ {

	}

	ascii.IterateOverASCIIStringLiteral(ascii.GetASCIIStringLiteral())
	ascii.IterateOverASCIIStringLiteral("Dette er en test")
}
