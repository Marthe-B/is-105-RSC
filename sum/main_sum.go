package main

import (
	"fmt"
	"os"
	"strconv"

	"./sum"
)

func main() {

	// `os.Args` provides access to raw command-line
	// arguments. Note that the first value in this slice
	// is the path to the program, and `os.Args[1:]`
	// holds the arguments to the program.
	//argsWithProg := os.Args
	//argsWithoutProg := os.Args[1:]

	// user types two numbers seperated with a space. They are stored in two variables
	arg1 := os.Args[1]
	arg2 := os.Args[2]

	int1, err := strconv.Atoi(arg1)
	if err == nil {
		//fmt.Println(int1)
	}

	int2, err := strconv.Atoi(arg2)
	if err == nil {
		//fmt.Println(int2)
	}

	//fmt.Println(argsWithProg)
	//fmt.Println(argsWithoutProg)

	//fmt.Println(arg1)
	//fmt.Println(arg2)

	//fmt.Println(int1)
	//fmt.Println(int2)

	//fmt.Println(int1 + int2)

	fmt.Println(sum.SumInt64(int64(int1), int64(int2)))
	fmt.Println(sum.SumFloat64(float64(int1), float64(int2)))
}

//func SumInt8(a, b int8) int8 {
//	return a + b
//}

//func SumFloat64(a, b float64) float64 {
//	return a + b
//}
