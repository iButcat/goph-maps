package utils

import (
	"fmt"
	"goph-maps/internal"
	"goph-maps/models"
	"io/ioutil"
	"log"
	"strconv"

	geojson "github.com/paulmach/go.geojson"
)

func readGeoJsonFile(fileName string) ([]byte, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func parseGeoJsonContent(fileName string) (*geojson.FeatureCollection, error) {
	geojsonContent, err := readGeoJsonFile(fileName)
	if err != nil {
		return nil, err
	}

	collectionFeatures, err := geojson.UnmarshalFeatureCollection(geojsonContent)
	if err != nil {
		panic(err)
	}

	return collectionFeatures, nil
}

var Graph = &internal.Graph{}

func GeoJsonToStruct(fileName string) {
	collectionFeatures, err := parseGeoJsonContent(fileName)
	if err != nil {
		panic(err)
	}

	features := collectionFeatures.Features
	Graph = internal.NewGraph(false)
	var sliceLineString []*models.LineString
	for i := 0; i < len(features); i++ {

		if features[i].Geometry.IsPoint() {
			convertIDToInt, err := strconv.Atoi(features[i].Properties["id"].(string))
			if err != nil {
				fmt.Println(err)
			}
			point := models.NewPoint(convertIDToInt,
				features[i].Properties["name"].(string),
				features[i].Geometry.Point)
			Graph.Add(*point)
			point = &models.Point{}
		}

		if features[i].Geometry.IsLineString() {
			lineString := models.NewLineString(features[i].Properties["route_id"].(string),
				features[i].Properties["route_long_name"].(string),
				features[i].Geometry.LineString)
			sliceLineString = append(sliceLineString, lineString)
			lineString = &models.LineString{}
		}
	}

	for i := 0; i < len(Graph.Vertices); i++ {
		for q := i + 1; q < len(Graph.Vertices); q++ {
			for j := 0; j < len(sliceLineString); j++ {
				ele := getFirstAndLastCoordinates(sliceLineString[j].Geometry)

				if equal(ele[0], Graph.Vertices[i].Point.Geometry) && equal(ele[1], Graph.Vertices[q].Point.Geometry) {
					if Graph.Vertices[q].ID > len(Graph.Vertices) {
						log.Println("problem with index")
					} else {
						Graph.AddEdge(Graph.Vertices[i].ID, Graph.Vertices[q].ID, *sliceLineString[j])
					}
				}
			}
		}
	}
}

func AddEdges() {

}

// Check if both slice are equal
func equal(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// First array and last are the coordinate of a point.
func getFirstAndLastCoordinates(lineStringGeometry [][]float64) [][]float64 {
	var coordinates [][]float64
	coordinates = append(coordinates, lineStringGeometry[0])
	coordinates = append(coordinates, lineStringGeometry[len(lineStringGeometry)-1])
	return coordinates
}

func checkifCoordinatesForEdges(lineStringGeometry [][]int, pointGeometry []int) bool {
	for _, value := range lineStringGeometry {
		for _, val := range value {
			for _, v := range pointGeometry {
				if val == v {
					return true
				}
			}
		}
	}
	return false
}
