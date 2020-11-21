package w3

func LetterCombinations(digits string) []string {
	var digitsMap = make(map[string][]string)
	digitsMap["1"]=[]string{}
	digitsMap["2"]=[]string{"a","b","c"}
	digitsMap["3"]=[]string{"d","e","f"}
	digitsMap["4"]=[]string{"g","h","i"}
	digitsMap["5"]=[]string{"j","k","l"}
	digitsMap["6"]=[]string{"m","n","o"}
	digitsMap["7"]=[]string{"p","q","r","s"}
	digitsMap["8"]=[]string{"v","t","u"}
	digitsMap["9"]=[]string{"y","w","x","z"}
	return recursionLetterCombinations("",digits,digitsMap)
}

//next is input next number string
//s is result string
func recursionLetterCombinations(s string,digits string ,digitsMap map[string][]string)[]string{
	var result []string
	//terminator
	if digits==""{
		result=append(result,s)
		return result
	}
	//process
	var rangeList = digitsMap[digits[0:1]]
	//drill down
	for i:=0;i<len(rangeList);i++{
		result=append( result,recursionLetterCombinations(s+rangeList[i],digits[1:],digitsMap)...)
	}
	//reverse
	return result
}