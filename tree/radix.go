package tree

import (
	"net/http"
	"strings"
)

type radixNode struct {
	value    string
	children []*radixNode
	isWild   bool
	handler  http.HandlerFunc
}

func NewRadix() *radixNode {
	return &radixNode{
		value:    "",
		children: make([]*radixNode, 0),
		isWild:   false,
		handler:  nil,
	}
}

func (r *radixNode) insert(pattern string, h http.HandlerFunc) {
	if pattern == "" || pattern[0] != '/' || h == nil {
		return
	}
	segs := strings.Split(pattern, "/")[1:]
	for _, key := range segs {
		child := r.matchChild(key)
		if child == nil {
			node := &radixNode{
				value:    key,
				children: make([]*radixNode, 0),
				isWild:   r.checkWildCard(key),
			}
			r.children = append(r.children, node)
			r = node
			continue
		}
		r = child
	}
	r.handler = h
}

func (r *radixNode) search(path string) http.HandlerFunc {
	if path == "" || path[0] != '/' {
		return nil
	}
	segs := strings.Split(path, "/")[1:]
	for _, key := range segs {
		child := r.matchChild(key)
		if child == nil {
			return nil
		}
		r = child
	}
	return r.handler
}

func (r *radixNode) matchChild(key string) *radixNode {
	for _, child := range r.children {
		if child.value == key || child.isWild {
			return child
		}
	}
	return nil
}

func (r *radixNode) checkWildCard(key string) bool {
	return key[0] == ':' || key[0] == '*' || (key[0] == '{' && key[len(key)-1] == '}')
}
