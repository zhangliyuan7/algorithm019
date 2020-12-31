package Week_09

import (
	"fmt"
	"log"
	"strings"
)

// 10
// . one rune  * is before rune repeat (0-n)
// hard ...递归可以理解透了 ，动态规划还是较为困难
func isMatch(s string, p string) bool {
	//这个地方注意，不能写成
	//if len(p)==0&&len(s)!=0{
	//	return false
	//}
	// 因为会忽略一个true的情况，即s也为0 ，这时，后续逻辑会将firstmatch设置为false，所以整个结果会错误

	if len(p)==0{
		return len(s)==0
	}
	firstMatch:=false
	// 一定要判断s长度，避免panic ，判断s的字符是否匹配p字符或者'.'
	if len(s)!=0&&(s[0]==p[0]||p[0]=='.'){
		firstMatch=true
	}
	// *号情况有两种 ，一种是p的 'n*'中的n出现零次，另外就是出现1-n次
	// 出现0次的结果 即将p的'n*'去掉，继续匹配s当前字符  s,p[2:]  注意p的长度判断，防止panic
	// >0次结果，即去掉当前匹配到的s字符，p规则不变，匹配s后续字符 s[1:],p
	if len(p)>1&&p[1]=='*'{
		return (firstMatch&&isMatch(s[1:],p))||isMatch(s,p[2:])
	}
	return firstMatch&&isMatch(s[1:],p[1:])
}

// 状态机版本
func debug(v ...interface{}) {
	log.Println(v...)
}

func toString(i interface{}) string {
	switch i.(type) {
	case int:
		return fmt.Sprintf("%v", i)
	case string:
		return fmt.Sprintf("%v", i)
	case bool:
		return fmt.Sprintf("%v", i)
	default:
		return fmt.Sprintf("%p", i)
	}
}

func isMatch2(s string, p string) bool {
	begin := new(Node)
	begin.C = '>'
	begin.Size = generatePattern(begin, p, 0)
	debug(begin.String())
	return check(begin, s, 0)
}

type Node struct {
	C        byte
	Parent   *Node
	Children map[byte][]*Node
	End      bool
	Size     int
}

func (n *Node) String() string {
	return n.StringLevel(0, make(map[*Node]bool))
}

func (n *Node) StringLevel(level int, finishNodes map[*Node]bool) string {
	r := make([]string, 0)
	if n.End {
		r = append(r, fmt.Sprintf("  id%s{%v};", toString(n), string(n.C)))
	} else {
		r = append(r, fmt.Sprintf("  id%s(%v);", toString(n), string(n.C)))
	}
	finishNodes[n] = true
	for k, v := range n.Children {
		for _, c := range v {
			if _, ok := finishNodes[c]; !ok {
				r = append(r, c.StringLevel(level+1, finishNodes))
			}
			r = append(r, fmt.Sprintf("  id%s -- %s --> id%s;", toString(n), string(k), toString(c)))
		}
	}
	return strings.Join(r, "\n")
}

func (n *Node) Append(c byte, child *Node) {
	m := n.Children
	if m == nil {
		m = make(map[byte][]*Node)
		n.Children = m
	}
	list := m[c]
	if list == nil {
		list = make([]*Node, 0)
	}
	for _, v := range list {
		if v == child {
			m[c] = list
			return
		}
	}
	list = append(list, child)
	m[c] = list
}

func generatePattern(now *Node, str string, idx int) int {
	if len(str) <= idx {
		now.End = true
		return now.Size
	}
	vnow := now
	switch str[idx] {
	case '*':
		now.Size = 0
		now.Append(now.C, now)
	default:
		node := new(Node)
		node.C = str[idx]
		now.Append(str[idx], node)
		node.Parent = now
		node.Size = 1
		vnow = node
	}
	ret := generatePattern(vnow, str, idx+1)
	if ret == 0 {
		now.End = true
	}
	addParent := now
	for addParent.Parent != nil {
		if addParent.Size == 0 {
			debug(toString(vnow), " -> ", toString(addParent.Parent))
			addParent.Parent.Append(vnow.C, vnow)
			addParent = addParent.Parent
		} else {
			break
		}
	}
	return now.Size + ret
}

func check(now *Node, str string, idx int) bool {
	if len(str) <= idx {
		return now.End
	}
	list := now.Children['.']
	for _, v := range now.Children[str[idx]] {
		list = append(list, v)
	}
	for _, v := range list {
		r := check(v, str, idx+1)
		if r {
			return true
		}
	}
	return false
}
