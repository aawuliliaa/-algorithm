package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil{
		return false
	}
	flag := false
	sumR(root,targetSum,0,&flag)
	return flag
}
func sumR(node *TreeNode,targetSum ,sum int,flag *bool)  {
	sum += node.Val
	if node.Left==nil&&node.Right==nil&&sum == targetSum{
		*flag = true
	}
	if node.Left!=nil&&!*flag{
		sumR(node.Left,targetSum,sum,flag)
	}
	if node.Right!=nil&&!*flag{
		sumR(node.Right,targetSum,sum,flag)
	}
}

