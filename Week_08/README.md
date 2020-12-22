学习笔记

##### 位运算
    ```
        >>  0110 -> 0011
        <<  0110 -> 1100
        n&n-1 清除最后一位的1 
        n&(-n)   ==  x & (~x + 1)
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
    