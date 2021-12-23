package base

import (
	"embed"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

// mapData represents the internal data structure for map YAML files.
type mapData struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Dimension   struct {
		X int `yaml:"x"`
		Y int `yaml:"y"`
	} `yaml:"dimension"`
	Layers []struct {
		Name        string `yaml:"name"`
		Description string `yaml:"description"`
		Objects     string `yaml:"objects"`
	} `yaml:"layers"`
}

//go:embed assets/*
var assets embed.FS

// MustLoadMap loads a Map from the given YAML file and panics if it fails.
func MustLoadMap(file string) *Map {
	f, err := assets.Open(file)
	if err != nil {
		panic(err)
	}
	var md mapData
	if err := yaml.NewDecoder(f).Decode(&md); err != nil {
		panic(err)
	}
	dim := NewVertex(md.Dimension.X, md.Dimension.Y)
	ls := make([]*Layer, len(md.Layers))
	for idx, v := range md.Layers {
		ls[idx] = NewLayer(LayerName(v.Name), dim, loadLayerFromStr(dim, v.Objects))
	}

	return NewMap(MapName(md.Name), dim, ls)
}

// loadLayerFromStr loads Objects from a string.
// Cf. object_data.go
func loadLayerFromStr(dim Vertex, s string) []*Object {
	ss := strings.ReplaceAll(s, "\n", " ")
	ss = strings.Trim(ss, " ")
	objStrList := strings.Split(ss, " ")
	objList := make([]*Object, len(objStrList))
	for idx, objStr := range objStrList {
		v, err := strconv.Atoi(objStr)
		if err != nil {
			objList[idx] = NewObject(ObjUndef, VertexFromIndex(dim, idx))
		}
		objList[idx] = NewObject(ObjectType(v), VertexFromIndex(dim, idx))
	}
	return objList
}
