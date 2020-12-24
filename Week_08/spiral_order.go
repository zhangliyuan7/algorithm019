package Week_08

//54
func SpiralOrder(matrix [][]int) []int {
	var r []int
	for len(matrix)!=0{
		r=append(r,matrix[0]...)
		matrix=Roll(matrix[1:])
	}
	return r
}

func Roll(matrix [][]int)[][]int{
	if len(matrix)==0{
		return [][]int{}
	}
	var matrixNew =make([][]int,len(matrix[0]))
	for i:=range matrixNew{
		matrixNew[i]=make([]int, len(matrix))
	}
	n:=len(matrixNew)-1
	newrow:=n
	for n>=0{
		for old_row:=0;old_row<len(matrix);old_row++ {
			matrixNew[newrow-n][old_row] =matrix[old_row][n]
		}
		n--
	}
	return matrixNew
}