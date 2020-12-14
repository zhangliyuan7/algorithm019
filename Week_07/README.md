学习笔记
##### Trie Tree 208
    ```
        type Trie struct{
            IsEnd bool
            Next map[string]*Trie || [26]*Trie
        }
        字典树 ：空间换时间
        特性：
            跟节点不包含字符，其余每个节点都只含有一个字符
            每个单词是由通路路径组成
            每个节点的子节点包含字符串都各不相同
       
            ～ 包含一个词结束符号
            ～ 可通过前缀查找单词
            ～ [26]int
        方法：
            插入词
            查询词
            前缀查询
            
        Map实现比用列表空间占用小得多
        查词 ：查找到最后一个字符，检查end标示
        查前缀：不检查end标示，只需要都搜索到即可
    ```
##### 单词搜索 no79
    ```
        从单词第一个字符开始进行深度优先搜索
        四连通，所以dx dy标示四个方向的变化
        //东南西北
        dx=[1,0,-1,0]
        dy=[0,-1,0,1]
        每次传递word[1:],当前首字符索引位置，和棋盘
        func dfs(board [][]byte,i,j int ,word string)bool{}
        递归中止条件：
            1 当i,j超出棋盘，中止当前这支深度搜索
            2 当word长度为1 判断棋盘i，j位置是否与其相同
            3 word[0]!=board[i][j] 下一位字符不同于单词该位置字符
        保存状态 替换掉遍历过的位置，防止扫到本身
            dfs {
                ...
                tmp:=board[i][j]
                board[i][j]=null
                for condition {
                    if dfs (board,i,j,word[1:]){ 
                        return true
                    }
                }
                board[i][j]=tmp
                ...
            }
            
    ```
##### 单词搜索II no212
    ```
    1   典型的trie树不存储string 只用路径，但是对于此题会超时
        用单词构建Trie Tree，再遍历二维数组，加dfs回溯 将结果添加到结果集
        t=trie Tree.
        result =[]string
        insert words to trie
        for i:=range board[i]{
            for j:=range board[j]{
                dfs(board,i,j,"",t)
            }
        }
        func dfs(board ,i , j ,word, trieT){
            if i,j outOfSlice {return}
            word=word+board[i][j]
            //回溯准备
            tmp:=board[i][j]
            board[i][j]='#'
            //前缀是否合法
            if trieT.searchPrefix(){
                if trieT.search(word){
                    //结果集
                    result=append(result,word)
                }
                // 四连通 dx dy
                for k:=0;k<4;k++{
                    dfs(board,i+dx[k],y+dy[k]，word,trieT)
                }
            }
            //回溯 恢复状态
            board[i][j]=tmp
            //前缀非法或者递归结束 退出
        }
    2   对于常规End的节点，存储对应的word，当访问到该节点时，直接塞入该word to 结果集
        不必每次都search完整，仅需要每次进一步，即下一个字符（e-s-w-n）
        code like this：
            func findWords(board [][]byte, words []string) []string {
            	trie,res := buildTrieTree(words),make([]string,0)
            	for i := 0; i < len(board); i++{
            		for j := 0; j < len(board[0]); j++{
            			dfsx(board,trie,i,j,&res)
            		}
            	}
            	return res
            }
            func dfsx(board [][]byte, node *trieNode, i, j int, res *[]string){
            	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]){
            		return
            	}
            	c := board[i][j]
            	if board[i][j] == '#' || node.nodes[c - 'a'] == nil{
            		return
            	}
            	node = node.nodes[c - 'a']
            	if node.word != ""{
            		*res = append(*res,node.word)
            		node.word = ""
            	}
            	board[i][j] = '#'
            	dfsx(board,node,i-1,j,res)
            	dfsx(board,node,i+1,j,res)
            	dfsx(board,node,i,j-1,res)
            	dfsx(board,node,i,j+1,res)
            	board[i][j] = c
            }
            type trieNode struct{
            	word string
            	nodes []*trieNode
            }
            func buildTrieTree(words []string)*trieNode{
            	root := trieNode{nodes:make([]*trieNode,26)}
            	for _, word := range words{
            		cur := &root
            		for i := 0; i < len(word); i++{
            			c := word[i]
            			if cur.nodes[c - 'a'] == nil{
            				cur.nodes[c - 'a'] = &trieNode{nodes:make([]*trieNode,26)}
            			}
            			cur = cur.nodes[c - 'a']
            		}
            		cur.word = word
            	}
            	return &root
            }
    ```
