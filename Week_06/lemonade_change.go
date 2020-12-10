package Week_06

func lemonadeChange(bills []int) bool {
	five := 0
	ten :=0
	twenty:=0
	for i := 0; i < len(bills); i++ {
		switch bills[i]{
		case 20:
			twenty+=1
			if ten>=1&&five>=1{
				ten-=1
				five-=1
			}else if five>=3 {
				five -= 3
			}else{
				return false
			}
		case 10:
			ten+=1
			if five==0{
				return false
			}
			five-=1
		case 5:
			five+=1
		}
	}
	return true
}
