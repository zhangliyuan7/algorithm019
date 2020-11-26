package Week_04

//二分
func isPerfectSquare(num int) bool {
	if num<=1{
		return true
	}
	l,r:=0,num
	for l<=r{
		mid:=l+(r-l)/2
		if mid*mid==num{
			return true
		}
		if mid*mid<num{
			l=mid+1
		}
		if mid*mid>num{
			r=mid-1
		}
	}
	return false
}