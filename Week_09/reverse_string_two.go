package Week_09

// 541
func reverseStr(s string, k int) string {
	r:=""
	for i:=0;i<len(s);i+=2*k{
		right:=i+2*k
		mid:=i+k
		left:=i
		if mid>len(s){
			r+=reverse(s[left:])
			break
		}
		if right>len(s)-1{
			right=len(s)
		}
		r+=reverse(s[left:mid])+s[mid:right]
	}
	return r
}

func reverse(s string )string{
	r :=""
	for i:=len(s)-1;i>=0;i--{
		r+=string(s[i])
	}
	return r
}