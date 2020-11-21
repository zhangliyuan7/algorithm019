package w3

// NO.236
var lowest *Node
//从上到下找到最后一个同时满足两个节点都是其子节点的节点为止
//too slow，submit timeout ；smile crying
func lowestCommonAncestor(root, p, q *Node) *Node {
	var searchq bool
	var searchp bool
	var queue []*Node
	queue=append(queue,root)
	for len(queue)!=0{
		last:=queue[0]
		queue=queue[1:]
		if last==nil{
			continue
		}
		if last==q{
			searchq=true
		}
		if last==p{
			searchp=true
		}
		if searchq&&searchp{
			lowest=root
			break
		}
		queue=append(queue,last.left)
		queue=append(queue,last.right)
	}
	lowestCommonAncestor(root.left,p,q)
	lowestCommonAncestor(root.right,p,q)
	return lowest
}

//record every node ->(his father)
func lowestCommonAncestor1(root, p, q *Node) *Node {
	var fathers = make(map[*Node]*Node)
	var visited = make(map[*Node]bool)
	var queue []*Node
	queue=append(queue,root)
	for len(queue)!=0{
		root=queue[0]
		queue=queue[1:]
		if root!=nil {
			if root.left != nil {
				fathers[root.left] = root
				queue=append(queue,root.left)
			}
			if root.right != nil {
				fathers[root.right] = root
				queue=append(queue,root.right)
			}
		}
	}
	for p!=nil{
		visited[p]=true
		p=fathers[p]
	}
	for q!=nil{
		if _,ok:=visited[q];ok{
			return q
		}
		q=fathers[q]
	}
	return nil
}

func lowestCommonAncestorB(root, p, q *Node) *Node {
	var visited = make(map[*Node]bool)
	var fathers = make(map[*Node]*Node)
	var s []*Node
	for root!=nil||len(s)!=0{
		if root!=nil{
			s=append(s,root)
			if root.left!=nil{
				fathers[root.left]=root
			}
			root=root.left
			continue
		}
		if len(s)!=0{
			root=s[len(s)-1]
			if root.right!=nil{
				fathers[root.right]=root
			}
			s=s[:len(s)-1]
			root=root.right
		}
	}
	for p!=nil{
		visited[p]=true
		p=fathers[p]
	}
	for q!=nil{
		if visited[q]==true{
			return q
		}
		q=fathers[q]
	}
	return nil
}









