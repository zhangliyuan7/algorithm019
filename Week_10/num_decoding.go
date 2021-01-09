package Week_10

//A-Z
func numDecodings(s string) int {
	if len(s)==0||s[0]=='0'{
		return 0
	}
	dp:=make([]int,len(s)+1)
	dp[0]=1
	dp[1]=1
	for i:=2;i<len(s);i++{
		if s[i]=='0'{
			if s[i-1]=='1'||s[i-1]=='2'{
				dp[i+1]=dp[i-1]
			}else{
				return 0
			}
		}else{
			if s[i-1]=='1'||(s[i-1]=='2'&&s[i]<='6'){
				dp[i+1]=dp[i-1]+dp[i]
			}else{
				dp[i+1]=dp[i]
			}
		}
	}
	return dp[len(s)]
}