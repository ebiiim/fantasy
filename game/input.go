package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
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
	go k.inputLoop()
	return &k
}

func (k *KBDInput) UserInputCh() <-chan UserInput {
	return k.nextInput
}

func (k *KBDInput) inputLoop() {
	for {
		<-time.After(1000 / 60 * time.Millisecond) // every frame
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
