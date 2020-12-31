package Week_09

import (
	"math"
	"strings"
)

// 8

func myAtoi(s string) int {
	return convert(clean(s))
}
func clean(s string )(sign int,abs string) {
	//space del
	s=strings.TrimSpace(s)
	if s==""{
		return 0,""
	}
	switch  s[0]{
	case '0','1','2','3','4','5','6','7','8','9':
		sign=1
		abs=s
	case '+':
		sign=1
		abs=s[1:]
	case '-':
		sign=-1
		abs=s[1:]
	default:
		return 0,""
	}
	for i:=0;i<len(abs);i++{
		if abs[i]<'0'||abs[i]>'9'{
			//后面的都非法，不用了
			abs=abs[:i]
			break
		}
	}
	return sign,abs
}
func convert(sign int,s string  )int {
	r:=0
	for i:=len(s)-1;i>=0;i--{
		if r>214748364&&(s[i]-'0')>7&&sign==1{
			return math.MaxInt32
		}
		if r>214748364&&(s[i]-'0')>8&&sign==-1{
			return math.MinInt32
		}
		r=r*10+int(s[i]-'0')
	}
	return r
}