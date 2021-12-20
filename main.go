package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ebiiim/fantasy/game"
)

var (
	version   string
	buildDate string
	goVersion string
)

func init() {
	// print build info
	setIfEmpty := func(src string, s string) string {
		if src == "" {
			return s
		}
		return src
	}
	println("version " + setIfEmpty(version, "dev"))
	println("build date " + buildDate)
	println(goVersion)
}

func main() {

	g := game.NewGame()
	x, y := g.Layout(0, 0)
	ebiten.SetWindowSize(x, y)
	ebiten.SetWindowTitle("fantasy")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
