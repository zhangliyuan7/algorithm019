学习笔记

##### 86 分隔链表
```golang 
    两个哨兵，其他与合并有序链表相类似，注意斩断拼接的节点与其next 即可
```
##### 300 最长上升子序列
```
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
```
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