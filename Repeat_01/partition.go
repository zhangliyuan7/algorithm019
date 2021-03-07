package Repeat_01

// 131

func partition(s string)[][]string{
	tmp:=[]string{}
    r := [][]string{}
	dig(s,0,tmp,&r)
	return r
}

func dig(s string,index int ,tmp []string,r *[][]string) {
	// 如果不拷贝 结果有错误
	// if len(s)==index{
	// 		*r=append(*r,tmp)
	// 		return
	// }
	// 分割结束，没有剩余字符串可以切分了
     if index == len(s) {
        lst := make([]string, len(tmp))
        copy(lst, tmp)
        *r = append(*r, lst)
        return
    }
	// 尝试不同大小的分割
	for i:=index;i<len(s);i++{
		if isValidHW(s[index:i+1]){
			tmp=append(tmp,s[index:i+1])
			dig(s,i+1,tmp,r)
			tmp=tmp[:len(tmp)-1]
		}
	}
}

func isValidHW(s string)bool{
	l,r:=0,len(s)-1
	for l<r{
		if s[l]!=s[r]{
			return false
		}
		l++
		r--
	}
	return true
}




func partitionII(s string) [][]string {
	r:=[][]string{}
	tmp:=[]string{}
	huisu(s,0,tmp,&r)
	return r
}

func huisu(s string ,index int ,tmp []string ,result *[][]string){
	if index==len(s){
		copyTmp:=make([]string,len(tmp))
		copy(copyTmp,tmp)
		*result=append(*result,copyTmp)
		return
	}
	for i:=index;i<len(s);i++{
		if valid(s[index:i+1]){
			tmp=append(tmp,s[index:i+1])
			huisu(s,i+1,tmp,result)
			tmp=tmp[:len(tmp)-1]
		}
	}
}

func valid(s string)bool{
	l,r:=0,len(s)-1
	for l<=r{
		if s[l]!=s[r]{
			return false
		}
		l++
		r--
	}
	return true
}