# LeetCode-Practise
The road of Leetcode Algorithm Chapter

```
Golang solution
```
#####DFS
<pre><code>
	/**
	 * DFS  的模版
	 *@src 原数组
	 *@index 递增或是递减的游标
	 *@target 不断变化的目标值
	 *@template 本次循环的解集
	 *@result 最终返回的解集
	 */
	func DFS(srcArr []int, index int,target int,tempSolution []int, result []int){
		target //不符合时退出此次遍历
		target //边界符合时给最终解集赋值
		for {//循环元数组或是该y轴
			if statement //寻找边界和特殊情况
			 continue //退出循环
			
			处理本次循环的结果
			 dfs()游标移动横向或是多游标进行多次移动进入下次迭代
			 reulst//给最终结果集赋值
		}
		return result
   }

</code></pre>
#####BFS
<pre><code>
	/**
	 * [BFS] 一般bfs有个灵活的集合需要用while 。go没有while 可以用for 解决
	 * 在每一个点的遍历都要做一个状态保存譬如二叉树或是图的当前路径的状态。以便找出最优
	 * 
	 * 
	 *
	 */
	func BFS(){
		while(arr){
			if statment //不符合条件跳出
			//符合或是部分符合继续查找下去
			……
		}
	}
</code></pre>
