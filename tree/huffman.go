package tree

import (
	"bytes"
	"cmp"
	"encoding/binary"
	"fmt"
	"slices"
	"strings"
)

type huffman struct {
	tree   *node
	encmap map[rune][]byte
	decmap map[string]rune
}

func newHuffman() *huffman {
	return &huffman{
		tree:   nil,
		encmap: make(map[rune][]byte),
		decmap: make(map[string]rune),
	}
}

type node struct {
	parent *node
	left   *node
	right  *node
	count  int
	value  rune
}

func (h *huffman) encode(input string) []byte {
	if input == "" {
		return nil
	}

	freq := calcFrequency(input)
	nodes := make([]*node, 0, len(freq))
	for value, count := range freq {
		n := &node{
			count: count,
			value: value,
		}
		nodes = append(nodes, n)
	}
	h.tree = h.trees(nodes)
	h.tables(h.tree, 0, 0)

	var encoded bytes.Buffer
	for _, r := range input {
		encoded.Write(h.encmap[r])
	}
	return encoded.Bytes()
}

func (h *huffman) trees(nodes []*node) *node {
	if len(nodes) < 2 {
		return nil
	}
	tmp := make([]*node, 0, 2)
	for len(nodes) >= 2 {
		//nodeスライスからcountが最小のnodeを2つ取り出す。
		//tmpに入れたら、取り出したnodeをnodeスライスから削除。
		for len(tmp) < 2 {
			min := slices.MinFunc(nodes, func(a, b *node) int {
				return cmp.Compare(a.count, b.count)
			})
			if min.parent != nil {
				continue
			}
			tmp = append(tmp, min)
			nodes = slices.DeleteFunc(nodes, func(n *node) bool {
				return n.value == min.value && n.count == min.count
			})
		}
		left, right := tmp[0], tmp[1]
		parent := &node{
			left:  left,
			right: right,
			count: left.count + right.count,
		}
		left.parent = parent
		right.parent = parent

		nodes = append(nodes, parent)
		tmp = slices.Delete(tmp, 0, len(tmp))
	}
	return nodes[0]
}

func (h *huffman) tables(n *node, r uint64, bits byte) {
	if n.value != 0 {
		buf := make([]byte, binary.MaxVarintLen64)
		if n.parent.right == n {
			r |= 1 << uint64(bits)
			i := binary.PutUvarint(buf, r)
			h.encmap[n.value] = buf[:i]
			h.decmap[string(buf[:i])] = n.value
		}
		if n.parent.left == n {
			i := binary.PutUvarint(buf, r)
			h.encmap[n.value] = buf[:i]
			h.decmap[string(buf[:i])] = n.value
		}
		bits++
		return
	}
	bits++
	h.tables(n.right, r, bits)
	h.tables(n.left, r, bits)
}

func (h *huffman) decode(encoded string) string {
	var decoded bytes.Buffer
	var buf bytes.Buffer
	for _, b := range []byte(encoded) {
		buf.WriteByte(b)
		if s, ok := h.decmap[buf.String()]; ok {
			decoded.WriteRune(s)
			buf.Reset()
		}
	}
	return decoded.String()
}

// textを1文字単位(rune)で探索し、文字ごとの登場頻度を算出する。
func calcFrequency(text string) map[rune]int {
	frequency := make(map[rune]int)
	//探索の重複を防ぐためにindexを設定。
	index := make(map[rune]struct{})
	for _, char := range text {
		if _, exists := index[char]; exists {
			continue
		}
		frequency[char] = strings.Count(text, string(char))
		index[char] = struct{}{}
	}
	return frequency
}

func printCmp(before, after string) {
	format := fmt.Sprintf(
		"{\n  元のサイズ: %d\n  圧縮後のサイズ: %d\n}",
		len([]byte(before)),
		len([]byte(after)),
	)
	fmt.Println(format)
}
