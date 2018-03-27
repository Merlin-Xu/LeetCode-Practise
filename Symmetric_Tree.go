/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil ||(root.Left == nil && root.Right == nil) {
        return true
    }
    if root.Left == nil || root.Right == nil{
        return false
    }
    return dfs(root.Left , root.Right)
}
func dfs(left *TreeNode,right *TreeNode) bool{
    if left == nil && right == nil {
          return true
    }
   
	if left.Val == right.Val && left != nil && right !=nil {
        le:=dfs(left.Left,right.Right)
        ri:=dfs(left.Right,right.Left)
        if le && ri{
		    return true
        }
        return false
	}
    return false
}