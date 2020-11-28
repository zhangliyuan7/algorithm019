学习笔记
##### 最少箭只射爆气球问题
```
    1 > 排序气球，以左边界或右边界为排序依据
    2 > 起始气球的右边界作为第一支箭出发点，尝试射穿多少气球，即气球的左边界小于箭的x坐标
    3 > 当气球左边界不符合小于射箭位置x的时候，以新的不符合条件的气球右边界作为下一支箭的出发点，继续上述步骤
    4 > 贪心，排序
```

##### 深度优先遍历，广度优先遍历模版
```
    深度优先template:
    func dfs_recursion(root *Node){
        if root==nil {return }
        for {
            root.view()
            visited.add(root)
            dfs(root.left)
            dfs(root.right)
        }
    }
    func dfs_with_stack(){
        var s stack 
        var visited
        if visited||root==nil{return}
        for {
            for root!=nil{
                root.visited
                s.add(root.left)
                root=root.left
            }
            if s!=nil{
                root=s.pop
                s=s[:len-1]
                root=s.right
            }
        }
    }
    n-tree dfs:
    func dfs(node,visited){
        if node in visited {return }
        visited.add(node)
        for next_node :=range  node.children(){
            if next_node not in visited{
                dfs(next_node,visited)
            }
        }
    }
    func dfs(node){
        if node==nil{return }
        visited,stack=[]int{} []int{}
        for stack.len!=0{
            node=stack.top
            visited.add(node)
            process(node)
            nodes=generateRelatedNodes(node)
            stack.push(nodes)
        }
    }

    广度优先template:
    func bfs(root *Node){
        var q queue
        if visited( return )
        for q!=nil{
            root=q[0]
            root.visited
            q.add(root.left)
            q.add(root.right)
        }
    }
```
##### 全排列
```
    没感觉这个问题要用深度优先搜索，感觉凭空增加编写和思考复杂度，有点反人类
    直接递归某固定数字之外的其他数字的全排列，返回结果合成结果集即可
    时间还是可以接受的 0ms ，内存占用较大，开了较多数组和递归栈，3.3MB
    func Permute(nums []int) [][]int {
   	    var r [][]int
   	    if len(nums)==1{
   		    return [][]int{nums}
   	    }
   	    for i:=0;i<len(nums);i++{
   		    var others []int  
   		    others=append(others,nums[:i]...)
   		    others=append(others,nums[i+1:]...)
   		    tmpR:=Permute(others)
   		    for x:=0;x<len(tmpR);x++{
   			    tmpR[x]=append(tmpR[x],nums[i])
   		    }
   		    r=append(r,tmpR...)
   	    }
   	    return r
   }
```

##### 完全二叉树的节点数
```
    1 > bfs,dfs,前中后序都可以，暴力法
    2 > 二分法，每次判断左子树和右子树的最左深度，所以左右子树分离时，左右两个子树中，一定有一边子树是完全二叉树
            1.如果左右深度相同，因为是完全二叉树，那么左边一定是满二叉树
              左子树是完全二叉树，所以个数直接可以计算(1<<左子树深度)-1（即2的n次方-1） ,即16-1
              右子树一定是不一定完全的子树，所以递归右子树，计算数量
            2.如果左右深度不同，因为是完全二叉树，那么左子树一定是比右子树深度深
              那么左子树一定是不一定完全的子树，右子树一定完全，所以
              递归左子树，求数量，加上完全右子树数量（1<<右子树深度)-1（即2的n次方-1）
            3.因为每层递归，都是以子树为满数量计算，所以子树的父节点是没有计算到其中的，这个时候需要函数每层层级结果+1
        所以，二分法最终结果为：
              若left深度==right深度 return calculate(root.right)+(1<<left)-1 +1(这是当前层的根节点)
              若left深度!=right深度 return calculate(root.left)+(1<<right)-1 +1(当前层root节点)
```

##### all sort visit b-tree
```
https://medium.com/@houzier.saurav/dfs-and-bfs-golang-d5818ec690d3
```

##### 二叉树层级遍历
    ```
    BFS遍历，每次取出queue中所有元素，将取出的val([]int)append到结果集，最后塞入当前层的所有left，right到干净的queue
    ```

##### 贪心算法
    ```
    Sum（局部最优）==>全局最优
    重点是判断是否贪心可解决，或者，如果不需要太过于精确的结果作为辅助，贪心可以一战
    通常伴随排序，部分问题需要换切入点来贪心，比如跳跃游戏中的倒序贪心
    如果可用贪心，几乎都是O（n）复杂度，因为不考虑次优解，所以几乎是最快最好的方式
    动态规划存储中间结果，贪心一条路走到黑，所以非贪心场景问题用贪心，结果错误率十分高～～
    ```

##### 1370上升下降字符串
    ```
    跳出条件：len(s)==len(result)
    1> 使用桶排序，桶索引可以是字母ascII码，也可以是asc码-'a',这样桶列表可以为{26}int{}
       每个字母对应位置便可以是该字母出现的次数
       整体循环中，根据题意，先正循环（从前到后->'a'到'z'）迭代桶列表，再倒循环桶列表，取出则记录的数字减一
    2> 使用方向，将s变成字符数组并排序，用方向字符控制每次是从前到后还是从后往前，取出元素则将其自字符数组中删除
       记录中间状态字符(防止添加重复元素)，某一方向至尽头，则改变方向，改变中间状态元素，反向相似逻辑
       相对1方法写起来较为麻烦，但是容易想到
       需要注意的地方：
            1- 每次正向或反向结束后，防重复的中间状态记录要重置 
            2- 取出要删除元素列表中值，改变方向时要每次获取索引起始点
            3- 改方向的条件是遍历当前方向结束
    ```
