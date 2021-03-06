学习笔记

##### templates
    ```
        recursion:
        python:
         def recursion(level,param1,param2...):
            //terminator
            if level>max_level:
                process_result
                return 
            //process
            process(level,data)
            //drill down
            self.recursion(level+1,p1,p2...)
            //reverse the current level status if you need
         c/c++:
            void recursion(int level,int param){
                //terminator
                if (level>MAX_LEVEL){
                    return ;
                }
                //process
                process(level,param)
                //drill down
                recursion(level+1,params)
                //reverse the current level status if needed
            }      
            
        dynamic programming:
        
        divide_conquer: //分开，分别-克服，攻克
        python:
        divide_conquer(problem,param1,param2...):
            //terminator
            if problem is None:
                print result
                return 
            //prepare data
            data=prepare_data(problem)
            subproblems=split_problem(problem,data)
            //conquer subproblems
            subresult1=self.divide_conquer(subproblem[1],p1...)
            subresult2=self.divide_conquer(subproblem[2],p1...)
            ......
            //process and generate the final result
            result=process_result(subresult1,subresult2...)
            //revert the current level status
        
        c/c++:
        int divide_conquer(problem *problem,int params){
            // recusion terminator
            if (problem == nullptr){
                process_result
                return return_result;
            }
            //process current problem
            subproblem = split_problem(problem,data)
            subresult1=divide_conquer(subproblem[1],param...)
            ......
            //merge 
            result = process_and_merge(subresult...)
            return 0;
        }
    ```
##### 单词接龙
    ```
        1 bfs 层序遍历思想+循环产生一位变化的新string ，与map中单词比对，存在则加入下层循环查找
          注意使用过的单词要清空，另外 go的正则实在是太tm慢了，不要用，不如循环生成单词直接查map来的快
        2 图方式，不是很明白  
        3 双向bfs ，如果相遇，就是2个方向的遍历层级和/2+1
    ```
##### 排序数组中查找目标元素的第一个位置和最后一个位置 no.34
    ```
        1。二分查找，找到目标target,向左找到左边界，再从mid开始用二分找右边界
        2。二分找到目标值一个索引，然后左右+1/-1扩散，获取两个边界（最差情况O(n))
    ```
##### 最小基因变化 no.433
    ```
        双向交替bfs，从end和start交替出发，找到交汇点之时就是结果层数
        注意点：
            浏览完当前层之后 要从visited和bankMap里去掉 important，防止下层又生成回原值
            visited只存有下一层遍历需要的两个方向的上层对比值 important
            入循环时记录队列长度，append下一层节点后，要截断队列为新队列
            判断start的queue和end的queue哪个长，哪个短则为下次循环的queue
            注意跳过节点自身值
    ```
    
##### 模拟行走机器人 874
    ```
        关键点 两个方向数组 两个数组保存各个方向的步调值
        //direct* = [4]int{north ,east, south, west}
        directForY := [4]int{1,0,-1,0} //y轴坐标四个方向的移动如何处理
        directForX := [4]int{0,1,0,-1} //x轴坐标四个方向的移动如何处理
        每次转弯计算max
        右转方向索引(direct+1)%4，左转(direct+3)%4 (turn to right->right->right == turn to left)
        每次按照方向一步一步走，判断每个点是否存在与障碍map中。
    ```
##### no.74 搜索二维矩阵
    ```
        二分查找，先找到对应矩阵，再对该矩阵二分查找
        注意特殊情况，比如矩阵不为空，但数组为空
        判断依据，因为有序，所以mid的数组最小值<=target&&mid数组最大值>=target即在此数组中有可能存在target
        另外，二分查找一定要记得 mid+1 / mid-1 / l<=r 判断左右边界千万别忘了"="号
    ```
