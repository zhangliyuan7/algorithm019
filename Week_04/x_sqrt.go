package Week_04

import "math"

// 69
//笨迭代
func mySqrt(x int) int {
	if x<1{
		return 0
	}
	var n = 1
	for {
		if n*n==x{
			return n
		}
		if ((n-1)*(n-1))<x&&n*n>x{
			break
		}
		n+=1
	}
	return n-1
}

//2 fen
func mySqrtErfen(x int) int {
	l, r := 0, x
	for l <= r {
		mid := l + (r - l) / 2
		if mid*mid==x{
			return mid
		}
		if mid * mid < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l-1
}
//jisuanqi 库函数算x的1/2次mi
func mySqrtunknown(x int) int {
	if x == 0 {
		return 0
	}
	ans := int(math.Exp(0.5 * math.Log(float64(x))))
	if (ans + 1) * (ans + 1) <= x {
		return ans + 1
	}
	return ans
}

//newton
func mySqrt3(x int) int {
  if x==0{
  	return 0
  }
  c,x0:=float64(x),float64(x)
  for {
  	xi:=0.5*(x0+c/x0)
  	if math.Abs(x0-xi)< 1e-7{
  		break
	}
	x0=xi
  }
  return int(x0)
}