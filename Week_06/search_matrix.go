package Week_06

//74
func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix)==0{
		return false
	}
	l,r:=0,len(matrix)-1
	for l<=r{
		mid:=l+(r-l)/2
		if len(matrix[mid])==0{
			return false
		}
		if matrix[mid][0]<=target&&matrix[mid][len(matrix[mid])-1]>=target{
			return half(matrix[mid],target)
		}
		if matrix[mid][0]>target{
			r=mid-1
		}
		if matrix[mid][len(matrix[mid])-1]<target{
			l=mid+1
		}
	}
	return false
}
func half(nums []int,target int )bool{
	l,r:=0,len(nums)-1
	for l<=r{
		mid:=l+(r-l)/2
		if nums[mid]==target{
			return true
		}
		if nums[mid]<target{
			l=mid+1
		}
		if nums[mid]>target{
			r=mid-1
		}
	}
	return false
}