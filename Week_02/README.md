学习笔记
##### hashFunc Map=HashFunc + LinkedList + List
      最好情况查询O(1),插入O(1),最差情况插入O(n),查找O(n),取决于数组长度及哈希函数的碰撞率

##### relative sort array
       ```
        1 > 自定义排序函数，排序
        自定义sort解决，注意理解题目，是相对顺序不变，并不是把所有的arr2的元素全放在前面
       ```
       
##### 删除k位元素，使结果最小 remove_k_digits.
    ```
        important: 
            1 > 不要用[]int数组操作10的n次方相乘相加方式转换为数字再转string的方式（轻松溢出），使用byte数组直接转string
            2 > rune 与 byte 不能直接相互转换，所以i,c:=range string的时候 
                c.(rune) string[i].([]byte)，操作[]byte 时 要通过string[i]去做处理
        用栈解决问题，从前到后删除k个数的过程，使数非严格单调递增，比较栈顶元素和新元素，如果不符合单调递减，弹出栈顶元素，循环比较，直到k==0或者符合单调递减
        最终将结果总头到尾从栈中取出（从头部取出，概念其实是双端队列，实现是数组而已）
        如果最终实现单调递减，但是删除个数不足k，为n ，那么去掉尾部n个数字
        最终结果左边是"0"的||result=="" 特殊处理一下，去掉左边0||输出"0"
    ```

##### 堆
    
    // 实现官方接口heap，默认小顶堆
    Push
    push方法，其官方注释要求实现的功能写的很简单，事实上也确实很简单。
    
    将x添加在数组最后即可，所以实现代码如下：
    
    func (h *IntHeap) Push(x interface{}) {
    	// Push and Pop use pointer receivers because they modify the slice's length,
    	// not just its contents.
    	*h = append(*h, x.(int))
    }
  
    你可能会疑惑，如果只是实现这样得Push函数，并没有保证堆的性质啊，你只是把这个数放到了数组最后。实际上这是官方为了避免我们写太多代码而做的设计，我们在push时实际上是调用heap.Push(h, 3) ,这是在heap.go中的一个函数，其具体实现为：
    
    func Push(h Interface, x interface{}) {
    	h.Push(x)
    	up(h, h.Len()-1)
    }
    func up(h Interface, j int) {
    	for {
    		i := (j - 1) / 2 // parent
    		if i == j || !h.Less(j, i) {
    			break
    		}
    		h.Swap(i, j)
    		j = i
    	}
    }
  
    上浮的操作由官方函数up完成。
    
    Pop
    func (h *IntHeap) Pop() interface{} {
    	old := *h
    	n := len(old)
    	x := old[n-1]
    	*h = old[0 : n-1]
    	return x
    }
   
    由官方注释可知，我们只需要将切片变为其[0:n-1]即可，那么如何保证获取的是最小/最大值呢，查看heap.go中Pop函数可知：
    
    func Pop(h Interface) interface{} {
    	n := h.Len() - 1
    	h.Swap(0, n)
    	down(h, 0, n)
    	return h.Pop()
    }
    func down(h Interface, i0, n int) bool {
    	i := i0
    	for {
    		j1 := 2*i + 1
    		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
    			break
    		}
    		j := j1 // left child
    		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
    			j = j2 // = 2*i + 2  // right child
    		}
    		if !h.Less(j, i) {
    			break
    		}
    		h.Swap(i, j)
    		i = j
    	}
    	return i > i0
    }
    // 
    包含多种实现，常见面试实现为二叉树或n叉树堆，查询O(1)，其他O(log(n))，时间简单，但工程上应用少，性能一般
    比较牛叉的是斐波那契堆和严格的斐波那契堆，除删除外，均可达到O(1),工程上较多,实现较复杂
    |Operation|find-max|delete-max|insert|increase-key|meld|
    |---|---|---|---|---|---| 
    | Binary | O(1) | O(log n) | O(log n) | O(log n) | O(n) |
    | Fibonacci | O(1) | O(log n) | O(1) | O(1) | O(1) |
    | Strict Fibonacci | O(1) | O(log n) | O(1) | O(1)| O(1) |
  
    (大顶堆为例，小顶堆第二条相反而已)   
    特性BHeap ：
    ```
      1 > 完全二叉树（除了叶子节点，其他节点是满的）
      2 > 父节点永远比子节点大
      3 > 父节点(索引为i)的左右子节点索引分别为 左：i*2+1，右：i*2+2
      4 > 子节点(索引为i)的父节点为 (i-1)/2
      5 > 插入：新元素插入堆最尾部，不断与其父节点比较，进行交替的过程，父节点大于新元素为止
          过程函数 HeapifyUp
      6 > 弹出/删除：堆顶元素与堆最尾部元素互换位置，然后列表size--，删除最尾部元素（最大值）
          堆顶此时是最小值，慢慢下沉.与其左右子节点比较，大者上浮，交换位置，迭代，直到子节点索引号>size即可停止（没有子节点）
          过程函数 HeapifyDown
   
    ```
   
##### 异位词anagram处理
     ```
     异位词判断
     1 > 用固定[26]int{}数组做每个字母的asc码值哈希到列表对应下标位置
     比较[26]int{}是否相同即可 (go中列表可以直接比较，且固定长度列表可以作为map key)
     2 > 排序字符串，比较字符串是否相同即可
     ```
     
