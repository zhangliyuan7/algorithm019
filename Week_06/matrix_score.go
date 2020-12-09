package Week_06

import "math"

//官方的比较巧妙，思路相同，都是先翻转所有行，使第一列为1 ，之后翻转所有列，
//官方题解不需要翻转，只需要记录列变换后可能的1的和直接和最终结果相加
func matrixScore(a [][]int) int {
	m, n := len(a), len(a[0])
	ans := 1 << (n - 1) * m
	for j := 1; j < n; j++ {
		ones := 0
		for _, row := range a {
			if row[j] == row[0] {
				ones++
			}
		}
		if ones < m-ones {
			ones = m - ones
		}
		ans += 1 << (n - 1 - j) * ones
	}
	return ans
}

// 861
func matrixScoreS(A [][]int) int {
	if len(A)==0{
		return 0
	}
	rowLength:=len(A[0])
	for i:=range A{
		if A[i][0]==0{
			A[i]=rowTurn(A[i])
		}
	}

	for i:=0;i<rowLength;i++{
		zeroCount,oneCount:=0,0
		for j:=0;j<len(A);j++{
			if A[j][i]==0{
				zeroCount+=1
			}else{
				oneCount+=1
			}
		}
		if zeroCount>oneCount{
			A=colTurn(A,i)
		}
	}
	return calculateBinary(A,rowLength)
}

func colTurn(num [][]int,k int )[][]int{
	for i:=0;i<len(num);i++{
		if num[i][k]==0{
			num[i][k]=1
		}else{
			num[i][k]=0
		}
	}
	return num
}

func rowTurn(num []int)[]int{
	for i:=range num{
		if num[i]==0{
			num[i]=1
		}else{
			num[i]=0
		}
	}
	return num
}

func calculateBinary(nums [][]int,rowlength int  )int {
	r:=0
	for n:=range nums {
		for i := rowlength - 1; i >= 0; i-- {
			r+=int(math.Pow(2, float64(rowlength-1-i))) * nums[n][i]
		}
	}
	return r
}