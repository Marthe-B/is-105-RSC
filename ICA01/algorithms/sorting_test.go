package algorithms

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func init() {
	seed := time.Now().Unix()
	rand.Seed(seed)
}

func perm(n int) (out []int) {
	for _, v := range rand.Perm(n) {
		out = append(out, v)
	}
	return
}

// Testfunksjon Bsort

func BenchmarkBSort100(b *testing.B) {
	benchmarkBSort(100, b)
}

func BenchmarkBSort1000(b *testing.B) {
	benchmarkBSort(1000, b)
}

func BenchmarkBSort10000(b *testing.B) {
	benchmarkBSort(10000, b)
}

func benchmarkBSort(i int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		values := perm(i)
		b.StartTimer()
		Bubble_sort(values)
	}
}

// Testfunksjoner for Quicksort algoritmen
func TestQSort(t *testing.T) {
	values := []int{1, 4, 2, 3, 7, 6, 8, 5, 3, 0}
	expected := []int{0, 1, 2, 3, 3, 4, 5, 6, 7, 8}

	QSort(values)

	if !reflect.DeepEqual(values, expected) {
		t.Fatalf("expected %d, actual is %d", 1, values[0])
	}
}

func BenchmarkQSort100(b *testing.B) {
	benchmarkQSort(100, b)
}

func BenchmarkQSort1000(b *testing.B) {
	benchmarkQSort(1000, b)
}

func BenchmarkQSort10000(b *testing.B) {
	benchmarkQSort(10000, b)
}

func benchmarkQSort(i int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		values := perm(i)
		b.StartTimer()
		QSort(values)
	}
}

//Ny testfunksjon BSortModified
func BenchmarkMBSort100(b *testing.B) {
	benchmarkBSortModified(100, b)
}

func BenchmarkMBSort1000(b *testing.B) {
	benchmarkBSortModified(1000, b)
}

func BenchmarkMBSort10000(b *testing.B) {
	benchmarkBSortModified(10000, b)
}

func benchmarkBSortModified(i int, b *testing.B) {
	for j := 0; j < b.N; j++ {
		b.StopTimer()
		values := perm(i)
		b.StartTimer()
		Bubble_sort_modified(values)
	}
}
