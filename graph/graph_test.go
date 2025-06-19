package graph

import (
	"fmt"
	"slices"
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

func TestDijkstra(t *testing.T) {
	s := &dijkNode{"start", nil}
	a := &dijkNode{"a", map[string]cost{"start": cost(5), "b": cost(8)}}
	b := &dijkNode{"b", map[string]cost{"start": cost(2)}}
	c := &dijkNode{"c", map[string]cost{"a": cost(4)}}
	d := &dijkNode{"d", map[string]cost{"a": cost(2), "b": cost(7), "c": cost(6)}}
	e := &dijkNode{"end", map[string]cost{"c": cost(3), "d": cost(1)}}
	g := map[string][]*dijkNode{
		s.name: {a, b},
		a.name: {c, d},
		b.name: {a, d},
		c.name: {d, e},
		d.name: {e},
		e.name: nil,
	}
	wantRoute := []string{s.name, a.name, d.name, e.name}
	wantCost := 8
	dk := newDijkstra(g)

	route, cost := dk.shortest(s, e)

	if !slices.Equal(route, wantRoute) {
		t.Errorf("route is wrong: (got: %+v, want: %+v)", route, wantRoute)
	}
	if cost != wantCost {
		t.Errorf("cost is wrong: (got: %+v, want: %+v)", cost, wantCost)
	}
}