##### 贪心和动态计算
```
    #一、贪心算法###例子假设有1元，5元，11元这三种面值的硬币，给定一个整数金额，比如28元，最少使用的硬币组合是什么？
        分析碰到这种问题，咱们很自然会想起先用最大的面值，再用次大的面值……这样得到的结果为两个11元，一个5元，一个1元，总共是四个硬币。
        思想贪心算法（又称贪婪算法）是指，在对问题求解时，总是做出在当前看来是最好的选择。也就是说，不从整体最优上加以考虑，他所做出的是在某种意义上的局部最优解。
        不足上面的例子，total = 28，得到的“11 11 5 1”恰巧是最优解。假如total = 15呢？total = 15时，结果为“11 1 1 1 1”，共用了五枚硬币。
        但是这只能算是较优解，不是最优解。因为最优解是“5 5 5”，共三枚硬币。所以贪心算法只能保证局部最优（第一枚11就是局部最优），不能保证全局最优。
    #二、动态规划算法咱们仍以15为例，换一种思路，看看如何得到最优解。
        面值为1时，最少需要一个一元硬
        面值为2时，最少需要两个一元硬币
        面值为3时，最少需要三个一元硬币
        面值为4时，最少需要四个一元硬币
        面值为5时，有两个方案:
        ① 在面值为4的基础上加一个1元的硬币，需要五个硬币
        ② 挑一个面值为5元的硬币，需要一个硬币取最小值，需要一个硬币
        面值为6时，两个方案：
        ① 比1元（一个硬币）多了5元（一个硬币），需要两个硬币
        ② 比5元（一个硬币）多了1元（一个硬币），需要两个硬币取最小值，需要两个硬币
        面值为7时，两个方案：
        ① 比1元（一个硬币）多了6元（两个硬币），需要三个硬币
        ② 比5元（一个硬币）多了2元（两个硬币），需要三个硬币取最小值，需要三个硬币
        面值为8时，两个方案：
        ① 比1元（一个硬币）多了7元（三个硬币），需要四个硬币
        ② 比5元（一个硬币）多了3元（三个硬币），需要四个硬币取最小值，需要四个硬币
        面值为9时，两个方案：
        ① 比1元（一个硬币）多了8元（四个硬币），需要五个硬币
        ② 比5元（一个硬币）多了4元（四个硬币），需要五个硬币取最小值，需要五个硬币
        面值为10时，两个方案：
        ① 比1元（一个硬币）多了9元（五个硬币），需要六个硬币
        ② 比5元（一个硬币）多了5元（一个硬币），需要两个硬币取最小值，需要两个硬币
        面值为11时，三个方案：
        ① 比1元（一个硬币）多了10元（两个硬币），需要三个硬币
        ② 比5元（一个硬币）多了6元（两个硬币），需要三个硬币
        ③ 取面值为11元的硬币，需要一个硬币取最小值，需要一个硬币
        面值为12时，三个方案：
        ① 比1元（一个硬币）多了11元（一个硬币），需要两个硬币
        ② 比5元（一个硬币）多了7元（三个硬币），需要四个硬币
        ③ 比11元（一个硬币）多了1元（一个硬币），需要两个硬币取最小值，需要两个硬币
       面值为13时，三个方案：
        ① 比1元（一个硬币）多了12元（两个硬币），需要三个硬币
        ② 比5元（一个硬币）多了8元（四个硬币），需要五个硬币
        ③ 比11元（一个硬币）多了2元（两个硬币），需要三个硬币取最小值，需要三个硬币（
       面值为14时，三个方案：① 比1元（一个硬币）多了13元（三个硬币），需要四个硬币② 比5元（一个硬币）多了9元（五个硬币），需要六个硬币③ 比11元（一个硬币）多了3元（三个硬币），需要四个硬币取最小值，需要四个硬币（15）面值为15时，三个方案：① 比1元（一个硬币）多了14元（四个硬币），需要五个硬币② 比5元（一个硬币）多了10元（两个硬币），需要三个硬币③ 比11元（一个硬币）多了4元（四个硬币），需要五个硬币取最小值，需要三个硬币（16）
       最终，得到的最小硬币数是3。并且从推导过程可以看出，计算一个数额的最少硬币数，比如15，必须把它前面的所有数额（1~14）的最少硬币数都计算出来。这够成了一个递推（注意不是递归）的过程 
       运行结果：
        面值为1的最小硬币数：1
        面值为2的最小硬币数：2
        面值为3的最小硬币数：3
        面值为4的最小硬币数：4
        面值为5的最小硬币数：1
        面值为6的最小硬币数：2
        面值为7的最小硬币数：3
        面值为8的最小硬币数：4
        面值为9的最小硬币数：5
        面值为10的最小硬币数：2
        面值为11的最小硬币数：1
        面值为12的最小硬币数：2
        面值为13的最小硬币数：3
        面值为14的最小硬币数：4
        面值为15的最小硬币数：3
    三、贪心算法与动态规划的区别
    （1）贪心是求局部最优，但不定是全局最优。若想全局最优，必须证明。dp是通过一些状态来描述一些子问题，然后通过状态之间的转移来求解。般只要转移方程是正确的，答案必然是正确的。
    （2）动态规划本质上是穷举法，只是不重复计算罢了。结果是最优的。复杂度高。贪心算法不一定最优。复杂度一般较低。
    （3）贪心只选择当前最有利的，不考虑这步选择对以后的选择造成的影响，眼光短浅，不能看出全局最优；动规是通过较小规模的局部最优解一步步最终得出全局最优解。
    （4）从推导过程来看，动态规划是贪心的泛化，贪心是动态规划的特例
```

