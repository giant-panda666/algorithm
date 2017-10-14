package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var a = []int{1, 2, 5, 4, 3, 0, 6, 9, 7, 8}
	fmt.Println("findMax1, ret:", findMaxN(a, 9))
	fmt.Println("findMax2, ret:", findMaxN(a, 8))
}

// 查找第n大的数，思想和快速排序一致，因为每次partition之后，用于区分两部分的数的位置是正确的，
// 所以只需要在这个数的左边或者右边继续查找就行了
func findMaxN(a []int, n int) int {
	if n < 0 || n > len(a)-1 {
		panic("wrong param")
	}

	l, r := 0, len(a)-1
	for {
		p := partition(a, l, r)
		if p == n {
			return a[p]
		}
		if p < n {
			l = p + 1
		} else {
			r = p - 1
		}
	}
}

func partition(a []int, l, r int) int {
	tmp := rand.Int()%(r-l+1) + l
	a[tmp], a[l] = a[l], a[tmp]

	v := a[l]
	// a[l+1...i) <= v, a(j...r] >= v
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
