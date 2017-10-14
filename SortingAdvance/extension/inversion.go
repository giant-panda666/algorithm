package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5}
	fmt.Println("expect 0, real is", inversion(a, 0, 4))
	var b = []int{5, 4, 3, 2, 1}
	fmt.Println("expect 10, real is", inversion(b, 0, 4))
}

// 查找逆序对的数量，思想和归并排序一致，在merge的时候增加统计逆序对的数量
func inversion(a []int, l, r int) int {
	if l >= r {
		return 0
	}

	mid := (r + l) / 2
	ret1 := inversion(a, l, mid)
	ret2 := inversion(a, mid+1, r)
	ret3 := merge(a, l, mid, r)
	return ret1 + ret2 + ret3
}

func merge(a []int, l, mid, r int) (ret int) {
	tmp := make([]int, r-l+1, r-l+1)
	k, i, j := 0, l, mid+1
	for i <= mid && j <= r {
		if a[i] > a[j] {
			tmp[k] = a[j]
			ret += mid - i + 1
			k++
			j++
		} else {
			tmp[k] = a[i]
			k++
			i++
		}
	}

	for i <= mid {
		tmp[k] = a[i]
		k++
		i++
	}

	for j <= r {
		tmp[k] = a[j]
		k++
		j++
	}

	for i := l; i <= r; i++ {
		a[i] = tmp[i-l]
	}

	return
}
