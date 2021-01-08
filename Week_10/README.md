学习笔记

##### 86 分隔链表
```golang 
    两个哨兵，其他与合并有序链表相类似，注意斩断拼接的节点与其next 即可
```
##### 300 最长上升子序列
```golang 
    首先另每个dp[i]=1 ，每个字符单独都是一个序列，最低为1 
    3个关键循环 
    Fir-loop：
       for i:=range dp{
            dp[i]=1
       }
    Sec-loop:
       for i:=0;i<length;i++{
            for j:=0;j<i;j++{
                 if nums[j]<nums[i]{ dp[i]=max(dp[i],dp[j]+1)}
            }
        }
    Thir-loop:
        max:=0
        for i:=range dp{
            if dp[i]>max{max=dp[i]}
        }
    return max
```
##### lcs ，edit distance 
```
    lcs 不同时取最长，相同时+1 
    editdistance 初始化dp 稍微特殊，dp[i]=i 不同时取最小+1 相同时不变，延续
```
##### tries ,并查集 
```
    tries 方法  searchPrefix(pre string  )bool,searchWord(w string )bool ,insert(w string )
    BCJ 通常的初始化 parents[i]=i ,方法 union（a, b int ）,parents (a )int
    tries 有些题目带上word作为节点值较好，但是较占内存
    bcj 应用场景是判断圈子数量或者相关集合数量 如何创建一个能够涵盖所有基础点的并查集列表是一个关键点，其思想和方法较为简单
    另外 ，并查集注意剪枝，每次调用parents都可以做剪枝操作，即去掉中间层级
```
##### LRU 
```golang 
    通常的实现structure:
        type LRU struct{
            cap int 
            length int 
            searchkey map[int]*linkedlist 
            head *linkedlist 
            tail *linkedlist 
        }
        type linkedlist struct {
            pre *linkedlist 
            next *linkedlist 
            key int 
            value int 
        }
        主要需要注意moveOut，moveToHead操作时的链表链接操作，因为是双向链表，前后都要链接，这点比较重要
        head和tail即哨兵，永远不变，通常不计算在length和cap之内
```
##### n queen 51
```golang 
    x&(-x) 最低位的1表示的值 
    x&(x-1) 消掉最低位的1 
    bits.Onecount 返回1的个数
    位运算 ：
        1 先令每个位置都是-1 初始化

           	solution =[][]string{}
           	queens:=make([]int,n)
           	for i:=range queens{
           		queens[i]=-1
           	}
           	solve(queens,n,0,0,0,0)
           	return solution

        2 子函数终止条件 当row==n 即生成board 添加到结果集（可循环到row==n 即每个皇后都有合适位置）
            if n==row{
            		board:=generateBoard(queens,n)
            		solution=append(solution,board)
            		return
            	}
        3 获取当前可用位置集合
            available= ((1<<n)-1)& (^(column|lefts|rights))
            全位置集合 左移1 n次，-1 另n个位置都是1 ，取当前不可占用位置的集合的反集合（或运算求并集，取并集反），求出剩余可用位置
        4 循环可用位置，尝试占用，
            for available！=0 // 存在可用位置，倒序逐个尝试{
                position：=available & (-available) //倒序尝试可用位置
                available = available & (available -1 ) // 占用
                column:=bit.onesCount(position-1 ) // 获取1的个数 即位置
                queens[row]=column
                // 本层占用确定，进入下层
                // 下层的列占用，左斜线占用和右斜线占用扩大，row+1 
                solve（queens,n,row+1,columns｜postion ,(lefts|postion)<<1,(rights|postion)>>1 ）
                //还原本层 回溯
                queens[row]=-1
            }
        5 结果集生成棋盘 ，利用byte数组进行Q的填充操作
            func generateBoard(queens []int,n int )[]string{
            	board:=[]string{}
            	for i:=0;i<n;i++{
            		row:=make([]byte,n)
            		for j:=0;j<n;j++{
            			row[j]='.'
            		}
            		row[queens[i]]='Q'
            		board=append(board,string(row))
            	}
            	return board
            }
```
##### largest rectangle area 84
```golang

    1 regular solution  O(n^2)
        double range loop search left and right , calculate the area and compare 
    2 one level loop with  stack ,get left 
        one level loop with stack ,get right 
        last loop calculate the area by ( height[i]*(right - left -1 )),compare area ,get max area
        通过单调栈获取单个边界
        左边界：
            循环索引 ，当栈为空时 ，用-1 作为左边界，即最左边，如果非空 左边界即为栈顶元素（是个索引号）
            当新元素入栈导致 栈将非单调时，不断出栈，栈定即为左边界
            左边界 正循环
            因为是栈 所以栈顶倒序，所以是左边界
            for i:=0;i<n;i++{
                // attention ,it's a  loop , is "FOR " not "IF "
                // 导致非单调
                for len(s)>0 && heights[stack[len(stack)-1]]>=heights[i]{
                       s=s[:len[s]-1]
                }
                if len(s)==0{
                        left[i]=-1 
                }else{
                        left[i]=s[len(s)-1]
                }
                s=append(s,i)
            }
        右边界：
            循环顺序，倒序，找右边界
            for i:=n-1;i>=0;i++{
                     // attention ,it's a  loop , is "FOR " not "IF "
                     // 导致非单调
                     for len(s)>0 && heights[stack[len(stack)-1]]>=heights[i]{
                            s=s[:len[s]-1]
                     }
                     if len(s)==0{
                            //stack is empty ,or clean 
                            //so ,the right is length of rectangle 
                             right[i]=n 
                     }else{
                             right[i]=s[len(s)-1]
                     }
                     s=append(s,i)
            }
        got max area :
            for i:=0;i<n ;i++{
                area=max(area,heights[i]*(right[i]-left[i]-1))
            }           
            return area 
```
##### back order  145 
```
    双栈方式最为直观 looks like picture ！
    code：
        func postorder(node *tree)[]int{
            r :=[]int{}
            if node==nil{ return r } // 异常情况
            // 第一个栈是用来确定添加到s2栈的元素顺序
            s1:=[]*tree{}
            s2:=[]*tree{}
            // 初始s1 遍历对象
            s1=append(s1,node)

            for s1!=nil{
                top:=s1[-1]
                s1=s1[:-1]
                // 添加到后序遍历队列 s2
                s2=append(s2,top)
                // 添加到s1处理队列
                if top.left!=nil{
                    s1=append(s1,top.left)
                }
                // 添加到s1处理队列
                if top.right!=nil{
                    s1=append(s1,top.right)
                }
            }
            // 真正的后序遍历
            for i:=len(s2)-1;i>=0;i--{
                r=append(r,s2[i].value)
            }
            return r 
        }
```
##### bit calculate 201 
```golang 
    search the intersection from range [m,n]
    regular way :
        for i = [m, n] {
            r=r&i
        }
        return r 
    smart way :
        //找出最大(n)最小(m)数字的二进制公共同位前缀，即为结果
        shift:=0
        for m<n{
            m,n=m>>1,n>>1 //两者都不断右移，砍掉尾巴，直到相同
            shift++ //移动次数
        }
        return m<<shift //将移动次数+当前仅剩的前缀，补齐位数
```
##### 399 除法求值
```
    并查集 
    初始化 ：
        func newbcj(l int )*bcj{
        	bcj:=&bcj{
        		make([]int,l),
        		make([]float64,l),
        	}
        	for i:=range bcj.fa{
        		bcj.fa[i]=i
                //每个字符都是一个独立单元，其初始权重为1 
        		bcj.w[i]=1
        	}
        	return bcj
        }

    找father：
        func (b *bcj)findParents(x int )int {
        	if b.fa[x]!=x{
        		ff:=b.findParents(b.fa[x])
                //权重累乘，上叠
        		b.w[x]*=b.w[b.fa[x]]
        		b.fa[x]=ff
        	}
        	return b.fa[x]
        }

    union：（from,to ,value）
        func (b *bcj)union(from,to int,value float64 ){
        	ff:=b.findParents(from)
        	ft:=b.findParents(to)
            //权重公式 
        	b.w[ff]=value*b.w[to]/b.w[from]
            //ff 's father is ft
        	b.fa[ff]=ft
        }
    search result :
            // 组成每个字符的ID map
        	id :=make(map[string]int)
        	for _,eq:=range equations{
        		fir,sec:=eq[0],eq[1]
        		if _,has:=id[fir];!has{
        			id[fir]=len(id)
        		}
        		if _,has:=id[sec];!has{
        			id[sec]=len(id)
        		}
        	}
            // 初始化并查集
        	bcj:=newbcj(len(id)) 
            // 并查集合并点
        	for i,eq:=range equations{
        		bcj.union(id[eq[0]],id[eq[1]],values[i])
        	}
        	ans:=make([]float64,len(queries))
            // 计算结果
        	for i,q:=range queries{
        		start,hasX:=id[q[0]]
        		end,hasY:=id[q[1]]
                // 存在且在同一域
        		if hasX && hasY&& bcj.findParents(start) == bcj.findParents(end){
                    // 做除法
        			ans[i]=bcj.w[start]/bcj.w[end]
        		}else{
                    // 无关者
        			ans[i]=-1
        		}
        	}
        	return ans
```
##### 115 不同子序列 hard
```
    string s 的子序列中包含多少string t ，返回个数
    外层循环t 内层循环s
        sl = length.s
        tl= length.t
    dp:=make([][]int,tl+1)
    for i:=range dp{
        dp[i]=make([]int,sl+1)
    }
    for i:=range dp[0]{
        // 空字符串是所有长度s的子序列
        dp[0][i]=1
    }
    状态转移 
       if t[i-1]=s[j-1]
            // 
            dp[i][j]=dp[i][j-1]+dp[i-1][j-1]
       else 
            dp[i][j]=dp[i][j-1]

    return dp[tl][sl]
```
##### 旋转数组 189 
```
    移动k步 即等于 移动 k%len(nums)步
    方法一： 原地，但是时间复杂度很高
        一个函数用来向右循环一步 
        移动k次即可
        代码如下：
            // stupid way ，just roll
            func Rotate(nums []int, k int)  {
            	m:=len(nums)
            	k=k%m
            	for ;k!=0;k--{
            		Roll(nums)
            	}
            }
            
            func Roll(nums []int ){
            	tmp:=nums[0]
            	for i:=1;i<len(nums);i++{
            		nums[i],tmp=tmp,nums[i]
            	}
            	nums[0]=tmp
            }
    方法二 翻转三次即可 效率较高，速度很快
        定义翻转函数，用来将长度为n的数组 整体reverse 
        func Reverse(nums []int){
        	m:=len(nums)
        	for i:=0;i<m/2;i++{
        		nums[i],nums[m-i-1]=nums[m-i-1],nums[i]
        	}
        }
        k仍然等于 k%len(nums)
        先整体翻转，然后再将k前面的部分和k后面的部分分别翻转
        code : 
            func RotateSmart(nums []int ,k int ){
            	m:=len(nums)
            	k=k%m
            	Reverse(nums)
            	Reverse(nums[:k])
            	Reverse(nums[k:])
            }
```