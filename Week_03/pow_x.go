package w3

// 50 x的n次幂
func MyPow(x float64, n int) float64 {
	if n==0{return 1}
	var pow func(base float64 ,times int )float64
	pow= func(b float64, t int) float64 {
		if t==1{return b} //terminator
		//process
		x:=pow(b,t/2) //sub_question //drill down
		if t%2!=0{
			return x*x*b //reverse
		}
		return x*x //reverse
	}
	// merge
	if n>0{
		return pow(x,n)
	}
	return 1.0/pow(x,-n)
}

