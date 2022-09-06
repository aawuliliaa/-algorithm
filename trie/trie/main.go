package main

type Node struct {
	isWord bool
	next   map[string]*Node
}
type trie struct {
	root *Node
}

func generateNode() *Node {
	return &Node{next: make(map[string]*Node)}
}
func NewTrie() *trie {
	return &trie{
		root: generateNode(),
	}
}

func (t *trie) Insert(word string) {
	cur := t.root
	for _, w := range []rune(word) {
		wStr := string(w)
		if cur.next[wStr] == nil {
			cur.next[wStr] = generateNode()
		}
		cur = cur.next[wStr]
	}
	cur.isWord = true

}
func (t *trie) Search(word string) bool{
	cur := t.root
	for _,w := range []rune(word){
		wStr := string(w)
		if cur.next[wStr]==nil{
			return false
		}
		cur=cur.next[wStr]
	}
	return cur.isWord
}
func (t *trie) Prefix(prefix string) bool{
	cur := t.root
	for _,w := range []rune(prefix){
		wStr := string(w)
		if cur.next[wStr]==nil{
			return false
		}
		cur=cur.next[wStr]
	}
	return true
}