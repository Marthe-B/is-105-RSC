package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slice skrevet med %s")
	fmt.Printf("%s\n\n", byteSlice())

	fmt.Println("Slice skrevet med %q")
	fmt.Printf("%q\n\n", byteSlice())

	fmt.Println("Slice skrevet med %+q")
	fmt.Printf("%+q\n\n", byteSlice())

	fmt.Println("Slice skrevet med %c")
	fmt.Printf("%c\n\n", byteSlice())

	fmt.Println("Slice skrevet med % X")
	fmt.Printf("% X\n\n", byteSlice())

	fmt.Println("Den omdefinerte slicen skrevet ut med %s")
	fmt.Printf("%s", newByteSlice())

}

func byteSlice() []byte {
	slice := []byte("\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98")
	return slice
}

/*func newByteSlice() []byte {
	new := []byte("\xc2\xbd\x20\xc2\xb2\x20\x3d\x20\xc2\xbc\x20\xc3\xa2\x20")
	return new
}
*/

func newByteSlice() []byte {
	new := []byte("\xc2\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98") //C2 added in fron of BD to produce 1/2 sign
	return new
}
