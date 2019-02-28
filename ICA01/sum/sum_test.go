package sum

import "testing"

// Check https://golang.org/ref/spec#Numeric_types and stress the limits!
var sum_tests_int8 = []struct {
	n1       int8
	n2       int8
	expected int8
}{
	{5, 2, 7},
	{4, 5, 9},
	{120, 1, 121},
}

func TestSumInt8(t *testing.T) {
	for _, v := range sum_tests_int8 {
		if val := SumInt8(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}
