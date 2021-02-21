package Repeat_01

//697
func findShortestSubArray(nums []int) int {
	m:=make(map[int][3]int)
	max:=1
	for i:=range nums {
		if v, ok := m[nums[i]]; ok {
			m[nums[i]]=[3]int{v[0],i,v[2]+1}
			if v[2]+1 > max {
				max = v[2]+1
			}
		} else {
			m[nums[i]] = [3]int{i, i, 1}
		}
	}
	//fmt.Println(m,max)
	shortest:=len(nums)
	for _,arr:=range m{
		if arr[2]==max{
			if shortest>arr[1]-arr[0]+1{
				shortest=arr[1]-arr[0]+1
			}
		}
	}
	return shortest
}

//sliding window
func findShortestSubArraySW(nums []int) int {
	m:=make(map[int]int)
	maxFreq:=0
	for i:=range nums{
		m[nums[i]]+=1
		if maxFreq<m[nums[i]]{
			maxFreq=m[nums[i]]
		}
	}
	l,r:=0,0
	m2:=make(map[int]int)
	ans:=len(nums)
	for r<len(nums){
		m2[nums[r]]++
		if m2[nums[r]]==maxFreq{

		}
		for m2[nums[r]]==maxFreq{
			if ans>r-l+1{
				ans=r-l+1
			}
			m2[nums[l]]--
			l++
		}
		r++
	}
	return ans
}