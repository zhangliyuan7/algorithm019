package practice
// 1 step or 2 step one move
// how many times can moving on n stair
// f1=1
// f2 =2
// f3=f1+f2 or f3 = f1+f1+f1 or f2+f1
//fi的状态方程，由于只能爬一步或者两步，fi是从fi-2 或fi-1来的，fi=fi-1+fi-2 两种方法数的集合

func ClimbStair(n int)int{
	if n<=2{
		return n
	}
	var p,q int
	p=1
	q=2
	for i:=3;i<=n;i++{
		temp:=p
		p=q
		q=temp+q
	}
	return p+q
}

func ClimbStair2(n int )int {
	if n<=2{
		return n
	}
	var f1 = 1
	var f2 = 2
	for i:=3;i<n;i++{
		tmp:=f1+f2
		f1=f2
		f2=tmp
	}
	return f1+f2
}