package practice


//stack
func isValid(s string) bool {
	if len(s)%2!=0{
		return false
	}
	var stack Stack
	for _,c:=range s{
		if string(c)=="("{
			stack.Push(")")
		}
		if string(c)=="["{
			stack.Push("]")
		}
		if string(c)=="{"{
			stack.Push("}")
		}
		if (string(c) == ")"||string(c)=="}"||string(c)=="]")&&stack.Pop()!=string(c){
			return false
		}
	}
	if stack.Len()==0{
		return true
	}
	return false
}