package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ebiiim/fantasy/game"
)

func main() {

	g := game.NewGame()
	x, y := g.Layout(0, 0)
	ebiten.SetWindowSize(x, y)
	ebiten.SetWindowTitle("fantasy")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
