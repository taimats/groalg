package graph

import (
	"maps"
	"slices"
)

type cost int

func (c cost) isEmpty() bool {
	return c == -1
}

type dijkNode struct {
	name  string
	costs map[string]cost
}

type dijkstra struct {
	graph    map[string][]*dijkNode
	costs    map[string]cost
	parents  map[string]string
	searched map[string]struct{}
}

func newDijkstra(graph map[string][]*dijkNode) *dijkstra {
	costs := make(map[string]cost)
	parents := make(map[string]string)
	for node := range graph {
		costs[node] = cost(-1)
		parents[node] = ""
	}
	searched := make(map[string]struct{})

	return &dijkstra{
		graph:    graph,
		costs:    costs,
		parents:  parents,
		searched: searched,
	}
}

func (d *dijkstra) shortest(start *dijkNode, end *dijkNode) (route []string, cost int) {
	//start地点の初期化処理
	nodes := d.graph[start.name]
	d.costs[start.name] = 0
	d.addSearched(start.name)
	//start地点から進めるnodeの各コストと親を保持。
	for _, n := range nodes {
		d.costs[n.name] = n.costs[start.name]
		d.parents[n.name] = start.name
	}
	//探索処理
	for len(d.searched) < len(d.graph) {
		node := d.minCostNode()
		cost := d.costs[node]
		neighbors := d.graph[node]
		for _, nbr := range neighbors {
			if d.costs[nbr.name].isEmpty() {
				d.costs[nbr.name] = d.costs[node] + nbr.costs[node]
			}
			if d.parents[nbr.name] == "" {
				d.parents[nbr.name] = node
			}
			newCost := cost + nbr.costs[node]
			if d.costs[nbr.name] > newCost {
				d.costs[nbr.name] = newCost
				d.parents[nbr.name] = node
			}
		}
		d.addSearched(node)
	}
	route = d.route(start.name, end.name)
	cost = int(d.costs[end.name])
	return
}

func (d *dijkstra) minCostNode() (node string) {
	//初期化時(-1)以外のcostの一覧を生成。
	vals := make(map[cost]string)
	for node, cost := range d.costs {
		if !d.isSearched(node) && !cost.isEmpty() {
			vals[cost] = node
		}
	}
	//昇順にソートして最小値(スライスの先頭)を取得
	min := slices.Sorted(maps.Keys(vals))[0]
	return vals[min]
}

func (d *dijkstra) addSearched(node string) {
	d.searched[node] = struct{}{}
}

func (d *dijkstra) isSearched(node string) bool {
	_, exists := d.searched[node]
	return exists
}

// endからたどり、最短経路を配列で返す。
func (d *dijkstra) route(start string, end string) (route []string) {
	route = make([]string, 0, len(d.parents))
	route = append(route, end)
	parent := d.parents[end]
	route = append(route, parent)
	for parent != start {
		parent = d.parents[parent]
		route = append(route, parent)
	}
	slices.Reverse(route)
	return route
}
