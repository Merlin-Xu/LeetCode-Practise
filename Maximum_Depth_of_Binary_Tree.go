/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "fmt"
func maxDepth(root *TreeNode) int {
    return dfs(root , 0, 0)
}
func dfs(node *TreeNode,result int,final int)(res int){
    if node==nil{
        res = 0
        return
    }
    result++
    if node.Left == nil && node.Right == nil{
        if result>final{
            fmt.Println(result)
            fmt.Println(final)
            res = result
        }else{
            res = final
        }
        return
    }
    if node.Left != nil {
        res = dfs(node.Left,result,final)
        final = res
    }
    if node.Right != nil{
        res = dfs(node.Right,result,final)
        final = res
    }
    
    return
    
}