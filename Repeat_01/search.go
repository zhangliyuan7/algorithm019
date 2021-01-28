package Repeat_01

//33
func Search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (r-l)/2 + l
		if nums[mid] == target {
			return mid
		}
		//不要忘了等于，判断mid 在左在右
		if nums[0]<=nums[mid]{
			if nums[mid]>target&&target>=nums[0]{
				r=mid-1
			}else{
				l=mid+1
			}
		}else{
			if nums[mid]<target&&target<=nums[len(nums)-1]{
				l=mid+1
			}else{
				r=mid-1
			}
		}
	}
	return -1
}

//if (nums[0] <= target && target < nums[mid]) {
//r = mid - 1;
//} else {
//l = mid + 1;
//}
//} else {
//if (nums[mid] < target && target <= nums[n - 1]) {
//l = mid + 1;
//} else {
//r = mid - 1;
//}
//}
