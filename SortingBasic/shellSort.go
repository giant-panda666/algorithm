package main

func shellSort(a []int) []int {
	n := len(a)
	if n < 2 {
		return a
	}

	stepSize := n / 2
	for stepSize > 0 {
		for i := stepSize; i < n; i++ {
			j := i
			tmp := a[i]
			for ; j >= stepSize && a[j-stepSize] > tmp; j -= stepSize {
				a[j] = a[j-stepSize]
			}
			a[j] = tmp
		}

		stepSize /= 2
	}

	return a
}
