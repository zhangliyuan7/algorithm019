package Week_06

//376 dp1
func wiggleMaxLength(nums []int) int {
	length:=len(nums)
	if length<2{
		return length
	}
	up,down :=make([]int,len(nums)),make([]int,len(nums))
	up[0]=1
	down[0]=1
	for i:=1;i<len(nums);i++{
		if nums[i]>nums[i-1]{
			up[i]=max376(up[i-1],down[i-1]+1)
			down[i]=down[i-1]
		}else if nums[i]<nums[i-1]{
			down[i]=max376(down[i-1],up[i-1]+1)
			up[i]=up[i-1]
		}else{
			up[i]=up[i-1]
			down[i]=down[i-1]
		}
	}
	return max376(up[length-1],down[length-1])
}

func wiggleMaxLength2(nums []int)int {
	length:=len(nums)
	if length<2{
		return length
	}
	up,down:=1,1
	for i:=1;i<length;i++{
		if nums[i]>nums[i-1]{
			up=max376(up,down+1)
		}
		if nums[i]<nums[i-1]{
			down=max376(down,up+1)
		}
	}
	return max376(up,down)
}
func max376(a,b int)int {
	if a > b {
		return a
	}
	return b
}