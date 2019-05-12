package unicode

import (
	"fmt"
)

// Kode for Oppgave 4a
func Translate(expression string, language string) string {

	/*var translate string
		if expression == "\x6e\x6f\x72\x64\x20\x6f\x67\x20\x73\xc3\xb8\x72" {
			if language == "jp" {
				translate = "\xe2\x80\x9c\xe5\x8c\x97\xe3\x81\xa8\xe5\x8d\x97\xe2\x80\x9d"
			} else {
				fmt.Println("unreconized language")
			}
		} else {
			fmt.Println("unregognized expression")
		}

		fmt.Println(expression)
		fmt.Println(translate)
	  return translate
	*/
	japan := "\xe2\x80\x9c\xe5\x8c\x97\xe3\x81\xa8\xe5\x8d\x97\xe2\x80\x9d"
	iceland := "\xe2\x80\x9c\x6e\x6f\x72\xc3\xb0\x75\x72\x20\x6f\x67\x20" +
		"\x73\x75\xc3\xb0\x75\x72\xe2\x80\x9d"
	var translate string

	if expression == "\x6e\x6f\x72\x64\x20\x6f\x67\x20\x73\xc3\xb8\x72" {
		switch {
		case language == "jp":
			fmt.Println("language = jp")
			//fmt.Println(japan)
			translate = japan
		case language == "is":
			fmt.Println("language = is")
			//fmt.Println(iceland)
			translate = iceland
		default:
			fmt.Println("Unrecognized language")
		}
	} else {
		fmt.Println("Unrecognized expression")
	}

	return translate

}

// “norður og suður” : e2809c6e6f72c3b07572206f67207375c3b07572e2809d
// “北と南” e2809ce58c97e381a8e58d97e2809d
// nord og sør = 6e6f7264206f672073c3b872 //\x6e\x6f\x72\x64\x20\x6f\x67\x20\x73\xc3\xb8\x72

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
