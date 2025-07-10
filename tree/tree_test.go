package tree

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
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

func TestRadixSearch(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		handler http.HandlerFunc
		path    string
		want    http.HandlerFunc
	}{
		{
			"success:case-1",
			"/hello",
			index,
			"/hello",
			index,
		},
		{
			"success:case-2",
			"/foo",
			index,
			"/foo",
			index,
		},
		{
			"success:case-3",
			"/hello/:name",
			hello,
			"/hello/taro",
			hello,
		},
		{
			"success:case-4",
			"/hello/:name/foo",
			hello2,
			"/hello/taro/foo",
			hello2,
		},
	}
	sut := NewRadix()
	for _, tt := range tests {
		sut.insert(tt.pattern, tt.handler)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := sut.search(tt.path)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Not equal: %s", diff)
			}
		})
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/hello/")
	fmt.Fprintf(w, "Hello, %s!\n", name)
}

func hello2(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/hello/")
	fmt.Fprintf(w, "Hello2, %s!\n", name)
}
