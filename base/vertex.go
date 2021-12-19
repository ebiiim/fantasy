package base

type Vertex struct {
	X int
	Y int
}

func NewVertex(x, y int) Vertex {
	return Vertex{x, y}
}

func (v0 Vertex) Add(v Vertex) Vertex {
	return Vertex{v0.X + v.X, v0.Y + v.Y}
}

func (v0 Vertex) Sub(v Vertex) Vertex {
	return Vertex{v0.X - v.X, v0.Y - v.Y}
}

func (v0 Vertex) Mul(v Vertex) Vertex {
	return Vertex{v0.X * v.X, v0.Y * v.Y}
}

func (v0 Vertex) Div(v Vertex) Vertex {
	return Vertex{v0.X / v.X, v0.Y / v.Y}
}

func (v Vertex) ToIndex(dim Vertex) int {
	return dim.X*v.Y + v.X
}

func VertexFromIndex(dim Vertex, idx int) Vertex {
	return Vertex{idx % dim.X, idx / dim.X}
}

func (v Vertex) IsOutside(dim Vertex) bool {
	return v.X < 0 || v.Y < 0 || v.X >= dim.X || v.Y >= dim.Y
}
