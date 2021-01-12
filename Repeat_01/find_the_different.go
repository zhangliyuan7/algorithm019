package Repeat_01

// 389

//异或 nb
func findTheDifferenceA(s, t string) (diff byte) {
	for i := range s {
		diff ^= s[i] ^ t[i]
	}
	return diff ^ t[len(t)-1]
}

//求和方法
func findTheDifferenceB(s,t string )(byte){
	ss:=0
	for _,c:=range s{
		ss+=int(c)
	}
	st:=0
	for _,c:=range t {
		st+=int(c)
	}
	return byte(st-ss)

}

func FindTheDifference(s string, t string) byte {
	m:=make(map[byte]int )
	for _,c:=range s{
		m[byte(c)]+=1
	}
	for _,c:=range t{
		if _,ok:=m[byte(c)];ok{
			m[byte(c)]-=1
			if m[byte(c)]<0{
				return byte(c)
			}
		}else{
			return byte(c)
		}
	}
	return byte(0)
}
