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
		return NewObject(OBJ_Err, loc)
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
	ss := strings.Split(s, "\n")
	// preprocess
	for idx, s := range ss {
		if strings.HasPrefix(s, "#") {
			s = ""
		}
		s = strings.ReplaceAll(s, " ", "")
		s = strings.TrimSuffix(s, ",")
		ss[idx] = s
	}
	objStr := strings.Join(ss, ",")
	// bad hack
	objStr = strings.ReplaceAll(objStr, ",,", ",")
	objStr = strings.ReplaceAll(objStr, ",,", ",")
	objStr = strings.ReplaceAll(objStr, ",,", ",")
	objStr = strings.ReplaceAll(objStr, ",,", ",")
	objStr = strings.ReplaceAll(objStr, ",,", ",")
	objStr = strings.TrimPrefix(objStr, ",")
	objStr = strings.TrimSuffix(objStr, ",")

	objStrList := strings.Split(objStr, ",")

	objTypeList := make([]ObjectType, len(objStrList))
	for idx, objStr := range objStrList {
		v, err := strconv.Atoi(objStr)
		if err != nil {
			objTypeList[idx] = OBJ_Err
		}
		objTypeList[idx] = ObjectType(v)
	}
	return objTypeList
}
