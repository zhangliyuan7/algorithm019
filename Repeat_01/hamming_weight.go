package Repeat_01

//191
func hammingWeight(num uint32) int {
	count:=0
	for num!=0{
		num=num&(num-1)
		count+=1
	}
	return count
}
