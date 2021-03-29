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


const (
	m1 = 0x55555555 // 01010101010101010101010101010101
	m2 = 0x33333333 // 00110011001100110011001100110011
	m4 = 0x0f0f0f0f // 00001111000011110000111100001111
	m8 = 0x00ff00ff // 00000000111111110000000011111111
)
//分治算法
func reverseBitsII(n uint32) uint32 {
	n = n>>1&m1 | n&m1<<1
	n = n>>2&m2 | n&m2<<2
	n = n>>4&m4 | n&m4<<4
	n = n>>8&m8 | n&m8<<8
	return n>>16 | n<<16
}

