package Repeat_01

type NumMatrix struct {
	dp [][]int
}

func ConstructorNumMatrix(matrix [][]int) NumMatrix {
	ln := len(matrix) + 1
	if ln-1 == 0 {
		return NumMatrix{}
	}
	ln2 := len(matrix[0]) + 1
	dp := make([][]int, ln)
	for i := range dp {
		dp[i] = make([]int, ln2)
	}
	for i := 0; i < ln-1; i++ {
		for j := 0; j < ln2-1; j++ {
			dp[i+1][j+1] = dp[i][j+1] + dp[i+1][j] + matrix[i][j] - dp[i][j]
		}
	}
	return NumMatrix{dp: dp}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.dp[row2+1][col2+1] - this.dp[row1][col2+1] - this.dp[row2+1][col1] + this.dp[row1][col1]
}

//type NumMatrix struct {
// 	dp [][]int
// }

// func Constructor(matrix [][]int) NumMatrix {
// 	ln:= len(matrix)
//     if ln==0{
//         return NumMatrix{}
//     }
// 	ln2:=len(matrix[0])
// 	dp:=make([][]int,ln)
// 	for i:=range dp{
// 		dp[i]=make([]int,ln2)
// 	}
// 	dp[0][0]=matrix[0][0]
//     for i:=1;i<ln;i++{
//         dp[i][0]=dp[i-1][0]+matrix[i][0]
//     }
//     for i:=1;i<ln2;i++{
//         dp[0][i]=dp[0][i-1]+matrix[0][i]
//     }
// 	for i:=1;i<ln;i++{
// 		for j:=1;j<ln2;j++{
// 			dp[i][j]=dp[i-1][j]+dp[i][j-1]+matrix[i][j]-dp[i-1][j-1]
// 		}
// 	}
// 	return NumMatrix{dp:dp}
// }

// func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
// 	return this.dp[row1-1][col1-1]+this.dp[row2][col2]-this.dp[row2][col1-1]-this.dp[row1-1][col2]
// }
