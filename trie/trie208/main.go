package main

type Node struct {
	isWord bool
	next   map[string]*Node
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{root: generateNode()}
}
func generateNode() *Node {
	return &Node{next: make(map[string]*Node)}
}
func (this *Trie) Insert(word string) {
	cur := this.root
	for _,w :=range word{
		wStr := string(w)
		if cur.next[wStr] ==nil{
			cur.next[wStr] = generateNode()
		}
		cur = cur.next[wStr]
	}
	cur.isWord = true
}

func (this *Trie) Search(word string) bool {
	cur := this.root
	for _,w :=range word{
		wStr := string(w)
		if cur.next[wStr] ==nil{
			return false
		}
		cur = cur.next[wStr]
	}
	return cur.isWord

}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for _,w :=range prefix{
		wStr := string(w)
		if cur.next[wStr] ==nil{
			return false
		}
		cur = cur.next[wStr]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
