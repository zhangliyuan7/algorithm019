package Week_02

type BSearchTree struct{
	val int
	left *BSearchTree
	right *BSearchTree
}

func (bst *BSearchTree)Search(n int)bool{
	root := bst
	for root.val!=n{
		if root.val<n{
			if root.right==nil{
				return false
			}
			root=root.right
		}
		if root.val>n{
			if root.left==nil{return false }
			root=root.left
		}
	}
	return true
}

func (bst *BSearchTree)Insert(n int ){
	root:=bst
	for root!=nil{
		if root.val==n{
			return
		}
		if root.val>n{
			if root.left==nil{
				root.left=&BSearchTree{val: n}
				return
			}
			root=root.left
		}
		if root.val<n{
			if root.right==nil{
				root.right=&BSearchTree{val: n}
				return
			}
			root=root.right
		}
	}
}

func (bst *BSearchTree)Delete(n int ){
	root:=bst
	for root!=nil{
		if root.val>n{
			root=root.left
		}
		if root.val<n{
			root=root.right
		}
		if root.val==n{
			//右子树为空，直接用左子树替换
			if root.right==nil{
				root=root.left
				return
			}
			//右子树不为空，找到右子树中最小元素替换
			//并将最小元素的right赋值给最小元素位置(如果存在的话非nil，不存在自然为nil了)
			rRight:=root.right
			for rRight!=nil{
				if rRight.left==nil{
					root.val=rRight.val
					rRight=rRight.right
					return
				}
				rRight=rRight.left
			}
		}
	}
}
