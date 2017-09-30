package main

func quickSort(a []int) []int {
	return quickSortImpl(a, 0, len(a)-1)
}

func quickSortImpl(a []int, l, r int) []int {
	// 在数量少的时候，也可以采用插入排序
	if l >= r {
		return nil
	}

	p := partition(a, l, r)
	quickSortImpl(a, l, p-1)
	quickSortImpl(a, p+1, r)

	return a
}

func partition(a []int, l, r int) int {
	v := a[l]
	j := l
	for i := l + 1; i <= r; i++ {
		if a[i] <= v {
			j++
			a[j], a[i] = a[i], a[j]
		}
	}
	a[l], a[j] = a[j], a[l]
	return j
}
