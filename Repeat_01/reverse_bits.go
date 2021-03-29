package Repeat_01

import "fmt"
// 190
func reverseBits(num uint32) uint32 {
	fmt.Println(num)
	var newn uint32
	power:=31
	for num!=0{
		x:=num&1
		newn+=x<<power
		num=num>>1
		power-=1
	}
	return newn
}
