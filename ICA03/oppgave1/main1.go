package main

import (
	"fmt"
	"../fileutils"
)

func main() {
	//Ingen parametere (utenom filepathen til den eksekverbare filen som alltid er første parameter), oppgave 1a)
	var text1 []byte = fileutils.FileToByteslice("../files/text1.txt")
	var text2 []byte = fileutils.FileToByteslice("../files/text2.txt")

	//Print ut antall bytes i hver tekst:
	fmt.Println("\ntext1 length: ", len(text1))
	fmt.Println("text2 length: ", len(text2))

	fmt.Printf("text 1: %q \n" , text1)
	fmt.Printf("text 2: %q \n" , text2)

	exit := false
	for i := 0; i < len(text1) && i < len(text2) && !exit; i++ {
		if (text1[i] != text2[i]) {

			fmt.Printf("Diff, text1: %U (%q) , text2: %U (%q)", text1[i],text1[i], text2[i],text2[i])
			exit = true
		} else {
			fmt.Printf("%q - ", text1[i])
			fmt.Printf("%q \n", text2[i])
		}
	}
}