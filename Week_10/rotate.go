package Week_10

// 189
// stupid way ，just roll
func Rotate(nums []int, k int)  {
	m:=len(nums)
	k=k%m
	for ;k!=0;k--{
		Roll(nums)
	}
}

func Roll(nums []int ){
	tmp:=nums[0]
	for i:=1;i<len(nums);i++{
		nums[i],tmp=tmp,nums[i]
	}
	nums[0]=tmp
}
//smart way
func RotateSmart(nums []int ,k int ){
	m:=len(nums)
	k=k%m
	Reverse(nums)
	Reverse(nums[:k])
	Reverse(nums[k:])
}
func Reverse(nums []int){
	m:=len(nums)
	for i:=0;i<m/2;i++{
		nums[i],nums[m-i-1]=nums[m-i-1],nums[i]
	}
}
