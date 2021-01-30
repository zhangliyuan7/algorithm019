package Repeat_01

// 81
func search(nums []int, target int) bool {
	nl:=len(nums)
	if nl==1{
		return nums[0]==target
	}
  	l:=0
  	r:=nl-1
  	for l<=r{
  		mid:=(r-l)/2+l
  		if nums[mid]==target{
  			return true
		}
		if nums[mid]==nums[l]{
			l++
		}else if nums[mid]==nums[r]{
			r--
		}else if nums[mid]<nums[r]{
			if target>nums[mid]&&target<=nums[r]{
				l=mid+1
			}else{
				r=mid-1
			}
		}else if nums[mid]>nums[l]{
			if nums[mid]>target&&target>=nums[l]{
				r=mid-1
			}else{
				l=mid+1
			}
		}
	}
	return false
}