##### 分割数组为连续子序列 659
    ```
        此题比较难想
        题解算法，自己有一半思路，但是没有想到用两个hash表
        使用一个哈希表记录每个值重复次数，另一个哈希表记录子序列队尾值，即该子序列截止点为key
        使用贪心算法尽可能延长子序列长度
        每当一个子序列可以延长时，那么就尽可能延长它
        1 当nums[i]的次数为0时跳过
        2 当发现nums[i-1]是一个子序列结尾，并且nums[i]存在可用次数时，将子序列延长至nums[i]
          即把记录子序列的结尾改为subqueue[nums[i]]+=1，并减少nums[i]可用次数,同时去取消上一个延长前尾节点记录subQueue[nums[i-1]]-=1
          此规则主要针对重复项
        3 当重复记录哈希表中的key[nums[i],nums[i+1],nums[i+2]]均为>0，则新开一个队列  
    ```
##### 最长公共子序列 1143
    ```
        动态规划
        dp[i][j]来自{dp[i-1][j-1]+1,dp[i-1][j],dp[i][j-1]}中的最大者
        if text1的第i-1个字符与text2中第j-1个字符相同
         => dp[i][j]=dp[i-1][j-1]+1
        else
            dp[i][j]=max(dp[i-1][j],dp[i][j-1])
      
        注意，初始数组的长度 要比字符长度大1 因为第一个dp[0][0]==0，为基准线
        循环记得要与<=len，否则少一次判断，因为判断都是判断text1[i-1]==text2[j-1]
        想起来比较麻烦的话 画一个二维数组看一下状态转移图
        需要多练习，最长字串子数组类的问题 状态较多较复杂的 均为动态规划题型
        代码：
        func x(text1,text2 string)int {
            m, n := len(text1), len(text2)
            	memo := make([][]int, m+1)
            	for i := 0; i <= m; i++ {
            		memo[i] = make([]int, n+1)
            	}
            
            	for i := 1; i <= m; i++ {
            		for j := 1; j <= n; j++ {
            			if text1[i-1] == text2[j-1] {
            				memo[i][j] = 1 + memo[i-1][j-1]
            			} else {
            				memo[i][j] = max(memo[i][j-1], memo[i-1][j])
            			}
            		}
            	}
            	return memo[m][n] 
            }
    ```
##### 四数之和
    ```
        最后一层循环通过双指针优化，其余还是要三重循环，
        另外注意continue条件，比如第二层 一定是要j>i+1&&j-1==j
    ```
##### dp
    ```
    最优子结构 opt [n]=best_of(opt[n-1],opt[n-2]...)
    存储中间状态：opt[i]
    递推公式(状态转移方程)
        fib: opt[i]=opt[i-1]+opt[i-2]
        二维路径：opt[i][j]=opt[i+1][j]+opt[i][j+1](判断i，j是否是空地)
    数学归纳法思维
    分治一般是递归存储状态，dp一般是存储中间状态
    过滤最佳值
    已知问题规模为n的前提A，求解一个未知解B。（我们用An表示"问题规模为n的已知条件"）
    
    此时，如果把问题规模降到0，即已知A0，可以得到A0->B.
    
    1）如果从A0添加一个元素，得到A1的变化过程。即A0->A1; 进而有A1->A2; A2->A3; …… ; Ai->Ai+1. 这就是严格的归纳推理，也就是我们经常使用的数学归纳法；
    
    对于Ai+1，只需要它的上一个状态Ai即可完成整个推理过程（而不需要更前序的状态）。我们将这一模型称为马尔科夫模型。对应的推理过程叫做"贪心法"。
    
    2）然而，Ai与Ai+1往往不是互为充要条件，随着i的增加，有价值的前提信息越来越少，我们无法仅仅通过上一个状态得到下一个状态，因此可以采用如下方案：
    
    {A1->A2}; {A1, A2->A3}; {A1,A2,A3->A4};……; {A1,A2,...,Ai}->Ai+1. 这种方式就是第二数学归纳法。
    
    对于Ai+1需要前面的所有前序状态（或几个状态）才能完成推理过程。我们将这一模型称为高阶马尔科夫模型。对应的推理过程叫做"动态规划法"。
    dp step:
        1 定义子问题 define subproblem
        2 猜测子问题迭代关系 guess relationship
        3 确定dp列表推算方法 make sure dp list calculate method 
        4 自底向上 bottom-up 推算
        5 集合/返回结果
    ```
