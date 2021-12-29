package main

import (
	"flag"
	"runtime"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/rs/zerolog"

	"github.com/ebiiim/fantasy/game"
	"github.com/ebiiim/fantasy/log"
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

var lg = log.NewLogger("main")

func main() {

	var logLevel int
	flag.IntVar(&logLevel, "v", 1, "log level")
	flag.Parse()

	switch logLevel {
	default:
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case 1:
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case 2:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case 3:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case 4:
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case 5:
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	}

	// stats
	go func() {
		startTime := time.Now()
		var mem runtime.MemStats
		for {
			<-time.After(5 * time.Second)
			n := runtime.NumGoroutine()
			runtime.ReadMemStats(&mem)
			lg.Info(log.TypeSystem, "main", "", "NumGoroutine=%d MemAlloc=%dKB Uptime=%v", n, mem.Alloc/1024, time.Since(startTime).Round(time.Second))
		}
	}()

	g := game.NewGame()
	x, y := g.Layout(0, 0)
	ebiten.SetWindowSize(x, y)
	ebiten.SetWindowTitle("fantasy")

	if err := ebiten.RunGame(g); err != nil {
		lg.Fatal(log.TypeSystem, "main", "", "RunGame returned error %v", err)
	}
}
