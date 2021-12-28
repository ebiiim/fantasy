package base

import (
	"math/rand"
	"time"
)

type Intelligent interface {
	Born(f *Field)
	Die()
	RecvCh() <-chan Action
	SendCh() chan<- Action

	Locatable
}

var _ Intelligent = (*Me)(nil)
var _ Intelligent = (*Sheep)(nil)

type Me struct {
	*BaseObject
	toFieldCh   chan Action
	fromFieldCh chan Action
	done        chan struct{}
	field       *Field
}

func NewMe(obj *BaseObject) *Me {
	m := Me{
		BaseObject:  obj,
		toFieldCh:   make(chan Action),
		fromFieldCh: make(chan Action),
		done:        make(chan struct{}),
	}
	return &m
}

func (x *Me) RecvCh() <-chan Action {
	return x.toFieldCh
}

func (x *Me) SendCh() chan<- Action {
	return x.fromFieldCh
}

func (x *Me) SendMe(act Action) {
	// log.Println("SendMe")
	go func() {
		x.toFieldCh <- act
	}()
}

func (x *Me) Born(f *Field) {
	x.field = f
	for {
		select {
		case act := <-x.fromFieldCh:
			switch act.Type {
			case ActMoved:
				// log.Println("Moved me")
				x.SetLoc(act.MovedLoc)
			}
		case <-x.done:
			return
		}
	}
}

func (x *Me) Die() {
	close(x.done)
}

type Sheep struct {
	*BaseObject
	toFieldCh   chan Action
	fromFieldCh chan Action
	done        chan struct{}
	field       *Field
}

func NewSheep(obj *BaseObject) *Sheep {
	s := Sheep{
		BaseObject:      obj,
		toFieldCh:   make(chan Action),
		fromFieldCh: make(chan Action),
		done:        make(chan struct{}),
	}
	return &s
}

func (x *Sheep) RecvCh() <-chan Action {
	return x.toFieldCh
}

func (x *Sheep) SendCh() chan<- Action {
	return x.fromFieldCh
}

func (x *Sheep) Born(f *Field) {
	x.field = f
	for {
		select {
		case <-x.done:
			return
		case <-time.After(time.Millisecond * time.Duration(500+rand.Intn(2000))):
			// log.Println("Sheep try to move")
			axis := rand.Intn(10)     // X:Y=7:3
			value := rand.Intn(3) - 1 // -1,0,1
			var moveAmount Vertex
			if axis < 3 {
				moveAmount = NewVertex(value, 0)
			} else {
				moveAmount = NewVertex(0, value)
			}
			act := Action{
				Type:       ActMove,
				MoveAmount: moveAmount,
			}
			x.toFieldCh <- act
		case act := <-x.fromFieldCh:
			switch act.Type {
			case ActMoved:
				// log.Println("Sheep moved")
				x.SetLoc(act.MovedLoc)
			}
		}
	}
}

func (x *Sheep) Die() {
	close(x.done)
}
