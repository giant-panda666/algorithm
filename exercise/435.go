package main

//Given a collection of intervals, find the minimum number of intervals you need to remove to make the rest of the intervals non-overlapping.
//
//Note:
//You may assume the interval's end point is always bigger than its start point.
//Intervals like [1,2] and [2,3] have borders "touching" but they don't overlap each other.
//Example 1:
//Input: [ [1,2], [2,3], [3,4], [1,3] ]
//
//Output: 1
//
//Explanation: [1,3] can be removed and the rest of intervals are non-overlapping.
//Example 2:
//Input: [ [1,2], [1,2], [1,2] ]
//
//Output: 2
//
//Explanation: You need to remove two [1,2] to make the rest of intervals non-overlapping.
//Example 3:
//Input: [ [1,2], [2,3] ]
//
//Output: 0
//
//Explanation: You don't need to remove any of the intervals since they're already non-overlapping.

/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
type IntervalSlice []Interval

func (s *IntervalSlice) Len() int {
	return len(*s)
}

func (s *IntervalSlice) Less(i, j int) bool {
	if s[i].End <= s[j].End {
		return true
	}
	return s[i].Start <= s[j].Start
}

func (s *IntervalSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func eraseOverlapIntervals(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Sort(IntervalSlice(intervals))

	var res = 1
	var prev = 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start >= intervals[prev].End {
			res++
			prev = i
		}
	}

	return len(intervals) - res
}
