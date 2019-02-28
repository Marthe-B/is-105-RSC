package unicode

import (
  "fmt"
)

// Kode for Oppgave 4a
func Translate(expression string, language string) string {
	return ""
}

// Kode for Oppgave 4b
func UnicodeCodePointDemo() {
	// Hva er dette?
	// Er det likt på MS Windows og macOS?
  fmt.Println("\xf0\x9F\x98\x80")
  fmt.Println("\xf0\x9F\x98\x97")
  // Demonstrerer at deler av et tegn representert med flere bytes
  // kan ikke tolkes innenfor koden (unicode)
  fmt.Println("\xf0\x9F\x98")
  fmt.Println("\xf0\x9F")
  fmt.Println("\xf0")
}

