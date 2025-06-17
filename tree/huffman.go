package tree

import (
	"fmt"
	"sort"
)

type node struct {
	parent *node
	left   *node
	right  *node
	count  freq
	char   rune
}

func nodes(c chars) []*node {
	nodes := make([]*node, 0, len(c.m))
	for k, v := range c.m {
		n := &node{
			count: v,
			char:  k,
		}
		nodes = append(nodes, n)
	}
	return nodes
}

func huffman(leaves []*node) *node {
	if len(leaves) < 1 {
		return nil
	}

	for len(leaves) > 1 {
		left, right := leaves[0], leaves[1]
		parent := &node{
			left:  left,
			right: right,
			count: left.count + right.count,
		}
		left.parent = parent
		right.parent = parent

		//降順になるようにparentを配置する
		rest := leaves[2:]
		index := sort.Search(len(rest), func(i int) bool {
			return rest[i].count >= parent.count
		})
		index += 2

		copy(leaves[1:], leaves[2:index])
		leaves[index-1] = parent
		leaves = leaves[1:]
	}
	return leaves[0]
}

// frequency
type freq int

// ある文字列における文字ごとの頻出数を保持する構造体。
type chars struct {
	m map[rune]freq
}

func newChars() *chars {
	m := make(map[rune]freq)
	return &chars{m: m}
}

// textを1文字単位で探索し、文字ごとの頻度を算出する。
func (c *chars) countChars(text string) {
	r := []rune(text)
	index := newIndex(len(r))
	for len(index.m) != 0 {
		for _, chr := range r {
			//すでにカウントした文字列の要素番号をindexから削除することで、
			//すべての文字が全探索を行うことを防いでいる。
			for i := range index.m {
				if chr == r[i] {
					c.m[chr] += 1
					index.delete(i)
				}
			}
		}
	}
}

type index struct {
	m map[int]struct{}
}

// 0~要素数(length)までの数字をすべて含んだmap。
func newIndex(length int) *index {
	m := make(map[int]struct{})
	i := 0
	for i < length {
		m[i] = struct{}{}
		i++
	}
	return &index{m: m}
}

func (i *index) delete(key int) {
	delete(i.m, key)
}

func printCmp(before, after string) {
	diffLength := fmt.Sprintf(
		"{\n  元の文字数: %d\n  圧縮後: %d\n}",
		len([]rune(before)),
		len([]rune(after)),
	)

	diffSize := fmt.Sprintf(
		"{\n  元のサイズ: %d\n  圧縮後: %d\n}",
		len([]byte(before)),
		len([]byte(after)),
	)

	fmt.Println(diffLength)
	fmt.Println(diffSize)
}
