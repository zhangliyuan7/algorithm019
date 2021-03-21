package Repeat_01


//73
func setZeroesI(matrix [][]int) {
	row, col := make(map[int]int), make(map[int]int)
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row[i] = 1
				col[j] = 1
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if row[i] == 1 || col[j] == 1 {
				matrix[i][j] = 0
			}
		}
	}
}
// 这是看的另一个题解 ，这个写法感觉比较好
func setZeroesII(matrix [][]int) {
	row := len(matrix)
	col := len(matrix[0])
	fr := false
	fc := false
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					fr = true
				}
				if j == 0 {
					fc = true
				}
				//标记行列是否存在0 这个包含了第一行和第一列的判断，重写的比官方题解的多，但是综合
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			// 这一步比较重要，通过第一列和第一行表示是否存在0，如果1<=row<=len-1 1<=col<=len-1 中 没有0 那么将采用原始的第一列和第一行进行标注
			// 所以结果正确，最后只需要维护第一行和一列的0即可
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if fr {
		for j := 1; j < col; j++ {
			matrix[0][j] = 0
		}
	}

	if fc {
		for i := 1; i < row; i++ {
			matrix[i][0] = 0
		}
	}
}
