package gopl

import (
	"bufio"
	"fmt"
	"os"
)

func InputDuplicate(){
	m:=make(map[string]int )
	input:=bufio.NewScanner(os.Stdin)
	for input.Scan(){
		m[input.Text()]+=1
	}
	fmt.Println(m)
	for k:=range m{
		fmt.Print(fmt.Sprintf(k))
	}
}

func ReadFileDuplicate(){
	count:=make(map[string]int)
	files:=os.Args[1:]
	if len(files)==0{
		countLine(os.Stdin,count)
	}else{
		for _,arg:=range files{
			f,err:=os.Open(arg)
			if err!=nil{
				fmt.Errorf("open file %s",err.Error())
				continue
			}
			countLine(f,count)
			f.Close()
		}
	}
	for line, n := range count {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLine(file *os.File,count map[string]int){
	input:=bufio.NewScanner(file)
	for input.Scan(){
		count[input.Text()]++
	}
}