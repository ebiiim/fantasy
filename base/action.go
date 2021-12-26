package base

type ActionType uint

const (
	ActUndefined ActionType = iota
	ActMove
	ActMoved
)

type Action struct {
	Type ActionType

	// Send to Field
	MoveAmount Vertex

	// Receive from Field
	MovedLoc Vertex
}
