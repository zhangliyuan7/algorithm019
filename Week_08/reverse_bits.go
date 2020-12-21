package Week_08

// 190
func reverseBits(num uint32) uint32 {
	var newn uint32
	power:=31
	for num!=0{
		newn+=num&1<<power
		num=num>>1
		power-=1
	}
	return newn
}


func reverseBits2(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		result += num >> i << 31 >> i
	}
	return result
}

func ReverseBits(num uint32) uint32 {
	ret := uint32(0)
	power := uint32(31)
	for num != 0 {
		ret += (num & 1) << power
		num = num >> 1
		power -= 1
	}
	return ret
}
