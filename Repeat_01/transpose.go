package Repeat_01

//867
func transpose(matrix [][]int) [][]int {
	row:=len(matrix)
	col:=len(matrix[0])
	newMatrix:=make([][]int,col)
	for i:=range newMatrix{
		newMatrix[i]=make([]int,row)
	}
	for i:=0;i<row;i++{
		for j:=0;j<col;j++{
			newMatrix[j][i]=matrix[i][j]
		}
	}
	return newMatrix
}