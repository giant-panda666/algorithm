package main

import "math/rand"

func quickSort3Ways(a []int) []int {
	return quickSort3WaysImpl(a, 0, len(a)-1)
}

func quickSort3WaysImpl(a []int, l, r int) []int {
	if l >= r {
		return nil
	}

	lt, gt := partition3Ways(a, l, r)
	quickSort3WaysImpl(a, l, lt-1)
	quickSort3WaysImpl(a, gt, r)

	return a
}

func partition3Ways(a []int, l, r int) (lt, gt int) {
	tmp := rand.Int()%(r-l+1) + l
	a[tmp], a[l] = a[l], a[tmp]

	v := a[l]
	i := l + 1
	lt, gt = l, r+1
	for i < gt {
		if a[i] == v {
			i++
			continue
		}
		if a[i] < v {
			lt++
			a[lt], a[i] = a[i], a[lt]
			i++
			continue
		}
		if a[i] > v {
			gt--
			a[gt], a[i] = a[i], a[gt]
		}
	}

	a[l], a[lt] = a[lt], a[l]
	return
}
