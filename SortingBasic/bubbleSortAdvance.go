package main

func bobbleSortAdvance(a []int) []int {
	var index int
	for i := len(a) - 1; i > 0; i-- {
		var isOrdered = true
		for j := index; j < i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				isOrdered = false
			}
			if isOrdered {
				index = j
			}
		}
	}

	return a
}
