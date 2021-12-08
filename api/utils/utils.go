package utils

import (
	"fmt"
	"goph-maps/models"
	"io/ioutil"

	geojson "github.com/paulmach/go.geojson"
)

func readGeoJsonFile(fileName string) ([]byte, error) {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func FileContentToStruct(fileName string) error {
	geojsonContent, err := readGeoJsonFile(fileName)
	if err != nil {
		return err
	}

	collectionFeatures, err := geojson.UnmarshalFeatureCollection(geojsonContent)
	if err != nil {
		panic(err)
	}

	fmt.Println("features collections type: ", collectionFeatures.Type)
	features := collectionFeatures.Features
	for i := 0; i < len(features); i++ {
		if features[i].Geometry.IsPoint() {
			fmt.Println("----------------------")
			point := models.NewPoint(features[i].Properties["id"].(string),
				features[i].Properties["name"].(string),
				features[i].Geometry.Point)
			fmt.Println("point: ", point)
			fmt.Println("----------------------")
			point = &models.Point{}
		}

		if features[i].Geometry.IsLineString() {
			fmt.Println("----------------------")
			lineString := models.NewLineString(features[i].Properties["route_id"].(string),
				features[i].Properties["route_long_name"].(string),
				features[i].Geometry.LineString)
			fmt.Println("line string: ", lineString)
			fmt.Println("----------------------")
			lineString = &models.LineString{}
		}
	}

	return nil
}

/*
func smh() {
	for i := 0; i < len(features); i++ {
		for j := i + 1; j < len(features); j++ {

			if features[i].Geometry.IsPoint() {
				point.New(features[i])
				fmt.Println("point: ", point)
				break
			}

			if features[i].Geometry.IsPoint() && features[j].Geometry.IsLineString() &&
				features[j].Geometry.LineString[0][0] == features[i].Geometry.Point[0] {
				fmt.Println("YES")
			}
		}
	}
}
*/
