package Repeat_01

//1006
func clumsy(N int) int {
	s:=[]int{N}
	N--
	index:=0
	for N>0{
		switch index%4{
		case 0:
			s[len(s)-1]*=N
		case 1:
			s[len(s)-1]/= N
		case 2:
			s=append(s,N)
		default:
			s=append(s,-N)
		}
		N--
		index++
	}
	r:=0
	for i:=range s{
		r+=s[i]
	}
	return r
}
