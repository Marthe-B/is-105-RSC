package sum

import "testing"

// Check https://golang.org/ref/spec#Numeric_types and stress the limits!
// Test int8
var sum_tests_int8 = []struct {
	n1       int8
	n2       int8
	expected int8
}{
	{5, 2, 7},
	{4, 5, 9},
	{300, 1, 301},
}

// Test uint32
var sum_tests_uint32 = []struct {
	n1       uint32
	n2       uint32
	expected uint32
}{
	{5, 2, 7},
	{4, 5, 9},
	{120, 1, 121},
}

// Test int32
var sum_tests_int32 = []struct {
	n1       int32
	n2       int32
	expected int32
}{
	{5, 2, 7},
	{4, 5, 9},
	{120, 1, 121},
}

// Test float64
var sum_tests_float64 = []struct {
	n1       float64
	n2       float64
	expected float64
}{
	{5.5, 5.5, 11},
	{4.2, 5.8, 10},
	{120, 1, 121},
}

// Test int64
var sum_tests_int64 = []struct {
	n1       int64
	n2       int64
	expected int64
}{
	{5, 2, 7},
	{4, 5, 9},
	{120, 1, 121},
}

// TestSumInt8
func TestSumInt8(t *testing.T) {
	for _, v := range sum_tests_int8 {
		if val := SumInt8(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

// TestSumUint32
func TestSumUint32(t *testing.T) {
	for _, v := range sum_tests_uint32 {
		if val := SumUint32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

// TestSumInt32
func TestSumInt32(t *testing.T) {
	for _, v := range sum_tests_int32 {
		if val := SumInt32(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}

// TestSumFloat64
func TestSumFloat64(t *testing.T) {
	for _, v := range sum_tests_float64 {
		if val := SumFloat64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%f, %f) returned %f, expected %f", v.n1, v.n2, val, v.expected)
		}
	}
}

// TestSumInt64
func TestSumInt64(t *testing.T) {
	for _, v := range sum_tests_int64 {
		if val := SumInt64(v.n1, v.n2); val != v.expected {
			t.Errorf("Sum(%d, %d) returned %d, expected %d", v.n1, v.n2, val, v.expected)
		}
	}
}
