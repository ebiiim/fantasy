package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ebiiim/fantasy/game"
)

func main() {
	ebiten.SetWindowSize(game.DimScreen.X, game.DimScreen.Y)
	ebiten.SetWindowTitle("fantasy")
	g := game.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
