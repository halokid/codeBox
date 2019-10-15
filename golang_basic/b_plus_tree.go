package main

import (
  "bytes"
  "fmt"
)

type Key int  //Key 是int类型的另外一个同义声明

//Node 是这个 struct 是另外一个同义声明
type Node struct {
  Leaf        bool
  N           int
  Keys        []Key
  Children    []*Node   // Children 就是自己
}

//Search node from B+ tree
func (x *Node) Search(k Key) (n *Node, idx int) {
  i := 0
  // N 为B+树的节点值， 在节点值的寻找范围内，并且节点的值要小于K,可见整个B+树的值，也就是数组是按从小到大排列的
  // 节点的值本身就是以数组来形式来储存的，只是整个 B+树 的数据结构，我们要建立一个结构体
  for i < x.N && x.Keys[i] < k {
    i += 1
  }

  if i < x.N && k == x.Keys[i] {    //找到了节点,节点的值等于寻找的值
    n, idx = x, i
  //沿着一条线都找不到值的话，而且节点也没有叶子了，那么就证明这个节点本身就是叶子，那么值就应该是在这个节点的子节点里面
  //所有就开始寻找这个节点的子节点
  } else if x.Leaf == false {
    n, idx = x.Children[i].Search(k)
  }

  return
}

//make a new B+ tree, it is a struct 
func newNode(n, branch int, leaf bool) *Node {
  return &Node {
    Leaf:       leaf,
    N:          n,
    Keys:       make([]Key, branch*2 - 1),
    Children:   make([]*Node, branch*2),
  }
}


func (parent *Node) Split(branch, idx int) { //  idx is Children's index
	full := parent.Children[idx]
	// make a new node, copy full's right most to it
	n := newNode(branch-1, branch, full.Leaf)
	for i := 0; i < branch-1; i++ {
		n.Keys[i] = full.Keys[i+branch]
		n.Children[i] = full.Children[i+branch]
	}
	n.Children[branch-1] = full.Children[2*branch-1] // copy last child
	full.N = branch - 1 // is half full now, copied to n(new one)
	// shift parent, add new key and children
	for i := parent.N; i > idx; i-- {
		parent.Children[i] = parent.Children[i-1]
		parent.Keys[i+1] = parent.Keys[i]
	}
	parent.Keys[idx] = full.Keys[branch-1]
	parent.Children[idx+1] = n
	parent.N += 1
}
func (tree *Btree) Insert(k Key) {
	root := tree.Root
	if root.N == 2*tree.branch-1 {
		s := newNode(0, tree.branch, false)
		tree.Root = s
		s.Children[0] = root
		s.Split(tree.branch, 0)
		s.InsertNonFull(tree.branch, k)
	} else {
		root.InsertNonFull(tree.branch, k)
	}
}
func (x *Node) InsertNonFull(branch int, k Key) {
	i := x.N
	if x.Leaf {
		for i > 0 && k < x.Keys[i-1] {
			x.Keys[i] = x.Keys[i-1]
			i -= 1
		}
		x.Keys[i] = k
		x.N += 1
	} else {
		for i > 0 && k < x.Keys[i-1] {
			i -= 1
		}
		c := x.Children[i]
		if c.N == 2*branch-1 {
			x.Split(branch, i)
			if k > x.Keys[i] {
				i += 1
			}
		}
		x.Children[i].InsertNonFull(branch, k)
	}
}
func space(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += " "
	}
	return s
}
func (x *Node) String() string {
	return fmt.Sprintf("{n=%d, Leaf=%v, Keys=%v, Children=%v}\n",
		x.N, x.Leaf, x.Keys, x.Children)
}
func (tree *Btree) String() string {
	return tree.Root.String()
}
type Btree struct {
	Root   *Node
	branch int
}
func New(branch int) *Btree {
	return &Btree{Root: newNode(0, branch, true), branch: branch}
}
func (tree *Btree) Search(k Key) (n *Node, idx int) {
	return tree.Root.Search(k)
}











