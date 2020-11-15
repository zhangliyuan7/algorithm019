package Week_02

import (
	"strings"
)

func RemoveKdigits(num string, k int) string {
	count := k
	r := []byte{}
	for i, _ := range num {
		for count>0&&len(r) != 0 && r[len(r)-1] > num[i] {
			r = r[:len(r)-1]
			count--
		}
		r = append(r, num[i])
	}
	if count!=0{
		r =r[:len(r)-count]
	}
	return strings.TrimLeft(string(r),"0")
}

//用这种方式转换大数字 溢出
//func Build(r []int )string {
//	var num int
//	for i:=range r {
//		num+=int(r[i])*int(math.Pow10(len(r)-1-i))
//	}
//	return strconv.Itoa(num)
//}
