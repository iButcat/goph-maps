package internal

import (
	"bytes"
	"fmt"
)

type Graph struct {
	Vertices []*Vertice

	directed bool
}

type Vertice struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Ajdacentes map[int]int
}

func NewGraph(directed bool) *Graph {
	if directed {
		return &Graph{
			directed: true,
		}
	}
	return &Graph{}
}

func (g *Graph) Add(name string) {
	id := len(g.Vertices)
	if !contains(g.Vertices, name) {
		g.Vertices = append(g.Vertices, &Vertice{
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
	v1 := g.Vertices[from]
	v2 := g.Vertices[to]

	if v1 == nil || v2 == nil {
		panic("error not all vertices exist")
	}

	if _, ok := v1.Ajdacentes[v2.ID]; ok {
		fmt.Println(ok)
		return
	}
	v1.Ajdacentes[v2.ID] = distance
	if !g.directed {
		v2.Ajdacentes[v1.ID] = distance
	}
}

func (g *Graph) GetVertexFromName(name string) *Vertice {
	for index, value := range g.Vertices {
		if value.Name == name {
			return g.Vertices[index]
		}
	}
	return nil
}

func (g *Graph) GetVertexFromID(id int) *Vertice {
	for index, value := range g.Vertices {
		if value.ID == id {
			return g.Vertices[index]
		}
	}
	return nil
}

func contains(points []*Vertice, name string) bool {
	for _, value := range points {
		if value.Name == name {
			return true
		}
	}
	return false
}

func (g *Graph) Print() {
	for _, value := range g.Vertices {
		fmt.Println("points: ", value)
		for _, value := range value.Ajdacentes {
			fmt.Println("adjacentes: ", value)
		}
	}
	fmt.Println()
}

func (g *Graph) String() string {
	var buffer bytes.Buffer
	for v, b := range g.Vertices {
		for w := range b.Ajdacentes {
			buffer.WriteString(fmt.Sprintf("%s - %s \n",
				g.GetVertexFromID(v).Name, g.GetVertexFromID(w).Name))
		}
	}
	return buffer.String()
}

func (g *Graph) Neighbors(id int) []int {
	var neighbors = make([]int, len(g.Vertices))
	for _, node := range g.Vertices {
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
	var edges = make([][3]int, 0, len(g.Vertices))
	for i := 0; i < len(g.Vertices); i++ {
		for k, v := range g.Vertices[i].Ajdacentes {
			edges = append(edges, [3]int{i, k, int(v)})
		}
	}
	return edges
}

func (g *Graph) Nodes() []int {
	nodes := make([]int, len(g.Vertices))
	for i := range g.Vertices {
		nodes[i] = i
	}
	return nodes
}

func (g *Graph) BFS(startingNode *Vertice, destination string) []Vertice {
	fmt.Println("starting node name: ", startingNode.Name)
	var res []Vertice
	var visited = make(map[string]bool)
	visited[startingNode.Name] = true

	var queue = queue{}
	queue.enqueue(startingNode)

	for !queue.isEmpty() {
		s := queue.dequeue()

		fmt.Println("current node: ", s.Name)

		if s.Name == destination {
			return res
		}

		for index, _ := range s.Ajdacentes {
			if !visited[g.GetVertexFromID(index).Name] {
				queue.enqueue(g.Vertices[index])
				res = append(res, *g.Vertices[index])
				visited[g.Vertices[index].Name] = true
			}
		}
	}
	fmt.Println("visited: ", visited)
	return res
}
