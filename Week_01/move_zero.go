package practice

//move all zero to end ,keep the sort of other numbers
//核心是填充，不是交换
//[0,1,0,3,12]
func MoveZero(nums []int )[]int{
	var lastNonZero int
	for i:=0;i<len(nums);i++{
		if nums[i]!=0{
			nums[lastNonZero]=nums[i]
			lastNonZero++
		}
	}
	for i:=lastNonZero;i<len(nums);i++{
		nums[i]=0
	}
	return nums
}

func MoveZero2(nums []int )[]int{
	var j int
	for i:=0;i<len(nums);i++{
		if nums[i]!=0{
			if i!=j{
				nums[j]=nums[i]
				nums[i]=0
			}
			j++
		}
	}
	return nums
}

func  MoveZero4(nums []int ){
	var lastindex int
	for i:=0;i<len(nums);i++{
		if nums[i]!=0{
			nums[lastindex]=nums[i]
			lastindex++
		}
	}
	for i:=lastindex;i<len(nums);i++{
		nums[i]=0
	}
}

func MoveZero3(nums []int ){
	var j int
	for i:=0;i<len(nums);i++{
		if nums[i]!=0{
			if i!=j {
				nums[j] = nums[i]
				nums[i] = 0
			}
			j++
		}
	}
}

//tianchong
func MoveZero5(nums []int ){
	var j int
	for i:=0;i<len(nums);i++{
		if nums[i]!=0{
			nums[j]=nums[i]
			j++
		}
	}
	for i:=j;i<len(nums);i++{
		nums[j]=0
	}
}
// double pointer
func MoveZero6(nums []int ){
	var j int
	for i:=0;i<len(nums);i++{
		if nums[i]!=0{
			if i!=j {
				nums[j] = nums[i]
				nums[i] = 0
			}
			j++
		}
	}
}