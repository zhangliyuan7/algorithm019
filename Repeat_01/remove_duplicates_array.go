package Repeat_01

func removeDuplicatesZ(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}
	slow, fast := 2, 2
	for fast < n {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func removeDuplicatesI(nums []int) int {
	ln:=len(nums)
	if ln<=2{
		return ln
	}
	last:=ln
	for i:=1;i<ln;i++{
		left:=i-1
		right:=i
		for right<last && nums[left] == nums[right]{
			right++
		}
		if right-left > 2{
			if right==last{
				return left+2
			}
			copy(nums[left+2:],nums[right:])
			last-=right-left-2
		}
	}
	return last
}


func removeDuplicatesII(nums []int) int {
	length := len(nums)
	i := 2
	for i < length {
		if nums[i] == nums[i-1] && nums[i-1] == nums[i-2] {
			copy(nums[i-1:], nums[i:])
			length--
		} else {
			i++
		}
	}
	return length
}
