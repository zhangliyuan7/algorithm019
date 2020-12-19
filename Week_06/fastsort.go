package Week_06


func fastsort(nums []int ){
	if len(nums)<=1{
		return
	}
	l,r:=0,len(nums)-1
	move:=1
	mid:=nums[l]
	for l<r{
		if nums[r]<mid{
			nums[r],nums[move]=nums[move],nums[r]
			r--
		}else{
			nums[l],nums[move]=nums[move],nums[l]
			l++
			move++
		}
	}
	fastsort(nums[0:l])
	fastsort(nums[l+1:r])
}