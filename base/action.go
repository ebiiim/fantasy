package base

type ActionType uint

const (
	ActUndefined ActionType = iota

	ActMove

	ActEcho

	// ActDie tells field to remove the actor
	ActDie
)

type Action struct {
	Type ActionType

	MoveAmount Vertex

	EchoWho  ObjectName
	EchoBody string
}
