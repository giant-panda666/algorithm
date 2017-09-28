package main

func selectionSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		var minIndex = i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[minIndex] {
				minIndex = j
			}
		}

		a[i], a[minIndex] = a[minIndex], a[i]
	}

	return a
}
