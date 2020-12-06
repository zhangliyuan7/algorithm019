package Week_06

import "fmt"

type Node struct{
	Val int
	L *Node
	R *Node
}

func bfs (root *Node){
	var q []*Node
	q=append(q, root)
	for len(q)!=0{
		root=q[0]
		q=q[1:]
		visit(root)
		if root.L!=nil{
			q=append(q, root.L)
		}
		if root.R!=nil{
			q=append(q,root.R)
		}
	}
}

func visit(root *Node){
	fmt.Println(root.Val)
}