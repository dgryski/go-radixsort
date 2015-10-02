package radixsort

import (
	"math/rand"
	"sort"
	"testing"
)

func benchmarkRadix(b *testing.B, size int) {

	rand.Seed(0)

	elems := make(Elems, size)

	for i := range elems {
		elems[i].Key = uint64(rand.Int63())
		elems[i].Val = elems[i].Key * elems[i].Key
	}

	shuf := make(Elems, size)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		copy(shuf, elems)
		Sort(shuf)
	}
}

type Elems []Elem

func (e Elems) Len() int           { return len(e) }
func (e Elems) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }
func (e Elems) Less(i, j int) bool { return e[i].Key < e[j].Key }

func benchmarkSort(b *testing.B, size int) {

	rand.Seed(0)

	elems := make(Elems, size)

	for i := range elems {
		elems[i].Key = uint64(rand.Int63())
		elems[i].Val = elems[i].Key * elems[i].Key
	}

	shuf := make(Elems, size)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		copy(shuf, elems)
		sort.Sort(shuf)
	}
}

func BenchmarkRadix10(b *testing.B)    { benchmarkRadix(b, 10) }
func BenchmarkRadix100(b *testing.B)   { benchmarkRadix(b, 100) }
func BenchmarkRadix1000(b *testing.B)  { benchmarkRadix(b, 1000) }
func BenchmarkRadix10000(b *testing.B) { benchmarkRadix(b, 10000) }
func BenchmarkRadix1e6(b *testing.B)   { benchmarkRadix(b, 1e6) }
func BenchmarkRadix10e6(b *testing.B)  { benchmarkRadix(b, 10e6) }
func BenchmarkRadix20e6(b *testing.B)  { benchmarkRadix(b, 20e6) }
func BenchmarkRadix30e6(b *testing.B)  { benchmarkRadix(b, 30e6) }

func BenchmarkSort10(b *testing.B)    { benchmarkSort(b, 10) }
func BenchmarkSort100(b *testing.B)   { benchmarkSort(b, 100) }
func BenchmarkSort1000(b *testing.B)  { benchmarkSort(b, 1000) }
func BenchmarkSort10000(b *testing.B) { benchmarkSort(b, 10000) }
func BenchmarkSort1e6(b *testing.B)   { benchmarkSort(b, 1e6) }
func BenchmarkSort10e6(b *testing.B)  { benchmarkSort(b, 10e6) }
func BenchmarkSort20e6(b *testing.B)  { benchmarkSort(b, 20e6) }
func BenchmarkSort30e6(b *testing.B)  { benchmarkSort(b, 30e6) }

func BenchmarkRadix250(b *testing.B) { benchmarkRadix(b, 250) }
func BenchmarkSort250(b *testing.B)  { benchmarkSort(b, 250) }

func TestSortSmall(t *testing.T) {

	A := []Elem{
		{7, 49},
		{2, 4},
		{3, 9},
		{6, 36},
		{10, 100},
		{4, 16},
		{9, 81},
		{5, 25},
		{1, 1},
		{8, 64},
	}

	Sort(A)

	if !sort.IsSorted(Elems(A)) {
		t.Errorf("A not sorted")
	}
}

func TestSortBig(t *testing.T) {

	rand.Seed(0)

	elems := make(Elems, 1e6)

	for i := range elems {
		elems[i].Key = uint64(rand.Int63())
		elems[i].Val = elems[i].Key * elems[i].Key
	}

	Sort(elems)

	if !sort.IsSorted(elems) {
		t.Errorf("million item array not sorted")
	}
}
