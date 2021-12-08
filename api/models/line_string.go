package models

type LineString struct {
	ID       string      `json:"id"`
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
