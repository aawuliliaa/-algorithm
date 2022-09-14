package mergeLink21

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1==nil&&list2==nil{
		return nil
	}
	newNode := &ListNode{}
	node := newNode
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			node.Val = list1.Val
            list1 = list1.Next
		}else{
			node.Val = list2.Val
			list2 = list2.Next
		}
		node.Next = &ListNode{}
		node = node.Next
	}
	for list1!=nil{
		node.Val = list1.Val
		list1 = list1.Next
		if list1!=nil{
			node.Next = &ListNode{}
			node = node.Next
		}

	}
	for list2!=nil{
		node.Val = list2.Val
		list2 = list2.Next
		if list2!=nil{
			node.Next = &ListNode{}
			node = node.Next
		}

	}
	return newNode
}
