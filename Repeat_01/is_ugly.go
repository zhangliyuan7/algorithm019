package Repeat_01


//263
func isUgly(n int) (ans bool) {
	if n==0{
		return false
	}
	if n==1{
		return true
	}
	switch {
	case n%2==0:
		ans = ans || isUgly(n/2)
	case n%3==0:
		ans = ans || isUgly(n/3)
	case n%5==0:
		ans = ans ||  isUgly(n/5)
	}
	return
}

