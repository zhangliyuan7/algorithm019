package Repeat_01

//归零律：a ⊕ a = 0
//恒等律：a ⊕ 0 = a
//交换律：a ⊕ b = b ⊕ a
//结合律：a ⊕b ⊕ c = a ⊕ (b ⊕ c) = (a ⊕ b) ⊕ c
//自反：a ⊕ b ⊕ a = b
//注意："前n个正整数的排列的异或" 这代表 这个数组中的数字 是从1-n的整数排列，只不过在数组中位置打散
// 1734
func decode(encoded []int) []int {
	n:=len(encoded)+1
	total:=0
	odd:=0
	perm:=make([]int,n)
	/*

	   encoded=  6   5   4   6
	            / \ / \ / \ / \
	   perm=   2   4   1   5   3

	   perm[0] = (2^4^1^5^3)^(4 ^ 1 ^ 5 ^ 3)
	                           \ /     \ /
	           = (2^4^1^5^3)^(  5   ^   6)
	           =（前n个正整数的排列的异或）^（encoded中第奇数位的异或，第1,3个）
	           = 1^3
	           = 2

	   perm[n]=perm[n-1]^encoded[n-1]
	*/
	for i:=1;i<=n;i++{
		total^=i
	}
	for i:=1;i<n;i+=2{
		odd^=encoded[i]
	}
	perm[0]=total^odd
	for i:=1;i<n;i++{
		perm[i]=perm[i-1]^encoded[i-1]
	}
	return perm
}

