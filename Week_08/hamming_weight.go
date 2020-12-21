package Week_08

//191 1的个数，直接消末尾1 消几次 即为1的个数
func hammingWeight(num uint32) int {
	count:=0
	for num!=0{
		num=num&(num-1)
		count++
	}
	return count
}