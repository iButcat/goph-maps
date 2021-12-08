package models

type Point struct {
	ID       string    `json:   "id"`
	Name     string    `json:   "name"`
	Geometry []float64 `json:"geometry"`
}

func NewPoint(id, name string, geometry []float64) *Point {
	return &Point{
		ID:       id,
		Name:     name,
		Geometry: geometry,
	}
}
