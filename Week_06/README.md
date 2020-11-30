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
    