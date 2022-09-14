package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {

	pathSumList := make([][]int, 0)
	if root == nil {
		return pathSumList
	}
	sumR(root, targetSum, 0, "", &pathSumList)
	return pathSumList
}
func sumR(node *TreeNode, targetSum, sum int, path string, pathSumList *[][]int) {
	sum += node.Val
	path = path +fmt.Sprintf("%d",node.Val) +"|"
	if node.Left == nil && node.Right == nil && sum == targetSum {
		pathList := strings.Split(path,"|")
		pathIntList := make([]int,0, len(pathList)-1)
		for _,p :=range pathList{
			if len(p)>0{
				pathInt ,err := strconv.Atoi(p)
				if err!=nil{
					panic(err)
				}
				pathIntList = append(pathIntList, pathInt)
			}
		}

		*pathSumList = append(*pathSumList, pathIntList)
	}
	if node.Left != nil {
		sumR(node.Left, targetSum, sum, path, pathSumList)
	}
	if node.Right != nil {
		sumR(node.Right, targetSum, sum, path, pathSumList)
	}
}