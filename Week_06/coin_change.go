package Week_06

//
func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}

	dp := make([]int, amount + 1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, coin := range coins {
			if i - coin < 0 || dp[i - coin] == -1 {
				continue
			}

			count := 1 + dp[i - coin]
			if dp[i] == -1 || dp[i] > count {
				dp[i] = count
			}
		}
	}

	return dp[amount]
}
//322
//search  差一个硬币值的最少的那个值
func CoinChange(coins []int, amount int) int {
	var dp =make([]int,amount+1)
	for i:=1;i<=amount;i++{
		dp[i]=-1
		for j:=0;j<len(coins);j++{
			if i-coins[j]<0||dp[i-coins[j]]==-1{
				continue
			}
			count:=dp[i-coins[j]]+1
			if dp[i]==-1||dp[i]>count{
				dp[i]=count
			}
		}
	}
	return dp[len(dp)-1]
}

func CoinChangeT1(coins []int, amount int) int {
	var dp =make([]int,amount+1)
	dp[0]=0
	for i:=1;i<=amount;i++{
		dp[i]=-1
		for _,k:=range coins{
			//
			if i-k<0||dp[i-k]==-1{
				continue
			}
			count:=dp[i-k]+1
			if dp[i]!=-1||dp[i]>count{
				dp[i]=count
			}
		}
	}
	return dp[amount]
}








