package Repeat_01


// 781
// 对应answers里相同大小的数字，如果它们的出现的次数小于等于它们的值加一
// 那么它们可能是相同颜色的兔子，而超出值+1范围的就肯定是其他颜色的兔子
func numRabbits(answers []int) (count int) {
	m:=make(map[int]int)
	for _,c:=range answers{
		m[c]++
	}
	for x,y:=range m{
		// 表达除该兔子外有x个相同颜色的兔子的个数 和实际这么说过的兔子的个数 做比较
		// 一层一层去除某种颜色可能的数量，再单独计算
		for y>x+1{
			count+=x+1
			y-=x+1
		}
		count+=x+1
	}
	return
}


func numRabbitsII(answers []int) (count int) {
	m:=make(map[int]int)
	for _,c:=range answers{
		m[c]++
	}
	for n,y:=range m{
		// 颜色数
		cn:=(n+y)/(n+1)
		// 颜色数 乘以个数
		count+=cn*(n+1)
	}
	return
}
