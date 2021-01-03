package Week_09

import (
	"math"
)

// 最大子矩阵
func GetMaxMatrix(matrix [][]int) []int {
	res:=make([]int,4)
	m:=len(matrix)
	n:=len(matrix[0])
	total:=math.MinInt32
	for i:=0;i<m;i++{
		for j:=i;j<m;j++{
			var sumColumn = make([]int,n)
			for k:=0;k<n;k++{
				for u:=i;u<=j;u++{
					sumColumn[k]+=matrix[u][k]
				}
			}
			l,r,max:=searchMaxSubsequence(sumColumn)
			//fmt.Println(l,r,max)

			if max>total{
				total=max
				res=[]int{i,l,j,r}
				//fmt.Println("res:",res)
			}
		}
	}
	return res
}
func searchMaxSubsequence(sc []int )(left int,right int,max int ){
	left,right=0,0
	tmpbegin:=left
	max=sc[0]
	dpi:=sc[0]
	for i:=1;i<len(sc);i++{
		if dpi>0{
			dpi+=sc[i]
		}else{
			dpi=sc[i]
			tmpbegin=i
		}
		if dpi>max{
			max=dpi
			left=tmpbegin
			right=i
		}
	}
	return
}