package base

import (
	"strconv"
	"strings"
)

type LayerName string

type Layer struct {
	Name      LayerName
	Dimension Vertex
	Objects   []*Object
}

func NewLayer(name LayerName, size Vertex, objList []ObjectType) *Layer {
	l := Layer{
		Name:      name,
		Dimension: size,
	}
	area := size.X * size.Y
	l.Objects = make([]*Object, area)
	for idx := 0; idx < area; idx++ {
		l.Objects[idx] = NewObject(objList[idx], VertexFromIndex(size, idx))
	}
	return &l
}

func (l *Layer) GetObject(loc Vertex) *Object {
	if loc.IsOutside(l.Dimension) {
		return NewObject(ObjUndef, loc)
	}
	return l.Objects[loc.ToIndex(l.Dimension)]
}

func LoadLayerFromStr(s string) []ObjectType {
	ss := strings.ReplaceAll(s, "\n", " ")
	ss = strings.Trim(ss, " ")
	objStrList := strings.Split(ss, " ")
	objTypeList := make([]ObjectType, len(objStrList))
	for idx, objStr := range objStrList {
		v, err := strconv.Atoi(objStr)
		if err != nil {
			objTypeList[idx] = ObjUndef
		}
		objTypeList[idx] = ObjectType(v)
	}
	return objTypeList
}