##### 122 买卖股票的最佳时机
    ```
    每次后一个比前一个贵，则卖掉赚差价，固定住利润
    
    ps: 此题难在如何想到使用贪心算法解决，而不在代码，暂时没想好怎么判断
        我自己总结贪心有很多情况需要排序，如果不排序的几乎都是要用动态规划解决
        只有特殊情况(如数学归纳法或者公式定理可以证明局部最优便是全局最优)
        否则这种每步状态不同的场景第一反应都应该是动态规划～
    ```
##### 二分查找模板
    ```
    func(a []int,target int ){
        left,right=0,len(a)-1
        for left<right{
            mid=left+(right-left)/2 
            //之所以不用(right+left)/2是防止left+right溢出
            if a[mid]==target{
                return true
            }
            if a[mid]<target{
                left=mid+1
            }
            if a[mid]>target{
                right=mid-1
            }
        }
        return false 
    }
    ```
##### 旋转数组最小值
    ```
    此题掉坑，想的情况太多太复杂，不要想一步拆出左单调或者右单调的情况
        1 > 思考情况不需要想太过细致的情况，只需要判断在左在右
        2 > 如果二分的mid比right大 那么最小一定在右边，left=mid+1
            其余情况一定在right左边，所以right可以right--收敛
            或者使用mid==right/mid<right区分
            if right==mid ==>right--
                因为无法确定是否存在数组重复项，导致最小值在mid位置和right中间,所以right收敛的慢一点
            if mid<right ==>right = mid
                至少mid会是比right所有都小的值，因为旋转了数组，那么right如果是单调的，right一定大，
                如果不是单调的，那么mid一定是包含最小值的那一部分单调，right一定大于等于最小值
        3 > 返回结果最终为left位置元素    
    ```
##### 求平方根 69
    ```
       1 > 二分法
            左边界l，右边界r ,初始值 0，目标C
            从l到C之间不断取mid=l+(r-l)/2，计算mid*mid是否小于等于C
            找出最大的小于C的mid*mid，然后返回mid的那个值即可
       2 > 牛顿法
            将求C的平方根看作找函数f(x)=x**2-C函数的零值点
            x0,C 初始值均为C
            x的收敛公式 xi=0.5*(x0+C/x0)
            if math.Abs(xi-x0)<1e-7 一个极小常数，那么认定为渐进幅度已经极小，即认定x0几乎为0值点
            终止循环 返回x0；else： x0=xi 继续迈步逼近
            牛顿法还是记住吧，不管怎么推的，会用知道就行了
            速度是真快，效率是真高，数学果然nb
    ```
##### no367 有效的完全平方数
    ```
        经典2分法，终止条件mid平方==n或者l>r
    ```
##### no279 完全平方数
    ```
        此题动态规划，dp[i]=min(dp[i-1]+1,dp[i-j*j]+1...)
        重点：问题是这个公式不知道为什么会是这样
        如果知道公式就比较好做了，使用一个数组 一直存储对应节点值，用公式推导就行了
        代码如下：
        func numSquares(n int) int {
        	dp := make([]int, n+1)
        	for i := 1; i <= n; i++ {
        		dp[i] = i
        		for j := 1; i-j*j >= 0; j++ {
        			dp[i] = min(dp[i], dp[i-j*j]+1)
        		}
        	}
        	return dp[n]
        }
        拉格朗日的四平方数定理和三平方数定理有点nb，不懂，但是不妨碍我记录一下～～
        四平方数：每个数字都可以表示为不超过四个完全平方数的和，普适性需要最少4个完全平方数的和，特殊数字可以只有一个，即完全平方数
        三平方数：n%8==7,则n至少需要4个完全平方数求和
        //拉格朗日 四平方数 和三平方数定理解法
        func numSquaresMath(n int) int {
        	sqrt := int(math.Sqrt(float64(n)))
        	if sqrt * sqrt == n {
        		return 1
        	}
        
        	for n % 4 == 0 {
        		n /= 4
        	}
        	if n % 8 == 7 {
        		return 4
        	}
        
        	for i := 0; i * i < n; i++ {
        		s := int(math.Sqrt(float64(n-i*i)))
        		if s * s == n - i * i {
        			return 2
        		}
        	}
        	return 3
        }

    ```
##### 454 四数相加
    ```
    四个数组中的数字组成的和为零的组合有多少
    使用map存储A，B，C，D四个数组的其中二个数组的所有值的sum及其个数
    for a:=range A{
        for b:=range B{
            map[a+b]++
        }
    }
    迭代C+D 找到map中C+D的相反数，把相反数的数量加入到结果集中
    for c:=range C{
        for d:=range D{
            r+=map[-(c+d)]
        }
    }
    return r
    直接加map中的值是因为map中记录的count也是A，B中不同索引的和，类似全排列思想
    所以可以直接加，不用担心重复，因为循环保证了不会重复索引位置
    此题标注标签为二分法，没看出来哪儿二分了，A+B=-(C+D)?
    ```
##### 链表倒数第k个值
    ```
        维护两个指针，第一个先走k步，走到之后，第二个指针同步前行，首指针到末尾，第二个指针即为倒数第k个node
    ```