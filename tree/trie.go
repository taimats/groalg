package tree

import "fmt"

func NewTrie() *TrieNode {
	return &TrieNode{
		prefix:   0,
		children: make([]*TrieNode, 0),
	}
}

type TrieNode struct {
	prefix   rune
	children []*TrieNode
}

func newNode(prefix rune) *TrieNode {
	return &TrieNode{
		prefix:   prefix,
		children: make([]*TrieNode, 0),
	}
}

func (t *TrieNode) Add(word string) {
	for _, r := range word {
		child := t.matchChild(r)
		if child == nil {
			node := newNode(r)
			t.insert(node)
			t = node
			continue
		}
		t = child
	}
}

func (t *TrieNode) Contain(key string) bool {
	if len(key) == 0 {
		return false
	}
	for _, r := range key {
		child := t.matchChild(r)
		if child == nil {
			return false
		}
		t = child
	}
	return true
}

func (t *TrieNode) ShowTrees() {
	if t.prefix == 0 && len(t.children) == 0 {
		fmt.Println("no trees")
	}
	depth := 1
	t.showTrees(depth, t.children)
}

func (t *TrieNode) showTrees(depth int, children []*TrieNode) {
	length := len(children)
	if length == 0 {
		return
	}
	nextDepth := make([]*TrieNode, 0)
	fmt.Printf("===== %d 階層 =====\n", depth)
	for i, node := range children {
		fmt.Printf("{%s}", string(node.prefix))
		if i == length-1 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
		nextDepth = append(nextDepth, node.children...)
	}
	fmt.Println()
	t.showTrees(depth+1, nextDepth)
}

func (t *TrieNode) matchChild(prefix rune) *TrieNode {
	if len(t.children) == 0 {
		return nil
	}
	for _, child := range t.children {
		if child.prefix == prefix {
			return child
		}
	}
	return nil
}

func (t *TrieNode) insert(node *TrieNode) {
	t.children = append(t.children, node)
}
