package main

import "fmt"

/**
G = <V, E>
V = struct, class
E = Matrix or Map<Vertex, *ListNode>
边的表示方法： 邻接矩阵， 邻接表
*/

type Vertex struct {
	label byte
	is_visited bool
}

type stack struct {
	data []int
	top int
	size int
}

func new_stack() stack {
	return stack{}
}
func stack_push(s stack, v int) {

}
func stack_peek(s stack) int {
	return -1
}
func stack_pop(s stack) (int, bool){
	return -1, false
}	

type Queue struct {
	data []int
	head, tail int
	size int
}
func new_queue() Queue {
	return Queue{}
}
func queue_push(s Queue, val int) {

}
func queue_pop(s Queue) (int, bool){
	return -1, false
}

type Graph struct {
	max_vertex_nums int
	vertex_array []Vertex
	adj_matrix [][]int
	vertex_len int
	stack stack
	queue Queue
}

func create_graph(max int) *Graph {
	var g *Graph = &Graph{
		max_vertex_nums: max,
		vertex_array: make([]Vertex, max, max),
		adj_matrix: make([][]int, max, max),
		vertex_len: 0,
		stack: new_stack(),
		queue: new_queue(),
	}
	for i := 0; i < max; i++ {
		g.adj_matrix[i] = make([]int, max, max)
	}
	return g
}

func add_vertex(g *Graph, label byte) bool{
	if g.vertex_len >= g.max_vertex_nums {
		return false
	}
	g.vertex_array[g.vertex_len] = Vertex{label: label, is_visited: false}
	g.vertex_len++
	return true
}

func add_edge(g *Graph, start, end int) {
	g.adj_matrix[start][end] = 1
	g.adj_matrix[end][start] = 1
}

func print_vertex(g *Graph, v int) {
	fmt.Printf("vertex=%v\n", g.vertex_array[v].label)
}
func get_adj_vertex(g *Graph, v int) int {
	for i := 0; i < g.max_vertex_nums; i++ {
		if g.adj_matrix[v][i] == 1 && g.vertex_array[i].is_visited == false {
			return i
		}
	}
	return -1
}

func dfs_visit_graph(g *Graph) {
	first := 0
	g.vertex_array[first].is_visited = true
	print_vertex(g, first)
	stack_push(g.stack, first)
	for g.stack.size != 0 {
		cur := stack_peek(g.stack)
		adj := get_adj_vertex(g, cur)
		if adj != -1 {
			// exit adj vertex
			g.vertex_array[adj].is_visited = true
			print_vertex(g, adj)
			stack_push(g.stack, adj)
		} else {
			stack_pop(g.stack)
		}
	}
	// restore
	for i := 0; i < g.max_vertex_nums; i++ {
		g.vertex_array[i].is_visited = false
	}
}

func bfs_visit_graph(g *Graph) {
	first := 0

	queue_push(g.queue, first)
	for g.queue.size != 0 {
		cur, _ := queue_pop(g.queue)
		g.vertex_array[cur].is_visited = true
		print_vertex(g, cur)
		for adj := get_adj_vertex(g, cur); adj != -1; adj = get_adj_vertex(g, cur){
			if adj == -1 {
				break
			}
			queue_push(g.queue, adj)
		}
	}
	// restore
	for i := 0; i < g.max_vertex_nums; i++ {
		g.vertex_array[i].is_visited = false
	}
}

func min_spanning_tree(g Graph) {
	//todo
}

func main() {

}