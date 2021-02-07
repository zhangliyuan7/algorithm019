package Repeat_01


//665
// 需要判断连续三项值才可决定
func checkPossibility(nums []int) bool {
	done:=0
	for i:=0;i<len(nums)-1;i++{
		if nums[i]>nums[i+1]{
			done+=1
			if done>1{
				return false
			}
			// 当后一项小于前一项 则调整后一项值
			if i>0&&nums[i+1]<nums[i-1]{
				nums[i+1]=nums[i]
			}
		}
	}
	return true
}
