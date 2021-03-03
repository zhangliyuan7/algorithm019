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


//
//偶数的二进制1个数超级简单，因为偶数是相当于被某个更小的数乘2，乘2怎么来的？
//在二进制运算中，就是左移一位，也就是在低位多加1个0，那样就说明dp[i] = dp[i / 2]
//奇数稍微难想到一点，奇数由不大于该数的偶数+1得到，偶数+1在二进制位上会发生什么？
//会在低位多加1个1，那样就说明dp[i] = dp[i-1] + 1，当然也可以写成dp[i] = dp[i / 2] + 1

func countBitsB(num int) []int {
	dp:=make([]int,num+1)
	for i:=1;i<=num;i++{
		// dp[i>>1]==dp[i/2]
		// i&1==dp[i%2]
		dp[i]=dp[i>>1]+i&1
	}
	return dp
}


func countBitsI(num int) []int {
	dp:=make([]int,num+1)
	for i:=1;i<=num;i++{
		// dp[i>>1]==dp[i/2]
		// i&1==dp[i%2]I
		dp[i]=dp[i>>1]+i&1
	}
	return dp
}
