package treasure

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// PrintTreasureUTF8 tar en "string literal" som INN-data
// og skriver ut en korrekt text på med UTF-8 koding
// Koden er for Oppgave 3c
// Bruk strengen fra filen treasure.txt som in-data for denne funksjonen
func PrintTreasureUTF8(treasureString string) {
	fmt.Printf("%s", treasureString)

}

// GetTreasureString gets the treasure string
func GetTreasureString() string {
	b, err := ioutil.ReadFile("../treasure/treasure.txt")
	if err != nil {
		fmt.Print(err)
	}

	str := string(b)

	//r := strings.NewReplacer(" ", "", "\n", "", "\r", "")
	r := strings.NewReplacer("\\x", "")
	mods := r.Replace(string(str))
	//mods = strings.TrimSuffix(mods, "\n")

	decoded0, err := hex.DecodeString(mods)
	if err != nil {
		log.Fatal(err)
	}

	// hex to replace F8-C3B8, E6-C3A6, E5-C3 A5
	decoded1 := bytes.Replace(decoded0, []byte("\xF8"), []byte("\xC3\xB8"), -1)
	decoded2 := bytes.Replace(decoded1, []byte("\xE6"), []byte("\xC3\xA6"), -1)
	decoded3 := bytes.Replace(decoded2, []byte("\xE5"), []byte("\xC3\xA5"), -1)

	//dec := string(decoded0)

	return string(decoded3)

}
