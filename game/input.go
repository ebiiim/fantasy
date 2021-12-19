package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
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
	count     int
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
	curKey := ebiten.Key(-1)
	for {
		<-time.After(1000 / 60 * time.Millisecond) // every frame
		k.count += 1

		// every N frames
		if k.count%12 == 0 {
			keys := inpututil.AppendPressedKeys([]ebiten.Key{})
			if len(keys) > 0 {
				curKey = keys[0]
			}
			switch curKey {
			case ebiten.KeyW, ebiten.KeyArrowUp:
				k.nextInput <- IN_UP
			case ebiten.KeyA, ebiten.KeyArrowLeft:
				k.nextInput <- IN_LEFT
			case ebiten.KeyS, ebiten.KeyArrowDown:
				k.nextInput <- IN_DOWN
			case ebiten.KeyD, ebiten.KeyArrowRight:
				k.nextInput <- IN_RIGHT
			}
			curKey = ebiten.Key(-1)
		}
	}
}
