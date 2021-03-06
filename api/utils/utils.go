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

// Temporary to check content of the graph vertices
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
		} else if features[i].Geometry.IsLineString() {
			lineString := models.NewLineString(features[i].Properties["route_id"].(string),
				features[i].Properties["route_long_name"].(string),
				features[i].Geometry.LineString)
			sliceLineString = append(sliceLineString, lineString)
		}
	}

	for i := 0; i < len(Graph.Vertices); i++ {
		for j := i + 1; j < len(Graph.Vertices); j++ {
			link(&Graph.Vertices[i].Point, &Graph.Vertices[j].Point, sliceLineString)
		}
	}
	Graph.Print()
	log.Println(len(Graph.Edges()))
}

func link(point1, point2 *models.Point, sliceLineString []*models.LineString) {
	for _, lineString := range sliceLineString {
		for i := 0; i < len(lineString.Geometry); i++ {
			for j := i + 1; j < len(lineString.Geometry); j++ {
				if equal(lineString.Geometry[i], point1.Geometry) && equal(lineString.Geometry[j], point2.Geometry) {
					//log.Println(point1.Name, "is connected to: ", point2.Name, "by:", lineString.Name)
					Graph.AddEdge(Graph.GetVertexFromName(point1.Name).ID, Graph.GetVertexFromName(point2.Name).ID, *lineString)
				}
			}
		}
	}
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
