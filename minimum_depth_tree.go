/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import "fmt"
func minDepth(root *TreeNode) int {
    arr:=[]*TreeNode{root}
    if nil == root {
        return 0
    }
    if root.Left == nil && root.Right == nil {
            return 1
    }
    root.Val = 1
    for i:=0;i<len(arr);i++{
        node := arr[i]
        if node.Left == nil && node.Right == nil {
            return node.Val
        }
        if node.Left == nil || node.Right == nil{
            if node.Left != nil{
                node.Left.Val = node.Val+1
                arr = append(arr,node.Left)
                fmt.Println(node.Left.Val)
            }else{
                node.Right.Val = node.Val+1
                arr = append(arr,node.Right)        
                fmt.Println(node.Right.Val)
            }
            
        }else{
            node.Left.Val = node.Val+1
            node.Right.Val = node.Val+1
            arr = append(arr,node.Left)
            arr = append(arr,node.Right)
        }
    }
    return 0
}