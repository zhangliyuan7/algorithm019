学习笔记

##### templates
    ```
        recursion:
        python:
         def recursion(level,param1,param2...):
            //terminator
            if level>max_level:
                process_result
                return 
            //process
            process(level,data)
            //drill down
            self.recursion(level+1,p1,p2...)
            //reverse the current level status if you need
         c/c++:
            void recursion(int level,int param){
                //terminator
                if (level>MAX_LEVEL){
                    return ;
                }
                //process
                process(level,param)
                //drill down
                recursion(level+1,params)
                //reverse the current level status if needed
            }      
            
        dynamic programming:
        
        divide_conquer: //分开，分别-克服，攻克
        python:
        divide_conquer(problem,param1,param2...):
            //terminator
            if problem is None:
                print result
                return 
            //prepare data
            data=prepare_data(problem)
            subproblems=split_problem(problem,data)
            //conquer subproblems
            subresult1=self.divide_conquer(subproblem[1],p1...)
            subresult2=self.divide_conquer(subproblem[2],p1...)
            ......
            //process and generate the final result
            result=process_result(subresult1,subresult2...)
            //revert the current level status
        
        c/c++:
        int divide_conquer(problem *problem,int params){
            // recusion terminator
            if (problem == nullptr){
                process_result
                return return_result;
            }
            //process current problem
            subproblem = split_problem(problem,data)
            subresult1=divide_conquer(subproblem[1],param...)
            ......
            //merge 
            result = process_and_merge(subresult...)
            return 0;
        }
    ```
##### 单词接龙
    ```
        1 bfs 层序遍历思想+循环产生一位变化的新string ，与map中单词比对，存在则加入下层循环查找
          注意使用过的单词要清空，另外 go的正则实在是太tm慢了，不要用，不如循环生成单词直接查map来的快
        2 图方式，不是很明白  
        3 双向bfs ，如果相遇，就是2个方向的遍历层级和/2+1
    ```
##### 排序数组中查找目标元素的第一个位置和最后一个位置 no.34
    ```
        1。二分查找，找到目标target,向左找到左边界，再从mid开始用二分找右边界
        2。二分找到目标值一个索引，然后左右+1/-1扩散，获取两个边界（最差情况O(n))
    ```
##### 最小基因变化 no.433
    ```
        双向交替bfs，从end和start交替出发，找到交汇点之时就是结果层数
        注意点：
            浏览完当前层之后 要从visited和bankMap里去掉 important，防止下层又生成回原值
            visited只存有下一层遍历需要的两个方向的上层对比值 important
            入循环时记录队列长度，append下一层节点后，要截断队列为新队列
            判断start的queue和end的queue哪个长，哪个短则为下次循环的queue
            注意跳过节点自身值
    ```