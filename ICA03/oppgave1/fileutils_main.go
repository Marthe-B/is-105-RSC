package main

import (
	"fmt"
	//"strings"
	"./fileutils"
)

func main() {
	text1 := fileutils.FileToByteslice("../files/text1.txt")
	for i := 0; i < len(text1); i++ {
		fmt.Printf("%U \n", text1[i])
	}

	text2 := fileutils.FileToByteslice("../files/text2.txt")
	for i := 0; i < len(text2); i++ {
		fmt.Printf("%U", text2[i])
	}

	//fmt.Println(strings.Join("text1 length: ", len(text1)))
}
