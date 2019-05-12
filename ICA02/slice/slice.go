package slice

//import "fmt"

// AllocateVar har INN-argument b
// b - antall bytes brukeren ønsker å allokere
// Returnerer en slice av type []byte
//
func AllocateVar(b int) []byte {
	// Kode for Oppgave 5a
	var slice = make([]byte, b)

	return slice

}

// AllocateMake tar lengde og kapasitet som b og lager en ny slice
func AllocateMake(b int) []byte {
	// Kode for Oppgave 5a
	slice := make([]byte, b)

	return slice
}

// Reslice takes a slice and reslices it
func Reslice(slc []byte, lidx int, uidx int) []byte {
	// kode her for 5b
	newSlice := slc[lidx:uidx]

	return newSlice
}

// CopySlice kopierer en slice og lager en ny
func CopySlice(slc []byte) []byte {

	//oldslice := slc
	newslice := make([]byte, len(slc))
	copy(newslice, slc)
	//fmt.Println(oldslice)
	//fmt.Println(newslice)

	return newslice
}
