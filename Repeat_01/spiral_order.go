package Repeat_01


// 54
func spiralOrder(matrix [][]int) []int {
	r:=[]int{}
	for len(matrix)!=0{
		r=append(r, matrix[0]...)
		matrix=Roll(matrix[1:])
	}
	return r
}

func Roll(matrix [][]int)[][]int{
	ln:=len(matrix)
	if ln== 0{
		return [][]int{}
	}
	ln2:=len(matrix[0])
	newMatrix:=make([][]int,ln2)
	for i:=range newMatrix{
		newMatrix[i]=make([]int,ln)
	}
	col:=len(newMatrix)-1
	for i:=0;i<ln;i++{
		for j:=0;j<ln2;j++{
			newMatrix[col-j][i]=matrix[i][j]
		}
	}
	return newMatrix
}