##### 数组交集
     ```
     1 > 先固定一个数组
     hashMap 记录其中每个数字的个数 number:count
     循环另一个数组，只要数字存在于map,加入到交集结果，同时count--，当count==0，删除该key
     2 > 排序两个数组，双指针分别代表两个数组的索引号
     nums1[i]==nums2[j] => i++&&j++ append->result
     nums1[i]>nums2[j] => j++;continue
     nums1[i]<nums2[j] => i++;continue     
     ```
     
##### 滑动窗口最大值
    ```
    1 > kernel ：单调队列 （由大到小）
        i<k-1时 一直添加元素，最多k-2个，到k-1个就要求第一个max了
        维护队列，窗口大小k以内，直接添加到队列，队列要求元素非严格递减，后续添加元素
        push:入队需要判断元素大小，踢出加入新元素后，元素顺序中不符合单调递减的元素
        for example:1 > queue A ; if A=[4,2,1]; new element 3; A.push(3) => A=[4，3]
                    2 > queue B ; if B=[5,4,2]; new element 4; B.push(4) => B=[5,4,4]
                    3 > queue C ; if C=[3,2,1]; new element 4; C.push(4) => C=[4] 
        
        pop:出窗口值，比较出队值是否与队列中最大值，即队列首个元素相同，如果相同，队列清掉当前最大值
        for exmaple:1 > queue A; if A=[4,2,1] ; pop element 4; A.pop(4); if A[0]==4 { A=A[1:]} => A=[2,1]
                    2 > queue B; if B=[4,2,1] ; pop element 3; B.pop(3); if B[0]!=3 {B not change } => B=[4,2,1]
        slidingwindown max :就是队列中的队首元素 queue[0]
                          
    2 > 半暴力 
        循环数组，暴力查找(可以用两个指针收敛式更快查找)窗口索引内的最大值，加入到结果集
    3 > 大顶堆  todo
    ```

##### 去掉最外层括号
     ```
     kernel 栈
     for _,c:=range S{
        if string(c)=="("{
            if len(s)!=0{ builder.Write }
        }
        if string(c)==")"{
            stack=stack[:len-1]
            if len(s)!=0{ builder.Write } 
        }
     }
     strings.builder.WriteString(string(c))
     栈空时的第一个入栈元素删除/不记录
     出栈后，栈空的最后一个出栈元素删除/不记录
     记录中间经过的其他string(c)即可 go中使用strings.builder拼接也很快
     ```
     
##### 树,二叉，二叉搜索树
    ```
    链表为树的特殊形态(单枝)，树为图的特殊形态(无环)
    树遍历按照遍历特性分为两种遍历类型
    1 > 广度优先BFS，深度优先DFS
    2 > 前序-根左右  中序-左根右 后序-左右根
    二叉搜索树 左子树永远小于父节点，右子树永远大于父节点 其操作
        1 > 插入 按照大小顺序不断比较，找到树中合适位置插入
        2 > 删除 找到删除节点的右子树中最靠近删除节点的值，进行补位
        （右子树最小节点/左子树最大节点）
        3 > 查找 类似二分查找过程
    树的操作近乎都是O(log n )，极端情况可能会退化为O(n)(如：链表形态)
    递归操作 如果合理使用缓存或者其他形式优化 速度并不慢，并不是说递归一定就低效，但是递归会开辟很多栈空间
        如果层级较多，开辟的空间较大，可能会溢出
    先序：用栈存好已经读取过的相对根，只要左子节点不为空，就一直读取值，然后入栈，直到左叶子节点为空，开始出栈，出栈后重复循环判断右子树的过程
    中序与先序入栈顺序相同,从上到下一路向左入栈，只不过读取值在出栈之后，右子树处理与先序同
    后序 todo
    ```
##### FizzBuzz
    ```
    题目扩展为多个因子时，用字符串拼接+hash散列比较优雅，暴力法代码行数太多，判断太多
    ```
##### 图
    ```
    表示形式分为邻接矩阵和邻接表两种形式
    概念包含点/边/度 
        点表示端点
        边表示是否连接和路径
        度表示一个点与多少个其他点连接(即该点是多少线段的端点)
    邻接矩阵中 横纵坐标皆为图中点一维坐标，矩阵值为点与点之间的（有无连接/路径长度/权重）
    无向图无权图 则仅有连接关系(对称(0-1)矩阵)
    无向图有权图 有路径长短关系(对称(0-weight)矩阵)
    有向有权图 有路径长短及方向关系(非对称(0-weight)矩阵)
    有向无权图 只有方向关系(非对称1-0矩阵)
    todo:
    BFS,DFS
    Dijkstra 优先队列+BFS
    ```
##### 奇偶链表
    ```
    拆成俩链表，最后合并，前进过程一次走两步，简化写
    ```
##### have to do
    ```
        堆:TOP-K问题 滑动窗口大顶堆 堆排序 堆实现
        图:Dijkstra(最短路径问题:
            priorityQ
            BFS
            visted
            入队权值相加) 
          BFS
          DFS
        
        树:前中后序，递归与非递归，二叉搜索树实现
        队列:priority Queue实现
    ```
