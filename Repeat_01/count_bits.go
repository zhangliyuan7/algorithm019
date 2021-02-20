package Repeat_01

// 338

// dp[1]=1
// dp[2]=1
// dp[3]=dp[2]+dp[1]
// dp[4]=1
// dp[5]=dp[4]+dp[1]
// dp[6]=dp[4]+dp[2]
// dp[7]=dp[4]+dp[3]
// dp[8]=1
// dp[9]=dp[8]+dp[1]
// ...
// 1=> dp[i]=dp[i-pow]+dp[pow]
// 2=> dp[i]=dp[i/2]+dp[i%2]

func countBits(num int) []int {
	dp:=make([]int,num+1)
	pow:=1
	for i:=1;i<=num;i++{
		if i==pow{
			dp[i]=1
		}else if i==pow*2{
			dp[i]=1
			pow*=2
		}else{
			dp[i]=dp[pow]+dp[i-pow]
		}
	}
	return dp
}

func countBitsB(num int) []int {
	dp:=make([]int,num+1)
	for i:=1;i<=num;i++{
		// dp[i>>1]==dp[i/2]
		// i&1==dp[i%2]
		dp[i]=dp[i>>1]+i&1
	}
	return dp
}

