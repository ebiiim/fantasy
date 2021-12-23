package base

import (
	"math/rand"
	"time"
)

type Intelligent interface {
	Born(f *Field)
	Die()
	Obj() *Object
	RecvCh() <-chan Action
	SendCh() chan<- Action
}

var _ Intelligent = (*Me)(nil)
var _ Intelligent = (*Sheep)(nil)

type Me struct {
	obj         *Object
	toFieldCh   chan Action
	fromFieldCh chan Action
	done        chan struct{}
	field       *Field
}

func NewMe(obj *Object) *Me {
	m := Me{
		obj:         obj,
		toFieldCh:   make(chan Action),
		fromFieldCh: make(chan Action),
		done:        make(chan struct{}),
	}
	return &m
}

func (m *Me) Obj() *Object {
	return m.obj
}

func (m *Me) RecvCh() <-chan Action {
	return m.toFieldCh
}

func (m *Me) SendCh() chan<- Action {
	return m.fromFieldCh
}

func (m *Me) SendMe(act Action) {
	// log.Println("SendMe")
	go func() {
		m.toFieldCh <- act
	}()
}

func (m *Me) Born(f *Field) {
	m.field = f
	for {
		select {
		case act := <-m.fromFieldCh:
			switch act.Type {
			case ActMoved:
				// log.Println("Moved me")
				m.obj.Loc = act.MovedLoc
			}
		case <-m.done:
			return
		}
	}
}

func (m *Me) Die() {
	close(m.done)
}

type Sheep struct {
	obj         *Object
	toFieldCh   chan Action
	fromFieldCh chan Action
	done        chan struct{}
	field       *Field
}

func NewSheep(obj *Object) *Sheep {
	s := Sheep{
		obj:         obj,
		toFieldCh:   make(chan Action),
		fromFieldCh: make(chan Action),
		done:        make(chan struct{}),
	}
	return &s
}

func (s *Sheep) Obj() *Object {
	return s.obj
}

func (s *Sheep) RecvCh() <-chan Action {
	return s.toFieldCh
}

func (s *Sheep) SendCh() chan<- Action {
	return s.fromFieldCh
}

func (s *Sheep) Born(f *Field) {
	s.field = f
	for {
		select {
		case <-s.done:
			return
		case <-time.After(time.Second):
			// log.Println("Sheep try to move")
			axis := rand.Intn(10)     // X:Y=7:3
			value := rand.Intn(3) - 1 // -1,0,1

			moveLoc := s.obj.Loc
			if axis < 3 {
				moveLoc = NewVertex(moveLoc.X+value, moveLoc.Y)
			} else {
				moveLoc = NewVertex(moveLoc.X, moveLoc.Y+value)
			}
			act := Action{
				Type:    ActMove,
				MoveLoc: moveLoc,
			}
			s.toFieldCh <- act
		case act := <-s.fromFieldCh:
			switch act.Type {
			case ActMoved:
				// log.Println("Sheep moved")
				s.obj.Loc = act.MovedLoc
			}
		}
	}
}

func (s *Sheep) Die() {
	close(s.done)
}
