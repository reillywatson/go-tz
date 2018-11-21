package gotz

import (
	"encoding/json"
	"errors"
)

var errNoTZID = errors.New("tzid for feature not found")

type FeatureCollection struct {
	featureCollection
}

type featureCollection struct {
	Features []*Feature `json:"features"`
}

type Feature struct {
	feature
}

type feature struct {
	Geometry   Geometry `json:"geometry"`
	Properties struct {
		Tzid string `json:"tzid"`
	} `json:"properties"`
}

type Geometry struct {
	geometry
}

type geometry struct {
	Type        string    `json:"type"`
	Coordinates [][]Point `json:"coordinates,omitempty"`
}

var jPolyType struct {
	Type       string      `json:"type"`
	Geometries []*Geometry `json:"geometries,omitempty"`
}

var jPolygon struct {
	Coordinates [][][]float64 `json:"coordinates,omitempty"`
}

var jMultiPolygon struct {
	Coordinates [][][][]float64 `json:"coordinates,omitempty"`
}

func (g *Geometry) UnmarshalJSON(data []byte) (err error) {
	if err := json.Unmarshal(data, &jPolyType); err != nil {
		return err
	}
	g.Type = "MultiPolygon"

	if jPolyType.Type == "Polygon" {
		if err := json.Unmarshal(data, &jPolygon); err != nil {
			return err
		}
		//Create a bounding box
		pol := make([]Point, 0, 50)
		for _, v := range jPolygon.Coordinates[0] {
			pol = append(pol, Point{v[0], v[1]})
		}
		b := getBoundingBox(pol)
		g.Coordinates = append(g.Coordinates, b)
		g.Coordinates = append(g.Coordinates, pol)
		return nil
	}

	if jPolyType.Type == "MultiPolygon" {
		if err := json.Unmarshal(data, &jMultiPolygon); err != nil {
			return err
		}
		for _, poly := range jMultiPolygon.Coordinates {
			pol := make([]Point, 0, 50)
			for _, v := range poly[0] {
				pol = append(pol, Point{v[0], v[1]})
			}
			b := getBoundingBox(pol)
			g.Coordinates = append(g.Coordinates, b)
			g.Coordinates = append(g.Coordinates, pol)
		}
		return nil
	}
	return nil
}
