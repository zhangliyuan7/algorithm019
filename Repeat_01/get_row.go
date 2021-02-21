package Repeat_01

// 119
// yanghui triangle
func getRow(rowIndex int )[]int{
	lastRow:=[]int{1}
	for i:=1;i<rowIndex+1;i++{
		ithRow:=make([]int,i+1)
		for j:=0;j<=i;j++{
			if j==i||j==0{
				ithRow[j]=1
			}else if j>0&&j<i{
				ithRow[j]=lastRow[j]+lastRow[j-1]
			}
		}
		lastRow=ithRow
	}
	return lastRow
}

// second way : 行内递推
func getRow2(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	for i := 1; i <= rowIndex; i++ {
		row[i] = row[i-1] * (rowIndex - i + 1) / i
	}
	return row
}