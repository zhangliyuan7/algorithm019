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


func NumberDecodingA(s string )int {
	length:=len(s)
	if length==1&&s[0]!=0{
		return 1
	}
	if s[0]==0{
		return 0
	}
	dp:=make([]int,length+1)
	dp[0]=1
	dp[1]=1
	for i:=1;i<length+1;i++{
		if s[i]=='0'{
			if s[i-1]=='1'||s[i-1]=='2'{
				dp[i+1]=dp[i-1]
			}else{
				return 0
			}
		}else {
			if s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6') {
				dp[i+1]=dp[i-1]+dp[i]
			}else{
				dp[i+1]=dp[i]
			}
		}
	}
	return dp[length]
}

func NumberDecodingB(s string )int {
	l:=len(s)
	if l==1&&s[0]!='0'{
		return 1
	}
	if s[0]=='0'{
		return 0
	}
	for i:=1;i<l;i++{
		if s[i]=='0'&&(s[i-1]=='0'||(s[i-1]!='1'&&s[i-1]!='2')){
			return 0
		}
	}
	dp:=make([]int,l+1)
	dp[0]=1
	dp[1]=1
	for i:=2;i<l+1;i++{
		if s[i-1]!='0'{
			dp[i]=dp[i-1]
		}
		if s[i-2:i]<="26"&& s[i-2:i]>="10"{
			dp[i]=dp[i]+dp[i-2]
		}
	}
	return dp[l]
}








