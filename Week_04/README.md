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
            root.visited
            root.add(visited)
            dfs(root.left)
            dfs(root.right)
        }
    }
    func dfs_with_stack(){
        var s stack 
        var visited
        if visited||root==nil{return}
        for {
            if root!=nil{
                root.visited
                s.add(root.left)
                root=root.left
                continue
            }
            if s!=nil{
                root=s.pop
                s=s[:len-1]
                root=s.right
            }
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