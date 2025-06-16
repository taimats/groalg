package tree

import (
	"fmt"
	"testing"
)

func TestHuffman(t *testing.T) {
	text := `ABBCACCEACBCCFABCDAFEABFFADBBC`
	c := newChars()
	c.countChars(text)
	for k, v := range c.m {
		fmt.Printf("{\n  文字:%s\n  頻度:%d\n}\n", string(k), v)
	}
	leaves := nodes(*c)
	ht := huffman(leaves)
	fmt.Printf("%+v", ht)
}
