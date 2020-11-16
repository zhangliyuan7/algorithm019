package w2

import (
	"strconv"
)

//这个题目唯一的构思是%3==0&&%5==0 => %15==0
//其他不知道有什么亮点
func fizzBuzz2(n int) []string {
	var r []string
	for i:=1;i<=n;i++{
		if i%15==0{
			r=append(r, "FizzBuzz")
		}else if i%3==0{
			r=append(r,"Fizz")
		}else if i%5==0{
			r=append(r,"Buzz")
		}else{
			r=append(r, strconv.Itoa(i))
		}
	}
	return r
}
