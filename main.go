package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ebiiim/fantasy/game"
)

func main() {
	ebiten.SetWindowSize(game.ScreenResolution.X, game.ScreenResolution.Y)
	ebiten.SetWindowTitle("fantasy")
	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
