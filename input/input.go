package input

import (
	"context"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ebiiim/fantasy/base"
)

type Button uint

const (
	BtnUndef Button = iota
	BtnUp
	BtnLeft
	BtnDown
	BtnRight
	BtnA
	BtnB
	BtnStart
)

// Device represents any input device.
type Device interface {
	// ButtonCh returns the channel for sending captured buttons.
	ButtonCh() <-chan Button
	// ListenLoop starts the listening loop.
	// This is blocking so consider using goroutine.
	ListenLoop(ctx context.Context)
}

var _ Device = (*JoinedDevice)(nil)

// JoinedDevice represents a virtual device that merges multiple input devices.
type JoinedDevice struct {
	btnCh chan Button
	devs  []Device
}

func NewJoinedDevice(devs ...Device) *JoinedDevice {
	var c JoinedDevice
	c.btnCh = make(chan Button)
	c.devs = devs
	return &c
}

func (c *JoinedDevice) ButtonCh() <-chan Button {
	return c.btnCh
}

func (c *JoinedDevice) ListenLoop(ctx context.Context) {
	for _, dev := range c.devs {
		dev := dev
		go func() {
			go dev.ListenLoop(ctx)
			ch := dev.ButtonCh()
			for {
				select {
				case <-ctx.Done():
					return
				case act := <-ch:
					c.btnCh <- act
				}
			}
		}()
	}
	<-ctx.Done()
}

var _ Device = (*Keyboard)(nil)

type Keyboard struct {
	btnCh chan Button
}

func NewKeyboard() *Keyboard {
	var k Keyboard
	k.btnCh = make(chan Button)
	return &k
}

func (k *Keyboard) ButtonCh() <-chan Button {
	return k.btnCh
}

func (k *Keyboard) ListenLoop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(1000 / 60 * time.Millisecond): // every frame
			switch {
			case ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyArrowUp):
				k.btnCh <- BtnUp
				<-time.After(1000 / 6 * time.Millisecond)
			case ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyArrowLeft):
				k.btnCh <- BtnLeft
				<-time.After(1000 / 6 * time.Millisecond)
			case ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyArrowDown):
				k.btnCh <- BtnDown
				<-time.After(1000 / 6 * time.Millisecond)
			case ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyArrowRight):
				k.btnCh <- BtnRight
				<-time.After(1000 / 6 * time.Millisecond)
			case ebiten.IsKeyPressed(ebiten.KeySpace):
				k.btnCh <- BtnA
				<-time.After(1000 / 6 * time.Millisecond)
			}
		}
	}
}

var _ Device = (*Mouse)(nil)

type Mouse struct {
	positionGridCenter base.Vertex
	sizeTilePixel      base.Vertex
	btnCh              chan Button
}

func NewMouse(centerTile, tilePixels base.Vertex) *Mouse {
	m := Mouse{
		centerTile,
		tilePixels,
		make(chan Button),
	}
	return &m
}

func (m *Mouse) ButtonCh() <-chan Button {
	return m.btnCh
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
				clicked := base.NewVertex(x, y).Div(m.sizeTilePixel)
				lr := clicked.X - m.positionGridCenter.X
				ud := clicked.Y - m.positionGridCenter.Y
				if abs(lr) == abs(ud) {
					continue
				}
				if abs(lr) < abs(ud) {
					if ud < 0 {
						m.btnCh <- BtnUp
					} else {
						m.btnCh <- BtnDown
					}
				} else {
					if lr < 0 {
						m.btnCh <- BtnLeft
					} else {
						m.btnCh <- BtnRight
					}
				}
				<-time.After(1000 / 6 * time.Millisecond)
			case ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight):
				m.btnCh <- BtnA
				<-time.After(1000 / 6 * time.Millisecond)
			}
		}
	}
}
