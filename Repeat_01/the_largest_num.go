package Repeat_01

import (
	"fmt"
	"sort"
	"strconv"
)

//179

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sy*x+y > sx*y+x
	})
	if nums[0] == 0 {
		return "0"
	}
	ans := []byte{}
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}
//bad didn't work
func largestNumberII(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		stacki:=[]int{}
		stackj:=[]int{}
		for i%10!=0&&i/10!=0{
			stacki=append(stacki,i%10)
			i=i/10
		}
		for j%10!=0&&j/10!=0{
			stackj=append(stackj,j%10)
			j=j/10
		}
		for len(stacki)!=0&&len(stackj)==0&&
			stacki[len(stacki)-1]==stackj[len(stackj)-1] {
			stacki = stacki[:len(stacki)-1]
			stackj = stackj[:len(stackj)-1]
		}
		if len(stacki)==0&&len(stackj)!=0{
			if stackj[len(stackj)-1]!=0{
				return false
			}
			return true
		}
		if len(stackj)==0 {
			return true
		}
		return stacki[len(stacki)-1]>stackj[len(stackj)-1]
	})
	res:=""
	for i:=range nums{
		res+=fmt.Sprintf("%d",nums[i])
	}
	return res
}

