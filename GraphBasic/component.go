package main

import (
	"container/list"
	"fmt"
)

type component struct {
	g     graph
	count int   // 子图的数量
	id    []int // 记录每个节点属于的子图
}

func newComponent(g graph) *component {
	var c = component{
		g:     g,
		count: 0,
		id:    make([]int, g.V()),
	}
	for i := 0; i < g.V(); i++ {
		c.id[i] = -1
	}
	return &c
}

// 从每一个节点开始进行深度优先遍历，并统计子图的个数
func (c *component) DFS() {
	var visited = make([]bool, c.g.V())
	fmt.Println("++++++++++++++Start dfs...")
	for i := 0; i < c.g.V(); i++ {
		if !visited[i] {
			fmt.Println("subgraph ", c.count)
			c.dfs(i, visited)
			c.count++
			fmt.Println()
		}
	}
	fmt.Println("subgraph number: ", c.count)
	fmt.Println("++++++++++++++End dfs...")
}

func (c *component) dfs(v int, visited []bool) {
	fmt.Print(v, " ")
	visited[v] = true
	c.id[v] = c.count
	iter := c.g.iter(v)
	for s := iter.begin(); !iter.end(); s = iter.next() {
		if !visited[s] {
			c.dfs(s, visited)
		}
	}
}

// 从每一个节点开始进行广度优先遍历, 并统计子图的个数
func (c *component) BFS() {
	var visited = make([]bool, c.g.V())
	fmt.Println("--------------Start bfs...")
	for i := 0; i < c.g.V(); i++ {
		if !visited[i] {
			fmt.Println("subgraph ", c.count)
			c.bfs(i, visited)
			c.count++
			fmt.Println()
		}
	}
	fmt.Println("subgraph number: ", c.count)
	fmt.Println("--------------End bfs...")
}

func (c *component) bfs(v int, visited []bool) {
	l := list.New()
	l.PushBack(v)
	visited[v] = true
	for e := l.Front(); e != nil; e = l.Front() {
		val := e.Value.(int)
		l.Remove(e)
		fmt.Print(val, " ")
		iter := c.g.iter(val)
		for s := iter.begin(); !iter.end(); s = iter.next() {
			if !visited[s] {
				l.PushBack(s)
				visited[s] = true
			}
		}
	}
}

func (c *component) isConnected(v, w int) bool {
	return c.id[v] == c.id[w]
}

// 利用深度优先搜索查找某一节点到其他节点的路径
type Path struct {
	g       graph
	node    int
	visited []bool
	from    []int
}

func newPath(g graph, v int) *Path {
	var p = Path{
		g:       g,
		node:    v,
		visited: make([]bool, g.V()),
		from:    make([]int, g.V()),
	}
	for i := 0; i < g.V(); i++ {
		p.from[i] = -1
	}
	p.dfs(v)
	return &p
}

func (p *Path) dfs(v int) {
	p.visited[v] = true
	i := p.g.iter(v)
	for j := i.begin(); !i.end(); j = i.next() {
		if !p.visited[j] {
			p.from[j] = v
			p.dfs(j)
		}
	}
}

func (p *Path) hasPath(w int) bool {
	return p.visited[w]
}

func (p *Path) showPath(w int) {
	if !p.hasPath(w) {
		fmt.Println("no path from", p.node, "to", w)
		return
	}
	fmt.Println("DFS showPath,", p.node, "to", w)
	var tmp = w
	var path []int
	for tmp != -1 {
		path = append(path, tmp)
		tmp = p.from[tmp]
	}
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Print(path[i])
		if i == 0 {
			fmt.Println()
		} else {
			fmt.Print(" -> ")
		}
	}
}

// 利用广度优先搜索查找某一节点到其他节点的最短路径
type ShortestPath struct {
	g       graph
	node    int
	from    []int
	visited []bool
	ord     []int
}

func newShortestPath(g graph, v int) *ShortestPath {
	var p = ShortestPath{
		g:       g,
		node:    v,
		from:    make([]int, g.V()),
		visited: make([]bool, g.V()),
		ord:     make([]int, g.V()),
	}

	for i := 0; i < g.V(); i++ {
		p.from[i] = -1
	}

	p.bfs(v)

	return &p
}

func (p *ShortestPath) bfs(v int) {
	p.visited[v] = true
	l := list.New()
	l.PushBack(v)

	for e := l.Front(); e != nil; e = l.Front() {
		val := e.Value.(int)
		l.Remove(e)

		i := p.g.iter(val)
		for j := i.begin(); !i.end(); j = i.next() {
			if !p.visited[j] {
				l.PushBack(j)
				p.visited[j] = true
				p.from[j] = val
				p.ord[j] = p.ord[val] + 1

			}
		}
	}
}

func (p *ShortestPath) hasPath(w int) bool {
	return p.visited[w]
}

func (p *ShortestPath) showPath(w int) {
	if !p.hasPath(w) {
		fmt.Println("no path from", p.node, "to", w)
		return
	}
	fmt.Println("BFS showPath,", p.node, "to", w)
	var tmp = w
	var path []int
	for tmp != -1 {
		path = append(path, tmp)
		tmp = p.from[tmp]
	}
	for i := len(path) - 1; i >= 0; i-- {
		fmt.Print(path[i])
		if i == 0 {
			fmt.Println()
		} else {
			fmt.Print(" -> ")
		}
	}
}

func (p *ShortestPath) length(w int) int {
	return p.ord[w]
}
