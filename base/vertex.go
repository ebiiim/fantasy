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

// ToIndex convertes the Vertex to a 1D index.
// E.g., v0={1,1} dim={5,4} => 6
func (v0 Vertex) ToIndex(dim Vertex) int {
	return dim.X*v0.Y + v0.X
}

// VertexFromIndex converts a 1D index to a 2D index.
// E.g., dim={5,4} idx=6 => {1, 1}
func VertexFromIndex(dim Vertex, idx int) Vertex {
	return Vertex{idx % dim.X, idx / dim.X}
}

// IsOutside checks whether the Vertex is outside of `dim`.
// E.g., v0={4,5}  dim={5,6} => true
//       v0={5,5}  dim={5,6} => false
//		 v0={0,0}  dim={5,6} => true
//		 v0={-1,0} dim={5,6} => false
func (v0 Vertex) IsOutside(dim Vertex) bool {
	return v0.X < 0 || v0.Y < 0 || v0.X >= dim.X || v0.Y >= dim.Y
}
