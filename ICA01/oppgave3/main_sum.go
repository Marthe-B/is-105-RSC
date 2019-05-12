package main

import (
	"fmt"
	"os"
	"strconv"

	"../sum"
)

func main() {

	// user types two numbers seperated with a space.
	// They are stored in two variables
	arg1 := os.Args[1]
	arg2 := os.Args[2]

	// was Atoi
	int1, err := strconv.ParseFloat(arg1, 64)
	if err != nil {
		fmt.Println("ERROR!", arg1, " is not a number")
		fmt.Println(err)
	}

	int2, err := strconv.ParseFloat(arg2, 64)
	if err != nil {
		fmt.Println("ERROR!", arg2, " is not a number")
		fmt.Println(err)
	}

	sumInt8 := sum.SumInt8(int8(int1), int8(int2))
	sumUint32 := sum.SumUint32(uint32(int1), uint32(int2))
	sumInt32 := sum.SumInt32(int32(int1), int32(int2))
	sumFloat64 := sum.SumFloat64(float64(int1), float64(int2))
	sumInt64 := sum.SumInt64(int64(int1), int64(int2))

	fmt.Printf("%T : %d \n", sumInt8, sumInt8)
	fmt.Printf("%T : %d \n", sumUint32, sumUint32)
	fmt.Printf("%T : %d \n", sumInt32, sumInt32)
	fmt.Printf("%T : %f \n", sumFloat64, sumFloat64)
	fmt.Printf("%T : %d \n", sumInt64, sumInt64)

	/*fmt.Println(sum.SumInt8(int8(int1), int8(int2)))
	fmt.Println(sum.SumUint32(uint32(int1), uint32(int2)))
	fmt.Println(sum.SumInt32(int32(int1), int32(int2)))
	fmt.Println(sum.SumFloat64(float64(int1), float64(int2)))
	fmt.Println(sum.SumInt64(int64(int1), int64(int2)))

	fmt.Printf("%T \n", sum.SumInt8(int8(int1), int8(int2)))
	fmt.Printf("%T \n", sum.SumUint32(uint32(int1), uint32(int2)))
	fmt.Printf("%T \n", sum.SumInt32(int32(int1), int32(int2)))
	fmt.Printf("%T \n", sum.SumFloat64(float64(int1), float64(int2)))
	fmt.Printf("%T \n", sum.SumInt64(int64(int1), int64(int2)))
	*/
}
