package Week_07

//repeat 22
func GenerateParenthesis(n int) []string {
	var r []string
	dgui(0,0,n,"",&r)
	return r
}

func dgui(left int ,right int , n int ,r string,res *[]string ){
	if left==right&&left==n{
		*res=append(*res,r)
		return
	}
	if left>n{
		return
	}
	if right>left{
		return
	}
	dgui(left+1,right,n,r+"(",res)
	dgui(left,right+1,n,r+")",res)
}