##### 不同路径 62
    ```
    因为只能向右或者向下：
    矩阵第一行都是1（只向右一条路线），矩阵第一列都是1（只向下一条路线）
    dp[i][j]=dp[i-1][j]+d[i][j-1]
    递推到dp[m-1][n-1]即可
    ```
##### 零钱兑换 322
    ```
    能兑换的钱数一定是减去硬币面值之后也可以兑换的钱数
    第一步先假设不能兑换，赋值为-1
    然后有能兑换的 就把能兑换的值赋给dp[i]
    后边还有能兑换的，假如比dp[i]小 ，就赋给dp[i]
    当第n块钱兑换时
        for i:=1;i<amount;i++{
            dp[i]=-1
            for x :=range coins{ 
                if dp[i-coin[x]]==-1{continue}
                if i-coin[x]<0{continue}
                count:=dp[i]=dp[i-coin[x]]+1
                if dp[i]==-1{
                    dp[i]=count //说明除了这个符合条件的之外，无法转化
                }
                if dp[i]>count{
                    dp[i]=count //说明之前可以转化的零钱数比这个硬币转化的数字要大
                }
            }
        }
        return dp[amount]即可
    better one :
    func CoinChangeT1(coins []int, amount int) int {
    	var dp =make([]int,amount+1)
    	dp[0]=0
    	for i:=1;i<=amount;i++{
    		dp[i]=-1
    		for _,k:=range coins{
    			// 硬币数大于amount或者上一个硬币无法转化的amount，跳过
    			if i-k<0||dp[i-k]==-1{
    				continue
    			}
    			//记录硬币可转化个数
    			count:=dp[i-k]+1
    			//如果非首次，也不是最小值则dp[i]重新赋值
    			if dp[i]!=-1||dp[i]>count{
    				dp[i]=count
    			}
    		}
    	}
    	return dp[amount]
    }
    ```
##### 三角形最小路径之和 120
    ```
        自底向上，从列表倒数第二行开始计算，每层的每个位置都是下层相邻两节点路径的最小值+当前位置数值
        如
        i:0,j:0       [1]
        i:1,j:0-1    [2,3]
        i:2,j:0-3   [4,3,1]
        i:3,j:0-4  [1,3,51,1]
          步骤即为：[i][j] = [i][j]   + min( [i+1][j] , [i+1][j+1] )
             当前位置路径和 = 当前元素值 + 正下方元素和斜下方元素中的较小值
        
        
        自顶向下：
        1 分治，当前值的最小路径和为下层元素中最小者+当前元素值 递推
        当到最底层 ，返回
            if level == len(triangle)-1 
            return triangle[level][index]
            // 最下层取当前值 即为当前最小路径（唯一路径）
        否则 递归 
            l:=（triangle，level+1,index）
            //下层左边值最小路径
            r:= (triangle,level+1,index+1)
            //下层右边值最小路径
            return min(l,r)+triangle[level][index]
            //取当前层最小路径
            
        2 和自底向上差不多，每层计算当前层的长度，即：
           triangle [i][j]=triangle[i-1][j-1]+triangle[i][j]||triangle[i-1][j]+triangle[i][j] 取小值，
           向下累加每个位置的最小路径，到最后一层，返回当前层计算结果中最小的值即可     
    ```
##### 翻转矩阵后得分 861 贪心
    ```
        1 先把矩阵每一行首位为0的全部变成1
        2 把每一列0的个数大于1的个数的 全部翻转
        这一步，可以直接计算列的0个数，1的个数就是矩阵行数减去0个数
        同时直接计算结果累加，不必进行实际翻转，官方题解666
        3 计算结果即可
    ```
