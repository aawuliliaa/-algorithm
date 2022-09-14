package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeAList := make([]*ListNode, 0)
	nodeBList := make([]*ListNode, 0)
	nodeA := headA
	nodeB := headB
	for nodeA != nil {
		nodeAList = append(nodeAList, nodeA)
		nodeA = nodeA.Next
	}
	for nodeB != nil {
		nodeBList = append(nodeBList, nodeB)
		nodeB = nodeB.Next
	}
	i := len(nodeAList) - 1
	j := len(nodeBList) - 1
	flage := false
	for i >= 0 && j >= 0 {
		if nodeAList[i].Val == nodeBList[j].Val && nodeAList[i] == nodeBList[j] {
			i--
			j--
			flage = true
		} else {
			break
		}
	}
	if flage {
		return nodeAList[i+1]
	}
	return nil
}
