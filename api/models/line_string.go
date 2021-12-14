package models

import (
	"encoding/json"
	"fmt"
)

type LineString struct {
	ID       string      `json:"route_id"`
	Name     string      `json:"name"`
	Geometry [][]float64 `json:"geometry"`
}

func NewLineString(id, name string, geometry [][]float64) *LineString {
	return &LineString{
		ID:       id,
		Name:     name,
		Geometry: geometry,
	}
}

func (l *LineString) Print() {
	fmt.Println("lineString: ", l)
}

func (l *LineString) String() string {
	str, err := json.Marshal(l)
	if err != nil {
		panic(err)
	}
	return string(str)
}
