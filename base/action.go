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
	MoveLoc Vertex

	// Receive from Field
	MovedLoc Vertex
}
