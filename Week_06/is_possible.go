package Week_06

// 659 贪心
func IsPossible(nums []int) bool {
	var numsMap =make(map[int]int)
	for i:=0;i<len(nums);i++{
		numsMap[nums[i]]+=1
	}
	var subQ = make(map[int]int )
	for _,n:=range nums{
		if numsMap[n]==0{
			continue
		}
		if subQ[n-1]>0{
			subQ[n]+=1
			subQ[n-1]-=1
			numsMap[n]-=1
			continue
		}
		if numsMap[n+1]>0&&numsMap[n+2]>0{
			numsMap[n+2]-=1
			numsMap[n+1]-=1
			numsMap[n]-=1
			subQ[n+2]+=1
			continue
		}
		return false
	}
	return true
}


func isPossible(nums []int) bool {
	left := map[int]int{} // 每个数字的剩余个数
	for _, v := range nums {
		left[v]++
	}
	endCnt := map[int]int{} // 以某个数字结尾的连续子序列的个数
	for _, v := range nums {
		if left[v] == 0 {
			continue
		}
		if endCnt[v-1] > 0 { // 若存在以 v-1 结尾的连续子序列，则将 v 加到其末尾
			left[v]--
			endCnt[v-1]--
			endCnt[v]++
		} else if left[v+1] > 0 && left[v+2] > 0 { // 否则，生成一个长度为 3 的连续子序列
			left[v]--
			left[v+1]--
			left[v+2]--
			endCnt[v+2]++
		} else { // 无法生成
			return false
		}
	}
	return true
}
