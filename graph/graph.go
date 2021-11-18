package main

import "fmt"

type Graph struct {
	vertices []*Node
}

type Node struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Ajdacentes map[int]int
}

func NewGraph() *Graph {
	return &Graph{}
}

func (g *Graph) Add(name string) {
	id := len(g.vertices)
	if !contains(g.vertices, name) {
		g.vertices = append(g.vertices, &Node{
			ID:         id,
			Name:       name,
			Ajdacentes: make(map[int]int),
		})
	} else {
		fmt.Println("Already exists")
		return
	}
}

func (g *Graph) AddEdge(from, to, distance int) {
	g.vertices[from].Ajdacentes[to] = distance
}

func (g *Graph) GetVertexFromName(name string) *Node {
	for index, value := range g.vertices {
		if value.Name == name {
			return g.vertices[index]
		}
	}
	return nil
}

func (g *Graph) GetVertexFromID(id int) *Node {
	for index, value := range g.vertices {
		if value.ID == id {
			return g.vertices[index]
		}
	}
	return nil
}

func contains(points []*Node, name string) bool {
	for _, value := range points {
		if value.Name == name {
			return true
		}
	}
	return false
}

func (g *Graph) Print() {
	for _, value := range g.vertices {
		fmt.Println("points: ", value)
		for _, value := range value.Ajdacentes {
			fmt.Println("adjacentes: ", value)
		}
	}
	fmt.Println()
}

func (g *Graph) Neighbors(id int) []int {
	var neighbors = make([]int, len(g.vertices))
	for _, node := range g.vertices {
		for edge := range node.Ajdacentes {
			if node.ID == id {
				neighbors = append(neighbors, edge)
			}
			if edge == id {
				neighbors = append(neighbors, node.ID)
			}
		}
	}
	return neighbors
}

func (g *Graph) Edges() [][3]int {
	var edges = make([][3]int, 0, len(g.vertices))
	for i := 0; i < len(g.vertices); i++ {
		for k, v := range g.vertices[i].Ajdacentes {
			edges = append(edges, [3]int{i, k, int(v)})
		}
	}
	return edges
}

func (g *Graph) Nodes() []int {
	nodes := make([]int, len(g.vertices))
	for i := range g.vertices {
		nodes[i] = i
	}
	return nodes
}

func (g *Graph) DFS() {
	var queue = make([]*Node, 0)
	for _, value := range g.vertices {
		fmt.Println("node: ", value.Name, "has been pushed")
		queue = append(queue, value)
	}
}

func main() {
	graph := NewGraph()

	graph.Add("Val de Briey")
	graph.Add("Amnéville")
	graph.Add("Luxembourg")
	graph.Add("Amsterdam")
	graph.Add("New York")
	graph.Add("Rotterdam")
	graph.Add("Dudelange")

	graph.AddEdge(0, 1, 30)
	graph.AddEdge(1, 2, 100)
	graph.AddEdge(2, 3, 200)
	graph.AddEdge(3, 4, 400)
	graph.AddEdge(4, 5, 440)
	graph.AddEdge(6, 8, 800)

	//fmt.Println("node 2 adjacentes: ", graph.Neighbors(2))

	//fmt.Println("edges: ", graph.Edges())
	//allID := graph.Nodes()
	//fmt.Println("all nodes id: ", allID)

	/*
		edges := graph.Edges()

		for i := 0; i < len(edges)-1; i++ {
			fmt.Println("city: ", graph.GetVertexFromID(edges[i][0]).Name,
				"is connected to: ", graph.GetVertexFromID(edges[i][1]).Name,
				"with a distance of: ", edges[i][2])
		}
	*/

	graph.DFS()

	graph.Print()
}
