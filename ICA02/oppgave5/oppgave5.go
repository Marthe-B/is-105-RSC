package main

import (
	"fmt"

	"../slice"
)

func main() {

	fmt.Println("AllocateVar")
	fmt.Println(slice.AllocateVar(16))
	fmt.Println("AlocateMake")
	fmt.Println(slice.AllocateMake(16))
	fmt.Println("")

	byteslice1 := (slice.AllocateMake(128))

	fmt.Println("&byteslice1[0]")
	fmt.Println(&byteslice1[0])
	aslice := slice.Reslice(byteslice1, 50, 100)
	fmt.Println("&aslice[0]")
	fmt.Println(&aslice[0])
	fmt.Println("&byteslice1[50]")
	fmt.Println(&byteslice1[50])

	fmt.Println("")

	oldslice := slice.AllocateMake(16)
	newslice := slice.CopySlice(oldslice)

	fmt.Println("&oldlice[0]")
	fmt.Println(&oldslice[0])
	fmt.Println(oldslice)
	fmt.Println("&newslice[0]")
	fmt.Println(&newslice[0])
	fmt.Println(newslice)

}
