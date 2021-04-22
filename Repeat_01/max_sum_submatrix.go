package Repeat_01

import "math"

//363
// standard resolution
func maxSumSubmatrix(matrix [][]int, k int) int {
	row:=len(matrix)
	col:=len(matrix[0])
	max:=math.MinInt32
	// 前缀和数组 不改变原数组
	preSum:=make([][]int,row)
	for x:=range preSum{
		preSum[x]=make([]int,col)
	}
	// 第一列初始化前缀和
	for i:=0;i<row;i++{
		preSum[i][0]=matrix[i][0]
	}
	// 填充前缀和数组
	for i:=0;i<row;i++{
		for j:=1;j<col;j++{
			preSum[i][j]=preSum[i][j-1]+matrix[i][j]
		}
	}
	// 固定左右边界 枚举区域大小
	for left:=0;left<col;left++{
		for right:=left;right<col;right++{
			// 上边界
			for top:=0;top<row;top++{
				area := 0
				//枚举上，下边界区域的和
				for bottom := top; bottom < row; bottom++ {
					// 最左边界
					if left==0{
						area += preSum[bottom][right]
					}
					// 偏右-左边界
					if left>0{
						area += preSum[bottom][right] - preSum[bottom][left-1]
					}
					if area > max && area <= k {
						max = area
					}
				}
			}
		}
	}
	return max
}
// bad dp ,it doesn't work ,but i like
func maxSumSubMatrix(matrix [][]int, k int) int {
	row:=len(matrix)
	col:=len(matrix[0])
	matrixMaxLength:=min(row,col)
	dp:=make([][][]int ,row)
	max:=0
	for i:=range dp{
		dp[i]=make([][]int,col)
		for j:=range dp[i]{
			dp[i][j]=make([]int,matrixMaxLength)
		}
	}
	//边长为1的矩阵初始化
	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			dp[i][j][1]=matrix[i][j]
			if dp[i][j][1]>max&&dp[i][j][1]<k{
				max=dp[i][j][1]
			}
		}
	}
	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			for x:=2;x<=matrixMaxLength;x++{
				if i+x>row||j+x>col{
					continue
				}
				span:=0
				for more:=0;more<x;more++{
					span+=dp[i][j+more][1]
					span+=dp[i+more][j][1]
				}
				//刨除一个重复的角
				span-=dp[i+x-1][j+x-1][1]
				dp[i][j][x]=dp[i][j][x-1]+span
				if dp[i][j][x]>max&&dp[i][j][x]<k{
					max=dp[i][j][x]
				}
			}
		}
	}
	return max
}
