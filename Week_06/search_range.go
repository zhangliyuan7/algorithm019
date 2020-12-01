package Week_06

func SearchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}
	var l, r = 0, len(nums) - 1
	var mid int
	for l <= r {
		mid = l + (r-l)>>1
		if nums[mid] == target {
			break
		}
		if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	var i, j = mid, mid
	for i >= 0 && nums[i] == target {
		res[0] = i
		i = i - 1
	}
	for j < len(nums) && nums[j] == target {
		res[1] = j
		j = j + 1
	}
	return res
}

func SearchRange2(nums []int, target int) []int {
	var l, r = 0, len(nums) - 1
	var mid int
	var res []int = []int{-1, -1}
	for l <= r {
		mid = l + (r-l)>>1
		if target == nums[mid] {
			var i, j int = mid, mid
			for i >= 0 && nums[i] == target {
				res[0] = i
				i -= 1
			}
			for j < len(nums) && nums[j] == target {
				res[1] = j
				j += 1
			}
			break
		}
		if target > nums[mid] {
			l = mid + 1
		}
		if target < nums[mid] {
			r = mid - 1
		}
	}
	return res
}
