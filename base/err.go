package base

import "errors"

var (
	ErrNoObjectFound = errors.New("ErrNoObjectFound")
	ErrFieldPut      = errors.New("ErrFieldPut")
	ErrFieldMove     = errors.New("ErrFieldMove")
)
