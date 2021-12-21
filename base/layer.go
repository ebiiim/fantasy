package base

import (
	"strconv"
	"strings"
)

type LayerName string

type Layer struct {
	Name      LayerName
	Dimension Vertex
	Objects   []Object
}

func NewLayer(name LayerName, size Vertex, objList []Object) *Layer {
	l := Layer{
		Name:      name,
		Dimension: size,
		Objects: objList,
	}
	return &l
}

func (l *Layer) GetObject(loc Vertex) Object {
	if loc.IsOutside(l.Dimension) {
		return ObjUndef
	}
	return l.Objects[loc.ToIndex(l.Dimension)]
}

func LoadLayerFromStr(s string) []Object {
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
