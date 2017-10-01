package main

import "math/rand"

func quickSort2(a []int) []int {
	return quickSort2Impl(a, 0, len(a)-1)
}

func quickSort2Impl(a []int, l, r int) []int {
	if l >= r {
		return nil
	}

	p := partition2(a, l, r)
	quickSort2Impl(a, l, p-1)
	quickSort2Impl(a, p+1, r)

	return a
}

func partition2(a []int, l, r int) int {
	tmp := rand.Int()%(r-l+1) + l
	a[tmp], a[l] = a[l], a[tmp]

	v := a[l]
	i, j := l+1, r
	for {
		for i <= r && a[i] < v {
			i++
		}
		for j >= l+1 && a[j] > v {
			j--
		}
		if i > j {
			break
		}
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	a[l], a[j] = a[j], a[l]
	return j
}
