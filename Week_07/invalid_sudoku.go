package Week_07

//36
func IsValidSudoku(board [][]byte) bool {
	row := make(map[int ]map[int]bool)
	col:=make(map[int]map[int]bool)
	setS:=make(map[int]map[int]bool)
	for i:=0;i<9;i++{
		for j:=0;j<9;j++{
			if board[i][j]!='.'{
				if len(row[i])==0{
					row[i]=make(map[int]bool)
				}
				if len(col[j])==0{
					col[j]=make(map[int]bool)
				}
				if len(setS[i/3*3+j/3])==0{
					setS[i/3*3+j/3]=make(map[int]bool)
				}
				if _,ok:=row[i][int(board[i][j]-'0')];ok{
					return false
				}else{
					row[i][int(board[i][j]-'0')]=true
				}
				if _,ok:=col[j][int(board[i][j]-'0')];ok{
					return false
				}else{
					col[j][int(board[i][j]-'0')]=true
				}
				if _,ok:=setS[i/3*3+j/3][int(board[i][j]-'0')];ok{
					return false
				}else{
					setS[i/3*3+j/3][int(board[i][j]-'0')]=true
				}

			}
		}
	}
	return true
}
//best 位图！
func isValidSudoku(board [][]byte) bool {
	var row, col, block [9]uint16
	var cur uint16
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			cur = 1 << (board[i][j] - '1')  // 当前数字的 二进制数位 位置
			bi := i/3 + j/3*3  // 3x3的块索引号
			if (row[i] & cur) | (col[j] & cur) | (block[bi] & cur) != 0 { // 使用与运算查重
				return false
			}
			// 在对应的位图上，加上当前数字
			row[i] |= cur
			col[j] |= cur
			block[bi] |= cur
		}
	}
	return true
}