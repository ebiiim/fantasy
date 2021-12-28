package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/zerolog"

	"github.com/ebiiim/fantasy/game"
)

var (
	version   string
	buildDate string
)

func init() {
	// inject build info
	setIfEmpty := func(src string, s string) string {
		if src == "" {
			return s
		}
		return src
	}
	parseDateOrZero := func(s string) time.Time {
		t, err := time.Parse(time.UnixDate, buildDate)
		if err != nil {
			return time.Time{}
		}
		return t
	}
	game.BuildInfo.Version = setIfEmpty(version, "dev")
	game.BuildInfo.BuildDate = parseDateOrZero(buildDate)
}

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	g := game.NewGame()
	x, y := g.Layout(0, 0)
	ebiten.SetWindowSize(x, y)
	ebiten.SetWindowTitle("fantasy")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
