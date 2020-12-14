package Week_07

import (
	"fmt"
	"testing"
)

func TestFindWords(t *testing.T) {
	//fmt.Println(FindWords([][]byte{{'o','a','a','n'},{'e','t','a','e'},{'i','h','k','r'},{'i','f','l','v'}},[]string{"oath","pea","eat","rain"}))
	//fmt.Println(FindWords([][]byte{{'o','a','a','n'},{'e','t','a','e'},{'i','h','k','r'},{'i','f','l','v'}},[]string{"oaao","pea","eat","rain"}))
	fmt.Println(FindWords([][]byte{{'a','a'}},[]string{"a"}))
}