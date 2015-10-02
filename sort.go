// Package radixsort implements a radix sort
package radixsort

// Elem is an key/value pair
type Elem struct {
	Key uint64
	Val uint64
}

// Sort sorts Elems in O(n) time with O(n) extra space.  This sort is stable.
// It is faster than sort.Sort() for n > ~250.
func Sort(E []Elem) {

	// Algorithm from CLRS, chapter 8

	const m = 256
	n := len(E)

	counts := make([]int, m)

	src := E
	dst := make([]Elem, n)

	for i := uint(0); i < 8; i++ {
		shift := i * 8

		for i := range counts {
			counts[i] = 0
		}

		for i := 0; i < n; i++ {
			j := byte(src[i].Key >> shift)
			counts[j]++
		}

		for j := 1; j < m; j++ {
			counts[j] += counts[j-1]
		}

		for i := n - 1; i >= 0; i-- {
			j := byte(src[i].Key >> shift)
			dst[counts[j]-1] = src[i]
			counts[j]--
		}

		src, dst = dst, src
	}
}
