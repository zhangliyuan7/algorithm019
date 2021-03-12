package Repeat_01

import "strings"

// https://leetcode-cn.com/problems/verify-preorder-serialization-of-a-binary-tree/solution/pai-an-jiao-jue-de-liang-chong-jie-fa-zh-66nt/

// 331
// 树 图 入度 出度
// 就是图的有向链接数目，节点的入度即为指向该节点的路径，出度就是该节点指向其他节点的路径
// 每个# 有1个入度 0个出度
// 每个!# 有2个出度 一个入度
// && 树的入度 == 出度
// Google: 顶点的度是指和该顶点相连的边的条数。 特别是对于有向图来说，顶点的出边条数称为该顶点的出度，顶点的入边条数称为该顶点的入度
func isValidSerialization(preorder string) bool {
	nodes:=strings.Split(preorder,",")
	//根节点入度为0 出度为2
	//首次进入循环 减去一个入度 所以root先要提前给一个入度1 为了减为0
	degree:=1
	for i:=range nodes{
		//减入度
		degree-=1
		if degree<0{
			return false
		}
		// 非空节点 加出度2
		if nodes[i]!="#"{
			degree+=2
		}
	}
	return degree==0
}
// 中左右 ，如果遇到叶子节点 中非空 左右为空 那么至少该子树是符合条件的
// 所以折叠树状态 ，令该节点为空"#"
// 不断坍塌 最终每个子节点都会折到root上
// 如果能完整的坍塌为空 即倒序坍塌状态成功
// 那么即是一个合格的树 ，那么通过中序倒序折叠过程成功折叠，即说明这是正确的中序遍历
// 9,3,4,#,# => 9,3,#，继续
// 9,3,#,1,#,# => 9,3,#,# => 9,# ，继续
// 9,#2,#,6,#,# => 9,#,2,#,# => 9,#,# => #，结束
// 过程很像向内坍塌的黑洞
func isValidSerializationII(preorder string) bool {
	 nodes:=strings.Split(preorder,",")
	 stack:=[]string{}
	 for _,node:=range nodes{
	 	stack=append(stack,node)
	 	// 不断从子树向上 逆序压缩，（两个子为空 父为非空的部分 向上折叠，三个都是空的话非法）
	 	for len(stack)>=3&&stack[len(stack)-1]=="#"&&stack[len(stack)-2]=="#"&&stack[len(stack)-3]!="#"{
	 		stack=stack[:len(stack)-3]
	 		stack=append(stack,"#")
		}
	 }
	 //坍塌到root了，仅剩下#了
	 return len(stack)==1&&stack[0]=="#"
}

