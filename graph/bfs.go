package graph

func bfs(start string, end string, g graph[string, string]) int {
	q := &queue{}
	q.enque(g[start]...)

	list := newSearched()
	steps := make(map[string]int)
	steps[start] = 0

	for {
		if q.isEmpty() {
			break
		}
		top := q.deque()
		names := g[top]
		for _, n := range names {
			if list.isSearched(n) {
				continue
			}
			steps[n] = steps[top] + 1
			list.add(n)
			q.enque(g[n]...)
		}
	}
	return steps[end]
}

type graph[K, V comparable] map[K][]V

type queue struct {
	data []string
	size int
}

func (q *queue) enque(s ...string) {
	q.data = append(q.data, s...)
	q.size++
}
func (q *queue) deque() string {
	if q.isEmpty() {
		return ""
	}
	top := q.data[0]
	q.data[0] = "" //prevention of memory leak

	q.data = q.data[1:]
	q.size--

	return top
}

func (q *queue) isEmpty() bool {
	return q.size == 0
}

type searched struct {
	m map[string]struct{}
}

func newSearched() *searched {
	return &searched{m: make(map[string]struct{})}
}

func (s *searched) isSearched(name string) bool {
	_, exists := s.m[name]
	return exists
}

func (s *searched) add(name string) {
	s.m[name] = struct{}{}
}
