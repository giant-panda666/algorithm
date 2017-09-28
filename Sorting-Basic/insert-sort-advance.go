package main

func insertSortAdvance(a []int) []int {
	for i := 1; i < len(a); i++ {
		tmp := a[i]
		j := i
		for j = i; j > 0 && a[j-1] > tmp; j-- {
			a[j-1] = a[j]
		}
		a[j] = tmp
	}

	return a
}
