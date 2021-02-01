package Repeat_01

import (
	"container/heap"
	"math"
)

// 1631
type point struct{ x, y, maxDiff int }
//实现heap接口函数
type hp []point
func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].maxDiff < h[j].maxDiff }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(point)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

type pair struct{ x, y int }
var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minimumEffortPathX(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	// for result
	maxDiff := make([][]int, n)
	for i := range maxDiff {
		maxDiff[i] = make([]int, m)
		for j := range maxDiff[i] {
			maxDiff[i][j] = math.MaxInt64
		}
	}
	maxDiff[0][0] = 0
	h := &hp{{}}
	for {
		p := heap.Pop(h).(point)
		// exit
		if p.x == n-1 && p.y == m-1 {
			return p.maxDiff
		}
		//
		if maxDiff[p.x][p.y] < p.maxDiff {
			continue
		}
		for _, d := range dir4 {
			x, y := p.x+d.x, p.y+d.y
			if  0 <= x && x < n && 0 <= y && y < m {
				diff := max(p.maxDiff, abs(heights[x][y]-heights[p.x][p.y]))
				//这步很关键，重复时 更新为取小
				if diff < maxDiff[x][y] {
					maxDiff[x][y] = diff
					heap.Push(h, point{x, y, diff})
				}
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
