package base

import "errors"

var (
	ErrNoObjectFound = errors.New("ErrNoObjectFound")
	ErrFieldPut      = errors.New("ErrFieldPut")
	ErrFieldDelete   = errors.New("ErrFieldDelete")
	ErrFieldMove     = errors.New("ErrFieldMove")
)
