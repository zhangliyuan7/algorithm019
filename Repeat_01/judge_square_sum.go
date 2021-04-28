package Repeat_01

import (
	"math"
)

//633
//
//最优解 费马定理:奇质数能表示为两个平方数之和的充分必要条件是该素数被4除余1
// 一个非负整数 cc 如果能够表示为两个整数的平方和，当且仅当 c的所有形如 4k+3 的质因子的幂均为偶数
// 质因数分解
// 欧拉证明五步：
// 第一步、“如果两个整数都能表示为两个平方数之和，则它们的积也能表示为两个平方数之和。”
// 第二步、“如果一个能表示为两个平方数之和的整数被另一个能表示为两个平方数之和的素数整除，则它们的商也能表示为两个平方数之和。”
// 第三步、“如果一个能表示为两个平方数之和的整数被另一个不能表示为两个平方数之和的整数整除，则它们的商也必有一个不能表示为两个平方数之和的因子。”
// 无穷递降法
// 扩展 ：拉格朗日四平方定理 ：每个自然数都可以表示为四个整数平方之和。 其中四个数字是整数
func judgeSquareSumIII(c int) bool {
	for base := 2; base*base <= c; base++ {
		// 如果不是因子，枚举下一个
		if c%base > 0 {
			continue
		}

		// 计算 base 的幂
		exp := 0
		for ; c%base == 0; c /= base {
			exp++
		}

		// 根据 Sum of two squares theorem 验证
		if base%4 == 3 && exp%2 != 0 {
			return false
		}
	}

	// 例如 11 这样的用例，由于上面的 for 循环里 base * base <= c ，base == 11 的时候不会进入循环体
	// 因此在退出循环以后需要再做一次判断
	return c%4 != 3
}

//双指针
func judgeSquareSumII(c int) bool {
	if c==0||c==1{
		return true
	}
	sqrt:=math.Sqrt(float64(c))
	if int(sqrt)*int(sqrt) == c {
		return true
	}
	i:=0
	j:=int(sqrt)
	for i<=j{
		total:=i*i+j*j
		if total>c{
			j--
		}else if total<c{
			i++
		}else{
			return true
		}
	}
	return false
}


// timeout ~
func judgeSquareSum(c int) bool {
	for i:=0;i<=c;i++{
		for j:=0;j<=c;j++{
			if i*i+j*j==c{
				return true
			}
		}
	}
	return false
}
// 改进的暴力法 仍然超时
func judgeSquareSumI(c int) bool {
	if c == 0 || c == 1 {
		return true
	}
	sqrt := math.Sqrt(float64(c))
	if int(sqrt)*int(sqrt) == c {
		return true
	}
	for i := 0; i <= int(sqrt)+1; i++ {
		for j := 0; j <= int(sqrt)+1; j++ {
			if i*i+j*j == c {
				return true
			}
		}
	}
	return false
}



