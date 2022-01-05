package base

import (
	"embed"
	"fmt"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/ebiiim/fantasy/log"
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

// MustLoadMap loads a Map from the given YAML file or log.Fatal.
func MustLoadMap(file string) *Map {
	f, err := assets.Open(file)
	if err != nil {
		lg.Fatal(log.TypeInit, "MustLoadMap", "", "load map err=", err)
	}
	var md mapData
	if err := yaml.NewDecoder(f).Decode(&md); err != nil {
		lg.Fatal(log.TypeInit, "MustLoadMap", "", "decode map err=", err)
	}
	dim := NewVertex(md.Dimension.X, md.Dimension.Y)
	ls := make([]*Layer, len(md.Layers))
	for idx, v := range md.Layers {
		ls[idx] = NewLayer(LayerName(v.Name), dim, loadBaseObjectsFromStr(dim, v.Objects))
	}

	return NewMap(MapName(md.Name), dim, ls)
}

// loadBaseObjectsFromStr loads BaseObjects from a string.
// Cf. object_data.go
func loadBaseObjectsFromStr(dim Vertex, s string) []Object {
	ss := strings.ReplaceAll(s, "\n", " ")
	ss = strings.Trim(ss, " ")
	objStrList := strings.Split(ss, " ")
	objList := make([]Object, len(objStrList))
	for idx, objStr := range objStrList {
		v, err := strconv.Atoi(objStr)
		vtx := VertexFromIndex(dim, idx)
		if err != nil {
			objList[idx] = NewObject(
				ObjectName(fmt.Sprintf("ErrFromMapLoader-%d_%d", vtx.X, vtx.Y)),
				ObjUndef, vtx)
		}
		objList[idx] = NewObject(
			ObjectName(fmt.Sprintf("ObjFromMapLoader-%d_%d", vtx.X, vtx.Y)),
			ObjectType(v), vtx)
	}
	return objList
}
