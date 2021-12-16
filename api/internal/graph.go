package internal

import (
	"bytes"
	"fmt"
	"goph-maps/models"
)

type Graph struct {
	Vertices []*Vertice `json:   "vertices"`

	directed bool
}

type Vertice struct {
	ID         int                        `json:   "id"`
	Point      models.Point               `json:   "point"`
	Ajdacentes map[int]*models.LineString `json:   "adjacentes"`
}

// graph constructor
func NewGraph(directed bool) *Graph {
	if directed {
		return &Graph{
			directed: true,
		}
	}
	return &Graph{}
}

// each vertex has his own identifier
var (
	id int = -1
)

// append a vertex to the graph
func (g *Graph) Add(point models.Point) {
	if !contains(g.Vertices, point.Name) {
		id++
		g.Vertices = append(g.Vertices, &Vertice{
			ID:         id,
			Point:      point,
			Ajdacentes: make(map[int]*models.LineString),
		})
	} else {
		return
	}
}

// add edge between two vertices
func (g *Graph) AddEdge(from, to int, road models.LineString) {
	v1 := g.Vertices[from]
	v2 := g.Vertices[to]

	if v1 == nil || v2 == nil {
		panic("error not all vertices exist")
	}

	if _, ok := v1.Ajdacentes[v2.ID]; ok {
		return
	}

	v1.Ajdacentes[v2.ID] = &road
	if !g.directed {
		v2.Ajdacentes[v1.ID] = &road
	}
}

// return a vertice from
func (g *Graph) GetVertexFromName(name string) *Vertice {
	for index, value := range g.Vertices {
		if value.Point.Name == name {
			return g.Vertices[index]
		}
	}
	return nil
}

func (g *Graph) GetVertexFromID(id int) *Vertice {
	for index, value := range g.Vertices {
		if value.Point.ID == id {
			return g.Vertices[index]
		}
	}
	return nil
}

func contains(points []*Vertice, name string) bool {
	for _, value := range points {
		if value.Point.Name == name {
			return true
		}
	}
	return false
}

func (g *Graph) Print() {
	for _, value := range g.Vertices {
		fmt.Println()
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
				g.GetVertexFromID(v).Point.Name,
				g.GetVertexFromID(w).Point.Name))
		}
	}
	return buffer.String()
}

func (g *Graph) Neighbors(id int) []int {
	var neighbors = make([]int, len(g.Vertices))
	for _, node := range g.Vertices {
		for edge := range node.Ajdacentes {
			if node.Point.ID == id {
				neighbors = append(neighbors, edge)
			}
			if edge == id {
				neighbors = append(neighbors, node.Point.ID)
			}
		}
	}
	return neighbors
}

type Edge struct {
	IDVertex     int
	IDAdjacente  int
	IDLineString string
}

func (g *Graph) Edges() []Edge {
	var edges = make([]Edge, 0, len(g.Vertices))
	for i := 0; i < len(g.Vertices); i++ {
		for k, v := range g.Vertices[i].Ajdacentes {
			edge := Edge{
				IDVertex:     i,
				IDAdjacente:  k,
				IDLineString: v.ID,
			}
			edges = append(edges, edge)
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
	fmt.Println("starting node name: ", startingNode.Point.Name)
	var res []Vertice
	var visited = make(map[string]bool)
	visited[startingNode.Point.Name] = true

	var queue = queue{}
	queue.enqueue(startingNode)

	for !queue.isEmpty() {
		s := queue.dequeue()

		fmt.Println("current node: ", s.Point.Name)

		if s.Point.Name == destination {
			return res
		}

		for index, value := range s.Ajdacentes {
			fmt.Println(value)
			if !visited[g.GetVertexFromID(index).Point.Name] {
				queue.enqueue(g.Vertices[index])
				res = append(res, *g.Vertices[index])
				visited[g.Vertices[index].Point.Name] = true
			}
		}
	}
	fmt.Println("visited: ", visited)
	return res
}
