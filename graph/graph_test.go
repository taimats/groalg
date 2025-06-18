package graph

import (
	"fmt"
	"testing"
)

func TestBFS(t *testing.T) {
	g := make(graph[string, string])
	g["cab"] = []string{"car", "cat"}
	g["cat"] = []string{"mat", "bat"}
	g["car"] = []string{"cat", "bar"}
	g["mat"] = []string{"bat"}
	g["bar"] = []string{"bat"}
	want := 2
	got := bfs("cab", "bat", g)
	if got != want {
		t.Errorf("Not Equal: (got: %d, want: %d)", got, want)
	}
}

func TestWalkDir(t *testing.T) {
	path := ".."
	files, err := walkDir(path)
	if err != nil {
		t.Errorf("error should be nil: (error: %v)", err)
	}
	for _, f := range files {
		fmt.Println(f)
	}
}
