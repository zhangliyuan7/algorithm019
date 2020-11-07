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

