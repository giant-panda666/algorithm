package main

func mergeSort(a []int) []int {
	return mergeSortImpl(a, 0, len(a)-1)
}

func mergeSortImpl(a []int, left, right int) []int {
	if left >= right {
		return nil
	}

	middle := (right + left) / 2
	mergeSortImpl(a, left, middle)
	mergeSortImpl(a, middle+1, right)
	if a[middle] > a[middle+1] {
		merge(a, left, middle, right)
	}
	return a
}

func merge(a []int, left, middle, right int) {
	k, i, j := 0, left, middle+1
	tmp := make([]int, right-left+1, right-left+1)

	for i <= middle && j <= right {
		if a[i] < a[j] {
			tmp[k] = a[i]
			i++
		} else {
			tmp[k] = a[j]
			j++
		}
		k++
	}

	if i > middle {
		for j <= right {
			tmp[k] = a[j]
			j++
			k++
		}
	}
	if j > right {
		for i <= middle {
			tmp[k] = a[i]
			i++
			k++
		}
	}
	for index := left; index <= right; index++ {
		a[index] = tmp[index-left]
	}
}
