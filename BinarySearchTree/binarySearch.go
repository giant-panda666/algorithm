package main

func binarySearch(a []int, target int) int {
	l, r := 0, len(a)-1
	for l <= r {
		mid := l + (r-l)/2
		if a[mid] == target {
			return mid
		}

		if a[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

func binarySearch2(a []int, target int) int {
	return __binarySearch2(a, 0, len(a)-1, target)
}

func __binarySearch2(a []int, l, r int, target int) int {
	if l > r {
		return -1
	}

	mid := l + (r-l)/2
	if a[mid] == target {
		return mid
	}
	if a[mid] < target {
		return __binarySearch2(a, mid+1, r, target)
	} else {
		return __binarySearch2(a, l, mid-1, target)
	}
}

func floor(a []int, target int) int {
	l, r := 0, len(a)-1
	if l == r && a[l] == a[r] {
		return l
	}

	for r > l {
		mid := l + (r-l)/2
		if a[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
		if r-l == 1 && r == mid {
			return r
		}
	}
	return l
}

func ceil(a []int, target int) int {
	l, r := 0, len(a)-1
	if l == r && a[l] == a[r] {
		return r
	}

	for r > l {
		mid := l + (r-l)/2
		if a[mid] <= target {
			l = mid
		} else {
			r = mid - 1
		}

		if r-l == 1 && l == mid {
			return r
		}
	}

	return r
}
