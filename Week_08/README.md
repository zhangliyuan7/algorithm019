学习笔记

##### 位运算
    ```
        >>  0110 -> 0011
        <<  0110 -> 1100
        n&n-1 清除最后一位的1 
        n&(-n)   ==  x & (~x + 1) 得到仅仅包含n的最后一位1标示的数字
        比如 12(1100) =>
            (12)&(-12)=>(1100)&(0100) //反码 按位取反，末位+1 
            => 0100
            => 8
             （12）&（-12）= 8
             ① 这个结果只有一位值是1， 其他位均是0 
             ② 这个值的末位0的个数与原值保持一致
        
        ^异或  相同为0 不同为1 
        x^0 = x
        x^1s = ~x  1s=～0  全零
        x^(~x) = 1s
        x^x = 0
        c=a^b => a^c =b b^c=a
        a^b^c=a^(b^c)=(a^b)^c
        x&~x = 0   
        
        x最右边的n位清零 x&(~0<<n)
        x的第n个位置 (x>>n)&1
        x的第n位的幂值 x&&(1<<n)
        将第n位置为1  x|(1<<n)
        第n位置换为0  x&(~(1<<n))
        x的最高位至第n位清零  x&((1<<n)-1)
        判断奇偶
            x&1==1 奇数
            x&1==0 偶数
        整除 
            x/2 == x>>1
        清零最低位的1
            x&(x-1)
        得到最低位的1
            x&-x
        
    ```
##### 2的幂次 231
    ```
        二进制有且仅有一个1 即为2的幂次
        二进制位消末位1 n&(n-1) 的次数即个数
        if n!=0 {
            count++
            if count>1{
                return false 
            }
            n=n&(n-1)
        }
        简化：
            if n<0{return false}
            return n&(n-1)==0
    ```
##### 1出现的次数 191 
    ```
        n&(n-1) 次数,即1的个数
    ```
##### 颠倒二进制 190
    ```
        取最后一位，倒序移位相加
        power:=31 (uint32)
        for n!=0{
            last:=n&1
            result+= last<<power
            n= n>>1
            power-=1
        }
    ```
##### bloom filter 布隆过滤器
    ```
        类似一个极速缓存，拦截在db或者缓存外边，做一个高效的粗略的过滤拦截
        对请求对象进行二进制存储，当获取到二进制位存在任何一个0 ，则可以确定该数据不在后边的db或者存储空间里，内容不存在
        当不存在0的时候，则有可能存在于后端存储，就可以允许其请求后端存储进行数据获取尝试
        优点：简单，快速，高效
        缺点：删除困难，因为有重复利用的二进制位，有一定误判率（可以确定不存在，但是不能确定一定存在）
    ```
##### LRU cache 
    ```
        最近最少使用元素
        通常由 hash map+double linkedlist实现
        查询O(1)-hash map的部分，插入和删除也是O(1) 链表的部分
        当数据进入，不断往链表顶部写入
        链表空间满时：新数据存储空间不足时，踢出链表中最尾部的元素
        当访问了存在于链表中的旧元素时，将旧元素翻新（删除该元素，重新插入回链表顶部）
        
        类似页面置换算法（LFU 最近访问频次最少元素）
    ```
##### 146 LRU缓存
    ```
        movetohead操作和movethetail操作 千万注意 将前后链接好，因为是双向链表，一定要两个方向都链接一次
        包括新增节点，一定要将head.next 和当前新增的node 和head链接好
        head.next.pre=node
        node.next=head.next
        head.next=node
        node.pre=head
        删除时 也是同理 另外，要保留两个标志节点，head和tail ，不要替换掉
        type LRUCache struct {
        	cap  int
        	length int
        	hashMap map[int]*dlinklist
        	head *dlinklist
        	tail *dlinklist
        }
        func Constructor(capacity int) LRUCache {
        	head:=dlinklist{pre:nil,next:nil}
        	tail:=dlinklist{pre:nil,next:nil}
        	head.next=&tail
        	tail.pre=&head
        	return LRUCache{
        		cap:capacity,
        		length:0,
        		hashMap:make(map[int]*dlinklist),
        		head:&head,
        		tail:&tail,
        	}
        }
        
        type dlinklist struct{
        	pre *dlinklist
        	next  *dlinklist
        	key int
        	value int
        }
        func (this *LRUCache) Get(key int) int {
        	if v,ok:=this.hashMap[key];ok{
        		this.moveToHead(v)
        		return v.value
        	}
        	return -1
        }
        
        func (this *LRUCache) Put(key int, value int)  {
        	if v,ok:=this.hashMap[key];ok{
        		v.value=value
        		this.moveToHead(v)
        		return
        	}
        	if this.cap==this.length{
        		this.removeTail()
        	}
        	node:=dlinklist{
        		pre:this.head,
        		next:this.head.next,
        		key:key,
        		value:value,
        	}
        	this.head.next.pre=&node
        	this.head.next=&node
        	this.hashMap[key]=&node
        	this.length++
        }
        func (this *LRUCache)moveToHead(d *dlinklist){
        	this.removeNode(d)
        	d.next=this.head.next
        	this.head.next.pre=d
        	d.pre=this.head
        	this.head.next=d
        }
        
        func (this *LRUCache)removeTail(){
        	node:=this.tail.pre
        	this.removeNode(node)
        	delete(this.hashMap,node.key)
        	this.length--
        }
        
        func (this *LRUCache)removeNode(d *dlinklist){
        	d.pre.next=d.next
        	d.next.pre=d.pre
        }

    ```
