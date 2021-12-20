package game

import (
	"context"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ebiiim/fantasy/base"
)

type InputDevice interface {
	ActionCh() <-chan InputAction
	ListenLoop(ctx context.Context)
}

type InputAction int

const (
	IN_UNDEF InputAction = iota
	IN_UP
	IN_DOWN
	IN_LEFT
	IN_RIGHT
)

var _ InputDevice = (*JoinedInput)(nil)

type JoinedInput struct {
	actCh chan InputAction
	devs  []InputDevice
}

func NewJoinedInput(devs ...InputDevice) *JoinedInput {
	var c JoinedInput
	c.actCh = make(chan InputAction, 10000) // HACK
	c.devs = devs
	return &c
}

func (c *JoinedInput) ActionCh() <-chan InputAction {
	return c.actCh
}

func (c *JoinedInput) ListenLoop(ctx context.Context) {
	for _, dev := range c.devs {
		dev := dev
		go func() {
			go dev.ListenLoop(ctx)
			ch := dev.ActionCh()
			for {
				select {
				case <-ctx.Done():
					return
				case act := <-ch:
					c.actCh <- act
				}
			}
		}()
	}
	<-ctx.Done()
}

var _ InputDevice = (*KBDInput)(nil)

type KBDInput struct {
	actCh chan InputAction
}

func NewKBDInput() *KBDInput {
	var k KBDInput
	k.actCh = make(chan InputAction, 10000) // HACK
	return &k
}

func (k *KBDInput) ActionCh() <-chan InputAction {
	return k.actCh
}

func (k *KBDInput) ListenLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(1000 / 60 * time.Millisecond): // every frame
			switch {
			case ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp):
				k.actCh <- IN_UP
				<-time.After(1000 / 6 * time.Millisecond)
			case ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft):
				k.actCh <- IN_LEFT
				<-time.After(1000 / 6 * time.Millisecond)
			case ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown):
				k.actCh <- IN_DOWN
				<-time.After(1000 / 6 * time.Millisecond)
			case ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight):
				k.actCh <- IN_RIGHT
				<-time.After(1000 / 6 * time.Millisecond)
			}
		}
	}
}

var _ InputDevice = (*MouseInput)(nil)

type MouseInput struct {
	centerTile base.Vertex
	actCh      chan InputAction
}

func NewMouseInput(centerTile base.Vertex) *MouseInput {
	m := MouseInput{centerTile: centerTile}
	m.actCh = make(chan InputAction, 10000) // HACK
	return &m
}

func (m *MouseInput) ActionCh() <-chan InputAction {
	return m.actCh
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func (m *MouseInput) ListenLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(1000 / 60 * time.Millisecond): // every frame
			switch {
			case ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft):
				x, y := ebiten.CursorPosition()
				clicked := base.NewVertex(x, y).Div(ObjectPixels)
				lr := clicked.X - m.centerTile.X
				ud := clicked.Y - m.centerTile.Y
				if abs(lr) == abs(ud) {
					continue
				}
				if abs(lr) < abs(ud) {
					if ud < 0 {
						m.actCh <- IN_UP
					} else {
						m.actCh <- IN_DOWN
					}
				} else {
					if lr < 0 {
						m.actCh <- IN_LEFT
					} else {
						m.actCh <- IN_RIGHT
					}
				}
				<-time.After(1000 / 6 * time.Millisecond)
			}
		}
	}
}
