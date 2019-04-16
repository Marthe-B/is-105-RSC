package main

import (
	"fmt"

	"../ascii"
)

func main() {
	ascii.IterateOverASCIIStringLiteral(ascii.GetASCIIStringLiteral())
	ascii.IterateOverASCIIStringLiteral("Dette er en test")

	fmt.Println("1c)")
	ascii.GreetingASCII()
}
