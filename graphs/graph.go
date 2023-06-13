package main

import (
	"fmt"
)

type Graph struct {
	vertices map[int][]int
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(map[int][]int),
	}
}

func (g *Graph) AddVertex(vertex int) {
	if _, ok := g.vertices[vertex]; !ok {
		g.vertices[vertex] = []int{}
	}
}

func (g *Graph) AddEdge(from, to int) {
	g.vertices[from] = append(g.vertices[from], to)
	g.vertices[to] = append(g.vertices[to], from) // Comment this line for a directed graph
}

func (g *Graph) PrintGraph() {
	for vertex, neighbors := range g.vertices {
		fmt.Printf("Vertex %d: ", vertex)
		for _, neighbor := range neighbors {
			fmt.Printf("%d ", neighbor)
		}
		fmt.Println()
	}
}

func main() {
	graph := NewGraph()

	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)

	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)
	graph.AddEdge(3, 4)
	graph.AddEdge(4, 1)

	graph.PrintGraph()
}