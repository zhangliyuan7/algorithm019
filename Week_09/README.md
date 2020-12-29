学习笔记

#####  day1
```
    // 忘掉好多，需要重新写一遍dp了，
    dp题目，尝试分解问题成二维 尤其是字符串问题
    dp table +数学归纳法，LCS问题需要反复看 
    
    
```
##### 188 买卖股票最佳时机
```
    hard problem ,need re-watch
```
##### 121 股票初级
    ```
        锁定利润，每天大于前一天，即买卖
        贪心:
        func maxProfit(prices []int) int {
            get:=0
            for i:=1;i<len(prices);i++{
                if prices[i]-prices[i-1]>0{
                  get+=prices[i]-prices[i-1]  
                }
            }
            return get
        }
        dp:
        func maxProfit(prices []int) int {
                  dp:=make([][2]int,prices)
                  //0 sell 1 buy
                  for i:=1;i<len(prices);i++{
                      dp[i][0]=max(dp[i-1][1]+prices[i],dp[i-1][0])
                      dp[i][1]=dp[i-1][0]-prices[i]
                  }
                  return dp[len(prices)-1][0]
        }
    ```
##### 72 编辑距离
```
    拆解dp比较麻烦，使用两个二维标示两个word
    dp[word1.len][word2.len]表示word1的len长度子串和word2的len长度子串的最小编辑距离
    使用两个word作为维度绘制table 
    dp[i][j]
         if word1[i]==word2[j] 那么 编辑距离不增加，即dp[i-1][j-1]
    否则
         else: dp[i][j]=min(dp[i-1][j],dp[i][j-1],dp[i-1][j-1]) +1
    递推公式不是很好想，先记住这个题目，以后套用
    注意 用word[i]==word[j]索引时，因为i，j从1开始 所以比较用word[i-1]==word[j-1]
    正因为如此，所以整个dp table范围要扩大1 即dp[m+1][n+1],return dp[m][n]
    全部代码：
       func minDistance(word1 string, word2 string) int {
       	m:=len(word1)+1
       	n:=len(word2)+1
       	dp:=make([][]int,m)
       	for i:=range dp{
       		dp[i]=make([]int,n)
       	}
       	for i:=0;i<m;i++{
       		dp[i][0]=i
       	}
       	for i:=0;i<n;i++{
       		dp[0][i]=i
       	}
       	for i:=1;i<m;i++{
       		for j:=1;j<n;j++{
       			// 否则少一次word的子母判断
       			if word1[i-1]==word2[j-1]{
       				dp[i][j]=dp[i-1][j-1]
       			}else{
       				dp[i][j]=minthree(dp[i-1][j-1]+1,dp[i-1][j]+1,dp[i][j-1]+1)
       			}
       		}
       	}
       	return dp[m-1][n-1]
       }
       
       func minthree(a,b,c int )int{
       	candidates:=[]int{a,b,c}
       	min:=math.MaxInt32
       	for i:=0;i<3;i++{
       		if  candidates[i]< min{
       			min=candidates[i]
       		}
       	}
       	return min
       }
  
```
##### 最小花费爬楼梯 746  
```
    此题之前做过
    每次迈1-2步，每次走到那个位置便需要加上对应体力消耗，
    dp[0]=cost[0] 从第一级起步
    dp[1]=cost[1] 从第二级起步
    return min(dp[len-1],dp[len-2])
    因为每次能迈1-2步，所以两个结果取小值
```
##### 121 只能交易一次的买卖股票最佳时机
```
   官方算法比较巧妙，直接一次循环找到最小price和向后的最大差值
   code:
        func MaxProfitOnce(prices []int)int{
        	minp:=math.MaxInt32
        	minprice:=minp
        	maxprofit:=0
        	for i:=0;i<len(prices);i++{
        		maxprofit=max(prices[i]-minprice,maxprofit)
        		minprice=min(prices[i],minprice)
        	}
        	return maxprofit
        }
```