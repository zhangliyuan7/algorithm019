package Week_06


// 	    [1],
// 	   [1,1],
// 	  [1,2,1],
// 	 [1,3,3,1],
// 	[1,4,6,4,1]
//[1,5,10,10,5,1]
func Generate(numRows int) [][]int {
	var r [][]int
	if numRows==0{
		return r
	}
	r=append(r,[]int{1})
	if numRows==1{
		return r
	}
	for i:=1;i<numRows;i++{
		var row []int
		for j:=0;j<=i;j++{
			if j-1<0{
				row=append(row,r[i-1][j])
			}else if j>len(r[i-1])-1{
				row=append(row,r[i-1][j-1])
			}else{
				row=append(row,r[i-1][j]+r[i-1][j-1])
			}

		}
		r=append(r,row)
	}
	return r
}