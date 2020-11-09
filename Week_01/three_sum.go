package practice

import (
	"sort"
)

// three nums sum equal zero
// a+b+c = 0
//found all the set of (a,b,c)
// two sum + repeat del + sort
func ThreeSum(nums []int) [][]int {
	result := [][]int{}
	BaseSort(nums)
	for i := 0; i < len(nums)-1; i++ {
		var hashM = make(map[int]int)
		target2 := 0 - nums[i]
		for j := i + 1; j < len(nums); j++ {
			if v, ok := hashM[target2-nums[j]]; ok {
				if j != v {
					candidate := []int{nums[i], nums[v], nums[j]}
					if !InList(result, candidate) {
						result = append(result, candidate)
					}
				}
			}
			hashM[nums[j]] = j
		}
	}
	return result
}

//收敛 双指针
func ThreeSum2(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	result := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		target := 0 - nums[i]
		fir := i + 1
		las := len(nums) - 1
		for fir < las {
			if nums[fir] == nums[fir-1] {
				fir++
				continue
			}
			if nums[fir]+nums[las] < target {
				fir++
			}
			if nums[fir]+nums[las] > target {
				las--
			}
			if nums[fir]+nums[las] == target && fir != las {
				if !InList(result, []int{nums[i], nums[fir], nums[las]}) {
					result = append(result, []int{nums[i], nums[fir], nums[las]})
				}
				fir++
				las--
			}
		}
	}
	return result
}

func InList(lists [][]int, candidate []int) bool {
	for _, v := range lists {
		var count int
		for i := 0; i <= 2; i++ {
			if candidate[i] != v[i] {
				break
			}
			count += 1
			if count == 3 {
				return true
			}
		}
	}
	return false
}

//冒泡
func BaseSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
}

//
//procedure quicksort(a, left, right)
//if right > left
//select a pivot value a[pivotIndex]
//pivotNewIndex := partition(a, left, right, pivotIndex)
//quicksort(a, left, pivotNewIndex-1)
//quicksort(a, pivotNewIndex+1, right)
func FastSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	p := left
	temp := nums[p]
	for i := left; i < right; i++ {
		if nums[left] > temp {
			nums[left], nums[i] = nums[i], nums[left]
			left++
		}
	}
	nums[left], nums[p] = nums[p], nums[left]

	FastSort(nums, left, p-1)
	FastSort(nums, p, right)
}

// best
// todo 去重部分还是不正常 ，需要继续研究
func ThreeSum3(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	result := [][]int{}
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		las := len(nums) - 1
		for fir := i + 1; fir < las; {
			sum := nums[fir] + nums[las] + nums[i]
			if sum < 0 {
				fir++
			}
			if sum > 0 {
				las--
			}
			if sum == 0 {
				result = append(result, []int{nums[i], nums[fir], nums[las]})
				fir++
				las--
			}
		}
	}
	return result
}

// repeat result
//todo analyze
func ThreeSumA(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	var last = len(nums) - 1
	var result = [][]int{}
	sort.Ints(nums)
	var tmpj, tmpl int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			return nil
		}
		for j := i + 1; j < len(nums)-1; {
			if nums[j] == tmpj {
				j++
			}
			if nums[last] == tmpl {
				last--
			}
			if j < len(nums)-1 {
				tmpj = nums[j]
			}
			tmpl = last
			if j == last {
				break
			}
			sum := nums[i] + nums[j] + nums[last]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[last]})
				j++
				last--
			}
			if sum < 0 {
				j++
			}
			if sum > 0 {
				last--
			}
		}
	}
	return result
}

func TwoSum(nums []int, target int) []int {
	var hashM = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := hashM[target-nums[i]]; ok {
			return []int{i, v}
		}
		hashM[nums[i]] = i
	}
	return nil
}

// ok ! 排序决定了相同值 相邻，所以可以通过前后比较排除已有相同连续值
// 所以 一定要做每个索引对应值和下一个索引对应值的比较
// 相同需要跳过
func ThreeSumB(nums []int) [][]int {
	if len(nums) < 3  {
		return nil
	}
	var result = [][]int{}
	var sec int
	var thir int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			return result
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		sec = i + 1
		thir = len(nums) - 1
		for sec < thir {
			sum := nums[i] + nums[sec] + nums[thir]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[sec], nums[thir]})
				for sec < thir  &&nums[sec+1] == nums[sec]  {
					sec++
				}
				for thir > sec&& nums[thir-1] == nums[thir]  {
					thir--
				}
				sec++
				thir--
			}
			if sum < 0 {
				sec++
			}
			if sum > 0 {
				thir--
			}
		}
	}
	return result
}

func ThreeSumC(nums []int)[][]int{
	if len(nums)<3{
		return nil
	}
	var result = [][]int{}
	quickSort(nums)
	for i:=0;i<len(nums);i++{
		if nums[i]>0{
			return result
		}
		if i>0&&nums[i]==nums[i-1]{
			continue
		}
		var j = i+1
		k := len(nums)-1
		for j<k {
			sum :=nums[i]+nums[j]+nums[k]
			if sum==0{
				result=append(result,[]int{
					i,j,k,
				})
				for j<k&&nums[k+1]==nums[k]{
					k++
				}
				for k>j&&nums[j]==nums[j-1]{
					j--
				}
				k++
				j--
			}
			if sum>0{
				k--
			}
			if sum <0{
				j++
			}
		}
	}
	return result
}


func quickSort(nums []int){

	var left,right int
	left=0
	right=len(nums)-1
	i := 1
	mid := nums[0]
	for left<right{
		if nums[i]>mid {
			nums[right], nums[i] = nums[i], nums[right]
			right--
		}else{
			nums[left],nums[i] = nums[i],nums[left]
			left+=1
			i++
		}
	}
	quickSort(nums[:left])
	quickSort(nums[left+1:])
}

func ThreeSumD(nums []int )[][]int{
	if len(nums)<3{
		return nil
	}
	var r =[][]int{}
	sort.Ints(nums)
	for i:=0;i<len(nums);i++{
		if nums[i]>0{
			return r
		}
		if i>0&&nums[i]==nums[i-1]{
			continue
		}
		j:=i+1
		k:=len(nums)-1
		for j<k{
			sum:=nums[i]+nums[j]+nums[k]
			if sum==0{
				r=append(r,[]int{nums[i],nums[j],nums[k]})
				if j<k&&nums[j]==nums[j+1] {
					j++
				}
				if k>j&&nums[k]==nums[k-1]{
					k--
				}
				j++
				k--
			}
			if sum<0{
				j++
			}
			if sum>0{
				k--
			}
		}
	}
	return r
}