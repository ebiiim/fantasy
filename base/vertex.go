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

func (v0 Vertex) ToIndex(dim Vertex) int {
	return dim.X*v0.Y + v0.X
}

func VertexFromIndex(dim Vertex, idx int) Vertex {
	return Vertex{idx % dim.X, idx / dim.X}
}

func (v0 Vertex) IsOutside(dim Vertex) bool {
	return v0.X < 0 || v0.Y < 0 || v0.X >= dim.X || v0.Y >= dim.Y
}
