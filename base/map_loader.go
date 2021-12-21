package base

import (
	"embed"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

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
		ls[idx] = NewLayer(LayerName(v.Name), dim, loadLayerFromStr(v.Objects))
	}
	return NewMap(MapName(md.Name), dim, ls)
}

func loadLayerFromStr(s string) []Object {
	ss := strings.ReplaceAll(s, "\n", " ")
	ss = strings.Trim(ss, " ")
	objStrList := strings.Split(ss, " ")
	objTypeList := make([]Object, len(objStrList))
	for idx, objStr := range objStrList {
		v, err := strconv.Atoi(objStr)
		if err != nil {
			objTypeList[idx] = ObjUndef
		}
		objTypeList[idx] = Object(v)
	}
	return objTypeList
}
