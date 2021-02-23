package Repeat_01

// 1052
// angry bookshop boss
func maxSatisfied(customers []int, grumpy []int, X int) int {
	l,r:=0,0
	maxSatisfied:=0
	nowSatisfied:=0
	for i:=range customers{
		nowSatisfied+=customers[i]*(1-grumpy[i])
	}
	//fmt.Println("不抑制愤怒的时候，总获得的满意数",nowSatisfied)
	for r<len(grumpy){
		for r<=X-1{
			if grumpy[r]==1{
				nowSatisfied+=customers[r]
			}
			maxSatisfied=nowSatisfied
			r++
			if r==len(grumpy){
				return maxSatisfied
			}
		}
		if grumpy[l]==1{
			nowSatisfied=nowSatisfied-customers[l]
		}
		if grumpy[r]==1{
			nowSatisfied=nowSatisfied+customers[r]
		}
		l++
		r++
		if nowSatisfied>maxSatisfied{
			maxSatisfied=nowSatisfied
		}
	}
	return maxSatisfied
}