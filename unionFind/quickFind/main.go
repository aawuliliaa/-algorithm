package main

// 我们的第一版Union-Find
type UnionFind struct {
	id []int // 我们的第一版Union-Find本质就是一个数组
}

func New(size int) *UnionFind {
	id := make([]int, size)
	for i := 0; i < len(id); i++ {
		id[i] = i
	}
	return &UnionFind{id: id}
}
func (u *UnionFind) GetSize() int {
	return len(u.id)
}

//查找元素p所对应的集合编号
//O(1)复杂度
func (u *UnionFind) find(p int) int {
	if p < 0 || p >= len(u.id) {
		panic("element out of range")
	}
	return u.id[p]
}

//查看元素p和元素q是否属于一个集合==即验证值是否相等
//O(1)复杂度
func (u *UnionFind) IsConnected(p int, q int) bool {
	return u.id[p] == u.id[q]
}

//合并元素p和元素q所属的集合==让p集合的值等于q集合的值或让q集合的值等于p集合的值
//O(n)复杂度
func (u *UnionFind) UnionElements(p int, q int) {
	pId := u.find(p)
	qId := u.find(q)
	for i:=0;i<len(u.id);i++{
		if u.id[i]==pId{
			u.id[i]=qId
		}
	}
}
