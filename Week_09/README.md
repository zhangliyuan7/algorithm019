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
##### 字符串距离，相似度类问题
    ```
        dp[word1][word2]+二维Table找规律，可破！
    ```
    
##### 回文子串  5
```
    1 动态递推，规则理解了，但是感觉代价很大啊，想的比较费劲,中心扩散更好理解且空间复杂度更低
    2 中心扩散法
        回文串有两种
            1 一种是奇数长度字符串，中心字母只有一个
            2 另外一种就是偶数长度字符串，所有字母都成对出现
        中心扩散就是同时 以每个字母作为奇数字符串中心点 ，以每个字符相邻字符作为偶数回文字符串中心点 
        遍历向外扩散， 不断找到最长的，赋值给结果左右边界
        最终返回string[left,right+1]
        注意扩散函数
        func extend(s ,left,right)(int,int){
            for ;i>=0&&j<=len(s)-1&&s[i]==s[j];left,right=left-1,righ1+1{}
            //因为循环终止条件为越界，所以缩回一步方为合法值
            return left+1,right-1
        }    
```
##### 反转字符串II 541 
```
    循环 ，每次步数为2k，每次反转k个字母 
    主要注意边界
        边界一：
            mid(i+k)大于字符串长度时 全部反转 reverse(s[left:]),break loops
            right(2k+i)大于字符串长度时，使其为 right=len(s) 
        设置left，mid ,right
    
        for i:=0;i<len(s);i+=2*k{ 
            left=i mid=i+k right=i+2*k
            result += reverse(s[left:mid])+s[mid:right]
        }
        return result 
```
##### Atoi 8 
```
    方法1 多条件判断
        1 先去除左右space 如果去除后 s为空 返回0
        2 判断首位 
            switch case 
            首位"+"/"-" sign=1/-1 abs = s[1:]
            首位数字  sign=1 abs=s
            其他 返回空
            循环此时的abs 截断非法部分，只取合法部分
                for i:=0;i<len(s);i++{
                    if s[i]<'0'||s[i]>'9'{
                        abs=abs[:i]
                        break
                    }
                }
            进行转换  // r的计算必须放在前面，否则会有判断不到的情况
                for i:=0;i<len(s);i++{
                    r=r*10+(s[i]-'0')
                    if r>maxint32&&sign==1{return max}
                    if sign==-1 &&sign*r<minint32{return min}
                }
                return sign*r
```       
##### 正则表达式匹配 10 
```
    方法1 ：递归  （可记忆化） 较好理解
    func isMatch(s string, p string) bool {
    	//这个地方注意，不能写成
    	//if len(p)==0&&len(s)!=0{
    	//	return false 
    	//}
    	// 因为会忽略一个true的情况，即s也为0 ，这时，后续逻辑会将firstmatch设置为false，所以整个结果会错误
    	
    	if len(p)==0{
    		return len(s)==0
    	}
    	firstMatch:=false
    	// 一定要判断s长度，避免panic ，判断s的字符是否匹配p字符或者'.'
    	if len(s)!=0&&(s[0]==p[0]||p[0]=='.'){
    		firstMatch=true
    	}
    	// *号情况有两种 ，一种是p的 'n*'中的n出现零次，另外就是出现1-n次
    	// 出现0次的结果 即将p的'n*'去掉，继续匹配s当前字符  s,p[2:]  注意p的长度判断，防止panic
    	// >0次结果，即去掉当前匹配到的s字符，p规则不变，匹配s后续字符 s[1:],p
    	if len(p)>1&&p[1]=='*'{
    		return (firstMatch&&isMatch(s[1:],p))||isMatch(s,p[2:])
    	}
    	return firstMatch&&isMatch(s[1:],p[1:])
    }
    方法2 动态规划 
        没十分理解
    方法3 有限状态机
        +1 
```  
##### 最长上升子序列 300 
```golang
    此题动态规划比较容易理解
    初始化dp比较特殊，dp每个位置都要设置为1 代表每个数字自身组成最短子序列
     然后双重循环依次迭代，子循环自首至当前元素索引
    递推公式比较难想 
    for i:=0;i<length;i++{
        for j:=0;j<i;j++{
        	if nums[i]>nums[j]{
                //即增序，j<i 
        		dp[i]=max(dp[i],dp[j]-1) //与每个可能的序列单独计算，即每个i都要计算子序列的可能性
            }           
        }           
    }
    return for loop ,search max of dp
    //return dp[max]
```
##### 滑动窗口最大值 239 
```golang
    复习滑动窗口一题
    使用双端队列，删除条件是( if q.len>0&&q[0]==n ) 删除nums[i-k+1]
    i <k-1 {
        q.push(nums[i])    
    }else{
        q.push(nums[i])
        r=append(r,q.top)
        q.pop(nums[i-k+1])
    }
    push:
        if len(q)==0{
            q=append(q,n)
        }else{
            for q.len!=0&&q.[q.len-1] < n{
                q=q[:q.len-1]
            }
            q=append(q,n)
        }
    pop:
        if q.len!=0&&q[0]==n{
           q=q[1:]
        }
```
##### 最大子矩阵 面试题17.24
```
    关键在于如何转化为一维，求最大连续子序列的和 
    其实就是将子矩阵 上下边界圈定，然后利用一个数组 ，存储这个圈定范围内的 每列的和
    再对这个圈定数组，求其最大连续子序列的和 
    上下固定之后 ，整个数组包含了整个行的每一列的值，所以 ，连续子序列求出的即是左右边界 
    加上循环的上下边界 ，便可以得到整个矩阵的上下左右边界 
        遍历矩阵行，i 从第1行- 第matrix size行 
        子循环 遍历矩阵行，从外层循环行数 i ，到matrix size
        每次形成一个列和的数组，需要每次行变换时缩减 
        对这个数组进行求最大连续子序列和，一维动态规划 
        求出与max对比 ，大 即记录    
     
    另外注意判断一维连续最大时，要利用tmp更新begin 
     
    func GetMaxMatrix(matrix [][]int) []int {
    	res:=make([]int,4)
    	m:=len(matrix)
    	n:=len(matrix[0])
    	total:=math.MinInt32
    	for i:=0;i<m;i++{
    		for j:=i;j<m;j++{
    			var sumColumn = make([]int,n)
    			for k:=0;k<n;k++{
    				for u:=i;u<=j;u++{
    					sumColumn[k]+=matrix[u][k]
    				}
    			}
    			l,r,max:=searchMaxSubsequence(sumColumn)
    			//fmt.Println(l,r,max)
    
    			if max>total{
    				total=max
    				res=[]int{i,l,j,r}
    				//fmt.Println("res:",res)
    			}
    		}
    	}
    	return res
    }
    func searchMaxSubsequence(sc []int )(left int,right int,max int ){
    	left,right=0,0
    	tmpbegin:=left
    	max=sc[0]
    	dpi:=sc[0]
    	for i:=1;i<len(sc);i++{
    		if dpi>0{
    			dpi+=sc[i]
    		}else{
    			dpi=sc[i]
    			tmpbegin=i
    		}
    		if dpi>max{
    			max=dpi
    			left=tmpbegin
    			right=i
    		}
    	}
    	return
    }
    划定上下边界，在利用连续最大子序和的计算方式，求出一维最大的左边界和右边界，sum 与 max比较确定是否替换结果集[top,left,bottom,right]
    前两重循环 用来界定上下界限，第三，四个循环 用来整合column的和数组，searchMaxSubsequence就是获取左右边界了，与一维的问题完全相同，只不过将left，right，max也要返回
    比较，确定最大的边界集合，返回即可

``` 
