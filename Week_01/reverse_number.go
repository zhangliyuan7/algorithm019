package practice


//number 7

//good !!!!
func ReverseNumber(x int) int {
	var res int
	var maxInt = 2147483647
	var minInt = -2147483648
	for ;x!=0;x=x/10{
		last :=x%10
		if res>maxInt/10 || res == maxInt/10 && last>8{
			return 0
		}
		if res<minInt/10 || res == minInt/10 && last< -7{
			return 0
		}
		res = res*10 +last
	}
	return res
}






















func reverseNumber(n int )int {
	var r  int
	var maxint = 2147483647
	var minint = -2147483648
	for ;n!=0;n=n/10{
		last:=n%10
		if r>maxint/10||r==maxint&&last>7{
			return 0
		}
		if r<minint/10||r==minint&&last< -8{
			return 0
		}
		r=r*10+last
	}
	return r
}
