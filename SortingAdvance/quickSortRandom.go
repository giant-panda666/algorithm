package main

import (
	"math/rand"
	"time"
)

func quickSortRandom(a []int) []int {
	rand.Seed(time.Now().Unix())
	return quickSortRandomImpl(a, 0, len(a)-1)
}

func quickSortRandomImpl(a []int, l, r int) []int {
	// 在数量少的时候，也可以采用插入排序
	if l >= r {
		return nil
	}

	p := partitionRandom(a, l, r)
	quickSortRandomImpl(a, l, p-1)
	quickSortRandomImpl(a, p+1, r)

	return a
}

func partitionRandom(a []int, l, r int) int {
	// 随机选择partition的比较值，防止在有序情况下，时间复杂度出现极端的O(n^2)
	tmp := rand.Int()%(r-l+1) + l
	a[l], a[tmp] = a[tmp], a[l]

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
