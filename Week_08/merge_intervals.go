package Week_08

import (
	"sort"
)

// 56
func MergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		}
		return false
	})
	left := intervals[0][0]
	right := intervals[0][1]
	var r [][]int
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= right {
			right = max(intervals[i][1], right)
			continue
		}
		r = append(r, []int{left, right})
		left = intervals[i][0]
		right = intervals[i][1]
	}
	r=append(r,[]int{left,right})
	return r
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
