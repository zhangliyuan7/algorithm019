package Week_09

//5
// 中心扩散法
func longestPalindrome(s string) string {
	l,r:=0,0
	for i:=0;i<len(s);i++ {
		l1, r1 := isPalindrome(s, i, i)
		l2, r2 := isPalindrome(s, i, i+1)
		if r2-l2 > r-l {
			r = r2
			l = l2
		}
		if r1-l1>r-l{
			r = r1
			l = l1
		}
	}
	return s[l:r+1]
}

func isPalindrome(s string,i,j int )(int ,int ) {
	for ;i>=0 && j<len(s) &&s[i]==s[j];i,j=i-1,j+1{}
	return i+1,j-1
}
