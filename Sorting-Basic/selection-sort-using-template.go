package main

type T interface {
	Len() int
	Swap(i, j int)
	Less(i, j int) bool
}

func selectionSortUsingTemplate(a T) T {
	var n = a.Len()
	for i := 0; i < n; i++ {
		var minIndex = i
		for j := i + 1; j < n; j++ {
			if a.Less(j, minIndex) {
				minIndex = j
			}
		}
		a.Swap(minIndex, i)
	}

	return a
}

type intSlice []int

func (s intSlice) Len() int {
	return len(s)
}

func (s intSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s intSlice) Less(i, j int) bool {
	return s[i] < s[j]
}
