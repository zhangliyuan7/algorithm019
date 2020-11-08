package practice

import (
	"container/list"
	"fmt"
)

//DFS/BFS 知识点分享
//深度优先DFS
//从顶点V，沿着一条路一直走到底，再从这条路的尽头的节点退回到上一个节点，再从另一条路走到底，不断递归重复，直到所有定点遍历完成，不撞南墙不回头
//
//遍历：
//1.从根节点1开始，像临界点有2，3，4，顺序为1—>2—>5—>9
//2.9后面没有别的节点，往上退，退回5—>2因为都没有其余节点，则从3开始，3—>6—>10
//3.到10后回退到3，3有右子节点，则--->7
//4.4---->8
//这个是完整的深度优先遍历
//实现方式(实现方式：前中后序遍历)
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

//栈
type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	list := list.New()
	return &Stack{list}
}

func (stack *Stack) Push(value interface{}) {
	stack.list.PushBack(value)
}

func (stack *Stack) Pop() interface{} {
	if e := stack.list.Back(); e != nil {
		stack.list.Remove(e)
		return e.Value
	}

	return nil
}

func (stack *Stack) Len() int {
	return stack.list.Len()
}

func (stack *Stack) Empty() bool {
	return stack.Len() == 0
}
func (stack *Stack) Top() interface{} {
	if e := stack.list.Back(); e != nil {
		return e.Value
	}
	return nil
}


//前序遍历.  利用栈的特性，先将右子树进行压栈，再将左子树进行压栈
func (root *Node) PreTravesal() {
	if root == nil {
		return
	}
	s := NewStack()
	s.Push(root)

	for !s.Empty() {
		cur := s.Pop().(*Node)
		fmt.Println(cur.Val)

		if cur.Right != nil {
			s.Push(cur.Right)
		}
		if cur.Left != nil {
			s.Push(cur.Left)
		}
	}
}

//中序遍历
func (root *Node) InTravesal() {
	if root == nil {
		return
	}

	s := NewStack()
	cur := root
	for {
		for cur != nil {
			s.Push(cur)
			cur = cur.Left
		}

		if s.Empty() {
			break
		}

		cur = s.Pop().(*Node)
		fmt.Println(cur.Val)
		cur = cur.Right
	}
}

//后序遍历：则按照先进后出，将根节点先进行压栈，然后右节点，左节点
func (root *Node) PostTravesal() {
	if root == nil {
		return
	}

	s := NewStack()
	out := NewStack()
	s.Push(root)

	for !s.Empty() {
		cur := s.Pop().(*Node)
		out.Push(cur)

		if cur.Right != nil {
			s.Push(cur.Right)
		}
		if cur.Left != nil {
			s.Push(cur.Left)
		}
	}

	for !out.Empty() {
		cur := out.Pop().(*Node)
		fmt.Println(cur.Val)
	}
}

//BFS层次遍历
//
//核心思想：每次出队一个元素，就将该元素的孩子节点加入队列中，直至队列中元素个数为0时，出队的顺序就是该二叉树的层次遍历结果
//1.根节点入队列，出队列时，将孩子节点即B，C入队列
//2.B节点出队列时，将孩子节点D入队列
//3.C节点出队列时，孩子节点F，G入队列
//4，以此类推
//广度优先：数组实现
func levelOrder(root *Node) [][]int {
	result = make([][]int, 0)
	if root == nil {
		return result
	}
	dfsHelper(root, 0)
	return result
}
func dfsHelper(node *Node, level int) {
	if node == nil {
		return
	}
	if len(result) < level+1 {
		result = append(result, make([]int, 0))
	}
	result[level] = append(result[level], node.Val)
	dfsHelper(node.Left, level+1)
	dfsHelper(node.Right, level+1)
}

//DFS/BFS 知识点分享
//深度优先DFS
//从顶点V，沿着一条路一直走到底，再从这条路的尽头的节点退回到上一个节点，再从另一条路走到底，不断递归重复，直到所有定点遍历完成，不撞南墙不回头
//
//遍历：
//1.从根节点1开始，像临界点有2，3，4，顺序为1—>2—>5—>9
//2.9后面没有别的节点，往上退，退回5—>2因为都没有其余节点，则从3开始，3—>6—>10
//3.到10后回退到3，3有右子节点，则--->7
//4.4---->8
//这个是完整的深度优先遍历
//实现方式(实现方式：前中后序遍历)
//前序遍历.  利用栈的特性，先将右子树进行压栈，再将左子树进行压栈
func (root *Node) PreTravesal1() {
	if root == nil {
		return
	}

	s := NewStack()
	s.Push(root)

	for !s.Empty() {
		cur := s.Pop().(*Node)
		fmt.Println(cur.Val)

		if cur.Right != nil {
			s.Push(cur.Right)
		}
		if cur.Left != nil {
			s.Push(cur.Left)
		}
	}
}

//中序遍历
func (root *Node) InTravesal2() {
	if root == nil {
		return
	}

	s := NewStack()
	cur := root
	for {
		for cur != nil {
			s.Push(cur)
			cur = cur.Left
		}

		if s.Empty() {
			break
		}

		cur = s.Pop().(*Node)
		fmt.Println(cur.Val)
		cur = cur.Right
	}
}

//后序遍历：则按照先进后出，将根节点先进行压栈，然后右节点，左节点
func (root *Node) PostTravesal2() {
	if root == nil {
		return
	}

	s := NewStack()
	out := NewStack()
	s.Push(root)

	for !s.Empty() {
		cur := s.Pop().(*Node)
		out.Push(cur)

		if cur.Right != nil {
			s.Push(cur.Right)
		}
		if cur.Left != nil {
			s.Push(cur.Left)
		}
	}

	for !out.Empty() {
		cur := out.Pop().(*Node)
		fmt.Println(cur.Val)
	}
}

//BFS层次遍历
//
//核心思想：每次出队一个元素，就将该元素的孩子节点加入队列中，直至队列中元素个数为0时，出队的顺序就是该二叉树的层次遍历结果
//1.根节点入队列，出队列时，将孩子节点即B，C入队列
//2.B节点出队列时，将孩子节点D入队列
//3.C节点出队列时，孩子节点F，G入队列
//4，以此类推
//广度优先：数组实现
var result [][]int

func levelOrder2(root *Node) [][]int {
	result = make([][]int, 0)
	if root == nil {
		return result
	}
	dfsHelper2(root, 0)
	return result
}
func dfsHelper2(node *Node, level int) {
	if node == nil {
		return
	}
	if len(result) < level+1 {
		result = append(result, make([]int, 0))
	}
	result[level] = append(result[level], node.Val)
	dfsHelper(node.Left, level+1)
	dfsHelper(node.Right, level+1)
}
