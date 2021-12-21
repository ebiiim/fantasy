package input

import (
	"context"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ebiiim/fantasy/base"
)

type Device interface {
	ActionCh() <-chan base.Action
	ListenLoop(ctx context.Context)
}

var _ Device = (*JoinedDevice)(nil)

type JoinedDevice struct {
	actCh chan base.Action
	devs  []Device
}

func NewJoinedDevice(devs ...Device) *JoinedDevice {
	var c JoinedDevice
	c.actCh = make(chan base.Action, 10000) // HACK
	c.devs = devs
	return &c
}

func (c *JoinedDevice) ActionCh() <-chan base.Action {
	return c.actCh
}

func (c *JoinedDevice) ListenLoop(ctx context.Context) {
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

var _ Device = (*Keyboard)(nil)

type Keyboard struct {
	actCh chan base.Action
}

func NewKeyboard() *Keyboard {
	var k Keyboard
	k.actCh = make(chan base.Action, 10000) // HACK
	return &k
}

func (k *Keyboard) ActionCh() <-chan base.Action {
	return k.actCh
}

func (k *Keyboard) ListenLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(1000 / 60 * time.Millisecond): // every frame
			switch {
			case ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp):
				k.actCh <- base.ActUp
				<-time.After(1000 / 6 * time.Millisecond)
			case ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft):
				k.actCh <- base.ActLeft
				<-time.After(1000 / 6 * time.Millisecond)
			case ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown):
				k.actCh <- base.ActDown
				<-time.After(1000 / 6 * time.Millisecond)
			case ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight):
				k.actCh <- base.ActRight
				<-time.After(1000 / 6 * time.Millisecond)
			}
		}
	}
}

var _ Device = (*Mouse)(nil)

type Mouse struct {
	centerTile base.Vertex
	tilePixels base.Vertex
	actCh      chan base.Action
}

func NewMouse(centerTile, tilePixels base.Vertex) *Mouse {
	m := Mouse{
		centerTile,
		tilePixels,
		make(chan base.Action, 10000), // HACK
	}
	return &m
}

func (m *Mouse) ActionCh() <-chan base.Action {
	return m.actCh
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func (m *Mouse) ListenLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(1000 / 60 * time.Millisecond): // every frame
			switch {
			case ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft):
				x, y := ebiten.CursorPosition()
				clicked := base.NewVertex(x, y).Div(m.tilePixels)
				lr := clicked.X - m.centerTile.X
				ud := clicked.Y - m.centerTile.Y
				if abs(lr) == abs(ud) {
					continue
				}
				if abs(lr) < abs(ud) {
					if ud < 0 {
						m.actCh <- base.ActUp
					} else {
						m.actCh <- base.ActDown
					}
				} else {
					if lr < 0 {
						m.actCh <- base.ActLeft
					} else {
						m.actCh <- base.ActRight
					}
				}
				<-time.After(1000 / 6 * time.Millisecond)
			}
		}
	}
}
