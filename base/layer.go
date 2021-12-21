package base

import (
	"strconv"
	"strings"
)

type LayerName int

type Layer struct {
	Name    LayerName
	Size    Vertex
	Objects []*Object
}

func NewLayer(name LayerName, size Vertex, objList []ObjectType) *Layer {
	l := Layer{
		Name: name,
		Size: size,
	}
	area := size.X * size.Y
	l.Objects = make([]*Object, area)
	for idx := 0; idx < area; idx++ {
		l.Objects[idx] = NewObject(objList[idx], VertexFromIndex(size, idx))
	}
	return &l
}

func (l *Layer) GetObject(loc Vertex) *Object {
	if loc.IsOutside(l.Size) {
		return NewObject(ObjUndef, loc)
	}
	return l.Objects[loc.ToIndex(l.Size)]
}

func (l *Layer) GetObjectOrError(loc Vertex) (*Object, error) {
	if loc.IsOutside(l.Size) {
		return nil, ErrNoObject
	}
	return l.Objects[loc.ToIndex(l.Size)], nil
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
