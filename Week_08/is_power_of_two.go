package Week_08

// 231
// n&n-1 清除最后一位的0
func isPowerOfTwo(n int) bool {
	if n==0{return false }
	count:=0
	for n!=0{
		count+=1
		n=n&(n-1)
		if count>1{
			return false
		}
	}
	return true
}

func isPowerOfTwo2(n int) bool {
	if n<=0{
		return false
	}
	return (n&(n-1))==0
}