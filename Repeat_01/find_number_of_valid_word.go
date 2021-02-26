package Repeat_01

// 	1178
// 	占位2进制 子集获取 important!!!!!!!!!!!!!!!!
//	n =	originBits
//	for n >	0 {
//		n = ( n - 1 ) && originBits
//	}
func findNumOfValidWords(words []string, puzzles []string) []int {
	// 结果集
	r:=make([]int,len(puzzles))

	// 统一进行的puzzle的占位二进制转换
	bitPuzzles:=[]int{}
	for i:=0;i<len(puzzles);i++{
		mask:=0
		for p:=range puzzles[i] {
			mask|=1<<(puzzles[i][p]-'a')
		}
		bitPuzzles=append(bitPuzzles,mask)
	}
	bitCount:=make(map[int]int)
	// 将words 全部转化为占位二进制，并用map存储压缩候选word数量，
	for i:=0;i<len(words);i++{
		mask:=0
		for w:=range words[i] {
			mask|=1<<(words[i][w]-'a')
		}
		bitCount[mask]+=1
	}
	//fmt.Println(bitCount,bitPuzzles)

	for i:=0;i<len(puzzles);i++{
		// 首个字母的占位二进制表示数 'c' => 100  'a'=>1   'd' =>1000   'z'=>10000000000000000000000000
		firBit:=1<<(puzzles[i][0]-'a')
		// 该puzzle对应的二进制值
		n:=bitPuzzles[i]
		// n 变量
		for n > 0 {
			// 该子集 必须包含首个字母 && 存在于words中,第一次进入循环直接判断puzzle自身是否存在于words中
			if (firBit&n != 0) && bitCount[n] > 0 {
				// 添加到结果集中
				r[i] += bitCount[n]
			}

			// 获取子集，去掉一个1
			n = (n - 1) & bitPuzzles[i]
			// example :
			// 11001 => |ade|
			//			& 11000 => 11000 => |de|
			//						-1
			//			& 10111 => 10001 => |ae|
			//						-1
			//			& 10000 => 10000 => |e|
			// 						-1
			//			& 01111 => 01001 => |ad|
			//						-1
			//			& 01000 => 01000 => |d|
			//						-1
			//			& 00111 => 00001 => |a|
			//						-1
			//			n = 00000 => break
		}
	}
	return r
}
