package Repeat_01

import "math"
// 1723
// ans 最短工作时间
// 最多同时处理k个任务
var ans int

func minimumTimeRequired(jobs []int, k int) int {
	ans = math.MaxInt32
	minimumTimeRequiredDFS(jobs, make([]int, k), 0, 0, 0)
	return ans
}

func minimumTimeRequiredDFS(jobs, sum []int, index, max, used int) {
	if max>ans {
		return
	}
	if index == len(jobs) {
		ans = Min(max, ans)
		return
	}
	if used < len(sum) {
		sum[used] = jobs[index]
		minimumTimeRequiredDFS(jobs, sum, index+1, Max(max, jobs[index]), used+1)
		sum[used] = 0
	}
	for i := 0; i < used; i++ {
		sum[i] += jobs[index]
		minimumTimeRequiredDFS(jobs, sum, index+1, Max(max, sum[i]), used)
		sum[i] -= jobs[index]
	}
}


func Max(args ...int) int {
	max := args[0]
	for _, item := range args {
		if item > max {
			max = item
		}
	}
	return max
}

func Min(a , b int )int {
	if a<b{
		return a
	}
	return b
}
