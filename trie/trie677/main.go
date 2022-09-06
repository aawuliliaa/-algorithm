package main

type Node struct {
	next map[string]*Node
	val  int
}

type MapSum struct {
	root *Node
}
func generateNode() *Node {
	return &Node{next: make(map[string]*Node)}
}
func Constructor() MapSum {
	return MapSum{root: generateNode()}
}

func (this *MapSum) Insert(key string, val int) {
	cur := this.root
	for _,w :=range []rune(key){
		wStr :=string(w)
		if cur.next[wStr]==nil{
			cur.next[wStr] = generateNode()
		}
		cur = cur.next[wStr]
	}
	cur.val = val
}

func (this *MapSum) Sum(prefix string) int {
	cur := this.root
	for _,w :=range []rune(prefix){
		wStr := string(w)
		if cur.next[wStr]==nil{
			return 0
		}
		cur = cur.next[wStr]
	}
	return this.sumR(cur)

}
func (this *MapSum)sumR(node *Node)int{
	if len(node.next)==0{
		return node.val
	}
	res := node.val

	for _,n :=range node.next{
		res += this.sumR(n)
	}
	return res

}
