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