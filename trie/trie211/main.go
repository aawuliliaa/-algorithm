package main

import "fmt"

type Node struct {
	next   map[string]*Node
	isWord bool
}
type WordDictionary struct {
	root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{root: generateNode()}
}
func generateNode() *Node {
	return &Node{next: make(map[string]*Node)}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this.root
	for _, w := range []rune(word) {
		wStr := string(w)
		if cur.next[wStr] == nil {
			cur.next[wStr] = generateNode()
		}
		cur = cur.next[wStr]
	}
	cur.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.searchR(word, this.root, 0)
}
func (this *WordDictionary) searchR(word string, node *Node, index int) bool {
	if index == len(word) {
		return node.isWord
	}
	c := string([]rune(word)[index])
	if c != "." {
		if node.next[c] == nil {
			return false
		} else {
			return this.searchR(word, node.next[c], index+1)
		}
	}else{
		for w,_ :=range node.next{
			if this.searchR(word, node.next[w], index+1){
				return true
			}
		}
		return false
	}
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
func main() {
	obj := Constructor()
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")

	fmt.Println(obj.Search("pad"))
	fmt.Println(obj.Search("bad"))
	fmt.Println(obj.Search(".ad"))
	fmt.Println(obj.Search("b.."))
}

