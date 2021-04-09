package Repeat_01

//153
func findMin(nums []int) int {
	if len(nums)==1{
		return nums[0]
	}
	if len(nums)==0{
		return -1
	}
	l:=0
	r:=len(nums)-1
	if nums[l]<nums[r]{
		return nums[l]
	}
	if nums[r]<nums[l]&&nums[r]<nums[r-1]{
		return nums[r]
	}
	for l<r{
		mid:=l+(r-l)/2
		if nums[mid]<nums[mid-1]&&nums[mid]<nums[mid+1]{
			return nums[mid]
		}
		if nums[mid]>nums[r]{
			l=mid
		}
		if nums[mid]<nums[l]{
			r=mid
		}
	}
	return -1
}

func findMinI(nums []int) int {
	l,r:=0,len(nums)-1
	for l<r{
		mid:=l+(r-l)/2
		if nums[mid]<nums[r]{
			r=mid
		}else{
			l=mid+1
		}
	}
	return nums[l]
}
// has duplicate num
func findMinII(nums []int) int {
	l,r:=0,len(nums)-1
	if r==0{
		return nums[r]
	}
	for l<r {
			mid := l + (r-l)/2
			if nums[mid]>nums[r]{
				l=mid+1
			}else if nums[mid]<nums[r]{
				r=mid
			}else{ //重复元素判断
				r--
			}
		}
	return nums[l]
}