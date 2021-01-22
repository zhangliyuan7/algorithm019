package Repeat_01

// 6
//3 PAYPALISHIRING =>PAHNAPLSIIGYIR
//P   A   H   N
//A P L S I I G
//Y   I   R
//
// 4 PINALSIGYAHRPI ==>  PAYPALISHIRING
//P     I    N
//A   L S  I G
//Y A   H R
//P     I
func Convert(s string, numRows int) string {
	if len(s)<=numRows||numRows<=1{
		return s
	}
	r:=make([]string,numRows)
	down:=false
	curRow:=0
	for _,c:=range s{
		r[curRow]+=string(c)
		if curRow==numRows-1||curRow==0{
			down=!down
		}
		if down{
			curRow+=1
		}else{
			curRow-=1
		}
	}
	res:=""
	for i:=range r{
		res+=r[i]
	}
	return res
}