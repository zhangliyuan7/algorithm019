package Week_06


func dfs(root *Node){
	var s []*Node
	for len(s)!=0||root!=nil{
		if root!=nil{
			visit(root)
			for root.L!=nil{
				s=append(s,root.L)
			}
		}
		if len(s)!=0{
			root=s[len(s)-1]
			s=s[:len(s)-1]
			root=root.R
		}
	}
}
func xianxu(root *Node){
	var s []*Node
	for root!=nil|| len(s)!=0{
		for root!=nil{
			visit(root)
			s=append(s,root)
			root=root.L
		}
		if len(s)!=0{
			root=s[len(s)-1]
			s=s[:len(s)-1]
			root=root.R
		}
	}
}
func zhongxu(root *Node){
	var s []*Node
	s=append(s,root)
	for root!=nil|| len(s)!=0{
		for root!=nil{
			s=append(s,root)
			root=root.L
		}
		if len(s)!=0{
			root=s[len(s)-1]
			visit(root)
			s=s[:len(s)-1]
			root=root.R
		}
	}
}
