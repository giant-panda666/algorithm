package main

func mergeSortBottomup(a []int) []int {
	for size := 1; size < len(a); size *= 2 {
		for i := 0; i+size < len(a); i += 2 * size {
			merge(a, i, i+size-1, min(i+2*size-1, len(a)-1))
		}
	}
	return a
}

func min(a1, a2 int) int {
	if a1 < a2 {
		return a1
	}
	return a2
}
