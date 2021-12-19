package game

import (
	"context"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ebiiim/fantasy/base"
)

type UserInput int

const (
	IN_UNDEF UserInput = iota
	IN_UP
	IN_DOWN
	IN_LEFT
	IN_RIGHT
)

type KBDInput struct {
	nextInput chan UserInput
}

func NewKBDInput() *KBDInput {
	k := KBDInput{}
	k.nextInput = make(chan UserInput, 10000) // HACK
	return &k
}

func (k *KBDInput) UserInputCh() <-chan UserInput {
	return k.nextInput
}

func (k *KBDInput) StartInputLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(1000 / 60 * time.Millisecond): // every frame
			switch {
			case ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp):
				k.nextInput <- IN_UP
				<-time.After(1000 / 6 * time.Millisecond)
			case ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft):
				k.nextInput <- IN_LEFT
				<-time.After(1000 / 6 * time.Millisecond)
			case ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown):
				k.nextInput <- IN_DOWN
				<-time.After(1000 / 6 * time.Millisecond)
			case ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight):
				k.nextInput <- IN_RIGHT
				<-time.After(1000 / 6 * time.Millisecond)
			}
		}
	}
}

type MouseInput struct {
	centerTile base.Vertex
	nextInput  chan UserInput
}

func NewMouseInput(centerTile base.Vertex) *MouseInput {
	m := MouseInput{centerTile: centerTile}
	m.nextInput = make(chan UserInput, 10000) // HACK
	return &m
}

func (m *MouseInput) UserInputCh() <-chan UserInput {
	return m.nextInput
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
func (m *MouseInput) StartInputLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(1000 / 60 * time.Millisecond): // every frame
			switch {
			case ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft):
				x, y := ebiten.CursorPosition()
				clicked := base.Vertex{x, y}.Div(ObjectPixels)
				lr := clicked.X - m.centerTile.X
				ud := clicked.Y - m.centerTile.Y
				if abs(lr) == abs(ud) {
					continue
				}
				if abs(lr) < abs(ud) {
					if ud < 0 {
						m.nextInput <- IN_UP
					} else {
						m.nextInput <- IN_DOWN
					}
				} else {
					if lr < 0 {
						m.nextInput <- IN_LEFT
					} else {
						m.nextInput <- IN_RIGHT
					}
				}
				<-time.After(1000 / 6 * time.Millisecond)
			}
		}
	}
}
