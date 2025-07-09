package tree

import (
	"fmt"
	"testing"
)

// 期待する結果
//
//	{
//	 A: 01,
//	 B: 00,
//	 C: 11,
//	 D: 1001,
//	 E: 1000,
//	 F: 101
//	}
func TestHuffman(t *testing.T) {
	hf := newHuffman()
	// input := `ABBCACCEACBCCFABCDAFEABFFADBBC`
	input := `すもももももももものうち`

	enc := hf.encode(input)
	fmt.Printf("encmap: %+v\n", hf.encmap)
	fmt.Printf("decmap: %+v\n", hf.decmap)
	fmt.Printf("encoded: %+v\n", enc)
	dec := hf.decode(string(enc))
	if dec != input {
		t.Errorf("\nNot Equal:\n{\ngot :%v\nwant:%v\n}\n", dec, input)
	}
	printCmp(input, string(enc))
}

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Add("word")
	trie.Add("wheel")
	trie.Add("world")
	trie.Add("hospital")
	trie.Add("mode")
	trie.ShowTrees()
	t.Log(trie.Contain("mo"))
}
