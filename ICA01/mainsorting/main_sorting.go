package main

import (
	"fmt"

	algorithms "../algorithms"
)

func main() {
	values := []int{1, 4, 2, 3, 7, 6, 8, 5, 3, 0}

	fmt.Println("List:", values)
	algorithms.Bubble_sort_modified(values)
	//algorithms.Bubble_sort(values)
	fmt.Println("List BubbleSorted:  ", values)
}
