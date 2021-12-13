package models

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	ID       int       `json:   "id"`
	Name     string    `json:   "name"`
	Geometry []float64 `json:"geometry"`
}

func NewPoint(id int, name string, geometry []float64) *Point {
	return &Point{
		ID:       id,
		Name:     name,
		Geometry: geometry,
	}
}

func (p *Point) Print() {
	fmt.Println("point: ", p)
}

func (p *Point) String() string {
	str, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return string(str)
}