##### 选择排序 n*n
```
    遍历剩组获取最小值，放到数组前面
    迭代剩余部分找到当前部分最小值，放到已安放数据的后边的位置，迭代
```
##### 插入排序 n*n
```
    遍历数组，将每次读取到的值，放在前面n个遍历过后的值的合适位置，不断迭代
```    
##### 冒泡排序 n*n
```
    不断两两比较，大的不断上浮，每次循环最大的都会浮到最右边
```
##### 快速排序 nlogn
```
    找到一个标杆，通过标杆不断拆分列表为 小于标杆的列表段，大于标杆的列表段
    不断递归排序即可
    分而治之
```
##### 归并排序 nlogn
```
    分而治之，拆分递归排序，最终合并有序列表
```
##### 堆排序 nlogn
```
    入堆和出堆 logn 
        heapifyUp（与father比较上浮）
        HeapifyDown（将尾元素提到最前面，不断下坠）
        注意如果是奇数个children,对于最后一个father*2+1,下坠时要记得最终处理一下
    father ：(i-1) >> 1
    children : i*2+1,i*2+2
    注意处理heapifydown时候的奇数 AGAIN
```
##### 螺旋矩阵 54
```
    采用读取一行后删除该行，然后将矩阵向左旋转90度，成为新矩阵，继续读取一行，继续删除，反复循环，直至矩阵为空
    翻转矩阵时注意，新矩阵行数=之前矩阵的列数，列数=新矩阵行数
    初始化新矩阵：
        if len(matrix)==0{
    		return [][]int{}
    	}
    	var matrixNew =make([][]int,len(matrix[0]))
    	for i:=range matrixNew{
    		matrixNew[i]=make([]int, len(matrix))
    	}
    翻转操作：
        n:=len(matrixNew)-1
    	newrow:=n
    	for n>=0{
    		for old_row:=0;old_row<len(matrix);old_row++ {
    			matrixNew[newrow-n][old_row] =matrix[old_row][n]
    		}
    		n--
    	}
    	return matrixNew
```
##### 区间合并 56
```
    先以左边界做排序，令左边最左的在前面
    用left，right做当前范围的标示符
    
    循环数组，当新访问元素左边界大于等于当前right时，right=新元素右边界与当前right的较大者
    当不符合上述条件时，将当前left，right添加到结果集，并将left=新元素左边界，right=新元素右边界，继续迭代
    当循环结束后，最后将left和right添加到结果集中，否则少一次结果
    部分代码如下：
    ...
    left := intervals[0][0]
	right := intervals[0][1]
	var r [][]int
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= right {
			right = max(intervals[i][1], right)
			continue
		}
		r = append(r, []int{left, right})
		left = intervals[i][0]
		right = intervals[i][1]
	}
	r=append(r,[]int{left,right})
    ...
```
##### 翻转对 493
```
    分治思想
    先写返回条件 len(nums)<=1 return 0 
    将当前数组拆分成两个单独的数组，left right
    分别递归计算单独数组包含的个数
    count=reversePair(left)+reversePair(right)
    此时的left与right已经做好了排序（递归函数统计完个数，最后归并排序其数组）
    类似归并合并的方式，比较两个数组合并时产生的count个数
    for l<left&&r<right{
        因为有序,当left的大于right*2时，left中比当前元素大的都比这个right*2大，所以count+=left.len()-l
        if left[l]>right[r]*2{
            count+=left.len()-l
            r++
        }else{
            l++
        }
    }
    最后做一次归并排序
    在原数组上排序，原地摆尾
    n:=len(nums)
    p1,p2:=0,0 
    for i:=0;i<n;i++{
        //这个if 的条件需要注意顺序，p2==len(right)时，即right先比left读完
        所以此时只添加left的值，注意"或"条件的左右两边顺序，一定要先判断p2，防止最后一个值和left比较而走else
        if p1<len(left)&&(p2==len(right)||left(p1)<right[p2])){ 
            nums[i]=left[p1]
            p1++
        }else{
            nums[i]=right[p2]
            p2++
        }
    }
    return count 
```
##### 