##### 打家劫舍I&&II 198&&213 
    ```
        198：基础版打家劫舍 
        维护二维数组，每个房屋的状态都有偷窃和不偷窃
        对于第i个房屋，决定偷当前房屋/不偷 如果偷，那么i-1个房屋就不可以偷，要取未偷i-1房屋的值
        如果不偷 到第i房间的值就可以来源于偷盗上一个房间的值
        用0，1 表示偷或者不偷当前房间，那么第i个房间的最大值为
        dp[i][0]=max(dp[i-1][1],dp[i-1][0]) （当前不偷，上一间偷或不偷的较大者，因为可以跳两个都不偷）
        dp[i][1]=dp[i-1][0]+nums[i] （当前如果偷，那么一定上一件房屋是不可以偷的，所以只取0状态的i-1加上当前房屋金额）
        最终返回max（dp[len-1][0]，dp[len-1][1]）
        简化版：
            因为第i间房屋的状态基本都由上一个i-1或者上两个i-2房屋的偷盗情况演化（dp[len-1]/dp[len-2]+nums[i]）
            并不会跳过三间房屋不偷，那必然不是最大值，所以
            dp[i]=max（dp[i-1]，dp[i-2]+nums[i])
            i-2包含了i-3和i-4等等。。逐层递推
            所以可以通过一维dp状态数组定义解决
            初始值dp[0]=nums[0] dp[1]=max(nums[1],nums[0])
        213 环形打家劫舍
        相比较基础版，此问题需要考虑首位相交，因为相连两个都偷要报警，那么只能取其一即可（两个都不取也包含在内）
        所以拆分成两个数组，第一个dp求出nums[0:len-2]范围的最大值，第二个dp求出nums[1:len-1]范围的最大值
        最终点比较两个dp求出的最大值取之即可
        dp1[0]=nums[0]
        dp1[1]=max(nums[0],nums[1])
        for i:=2;i<len-2;{}
        dp2[1]=nums[1]
        dp2[2]=max(nums[1],nums[2])
        for i:=3;i<len-1;{}
        max(dp1[-1],dp2[-1])
    ```
##### dota参议院 649
    ```
        运用两个队列，每次封掉对手的人重新进入队列，进行下一轮的较量
        队列中记录的为每个参议员的索引，索引小者将战胜索引大者
        r:=[]int{}
        d:=[]int{}
        for i:=range senate{
            if R{
            r=append(r,i)
            }
            if D{
            d=append(d,i)
            }
        }
        for r.len>0&&d.len>0{
            r1=r[0]
            d1=d[0]
            if r1<d1{ //r索引小，所以r先手
                r=append r,r1+len(senate) //加上当前索引，保证相对顺序 
            }else{
                d=append d,d1+len(senate) //加上当前索引
            }
            r=r[1:]
            d=d[1:]
        }
        //当最终某个队列为空，那么非空队列阵营的将取得胜利
        if r.length>0{
        }return r win 
        else {return d win}
    ```
##### 摆动序列 376
    ```
        递推要分开两种情况，一种是1，3，2，4这种的上升摆动序列 另一种是 4，1，3，2这样的下降摆动序列 
        当走到第i个位置时，要判断第i-1 到第i 是上升还是下降
        1 如果i>i-1 是上升 那么从下降序列转换过来将是必然合法的 所以down+=1,上升序列转换非法，
            此时，i的上升序列（下降转化/延长上升）就变为max（up，down+1）i的下降序列长度保持不变
            1 上升序列延续 up i = up i-1
            2 下降序列转化 up i = down i-1 +1
            取大者 up i = max（up i-1 ，down i-1 +1）
        2 而当i<i-1 那么i-1 到i 即为下降序列，所以两种可能 ：
            1 下降序列延续 即 down i= down i-1
            2 由上升序列转化 down i= up i-1 +1 
            取大者 down i = max（down i-1 ， up i-1 +1）
        当i==i-1时，延长up 和down 
        
        数组版代码比较容易看出思路：
            func wiggleMaxLength(nums []int) int {
            	length:=len(nums)
            	if length<2{
            		return length
            	}
            	up,down :=make([]int,len(nums)),make([]int,len(nums))
            	up[0]=1
            	down[0]=1
            	for i:=1;i<len(nums);i++{
            		if nums[i]>nums[i-1]{
            			up[i]=max376(up[i-1],down[i-1]+1)
            			down[i]=down[i-1]
            		}else if nums[i]<nums[i-1]{
            			down[i]=max376(down[i-1],up[i-1]+1)
            			up[i]=up[i-1]
            		}else{
            			up[i]=up[i-1]
            			down[i]=down[i-1]
            		}
            	}
            	return max376(up[length-1],down[length-1])
            }
        
    ```