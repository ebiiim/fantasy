package game

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/ebiiim/fantasy/base"
	"github.com/ebiiim/fantasy/camera"
	"github.com/ebiiim/fantasy/input"
)

var (
	// values decided at compile time and injected by main function
	BuildInfo struct {
		Version   string
		BuildDate time.Time
	}
)

var (
	tilePixels     = base.NewVertex(40, 40)
	dimCameraTiles = base.NewVertex(16, 12)
)

var _ ebiten.Game = (*Game)(nil)

type Game struct {
	Field    *base.Field
	FieldCam *camera.FieldCam
	Me       *base.Me
	ButtonCh <-chan input.Button
}

func NewGame() *Game {

	m := base.MustLoadMap("assets/map01.yaml")
	f := base.NewField(m)

	me := base.NewMe()
	sheep1 := base.NewSheep()
	sheep2 := base.NewSheep()
	sheep3 := base.NewSheep()
	f.PutIntelligent(me, base.NewVertex(6, 5))
	f.PutIntelligent(sheep1, base.NewVertex(4, 5))
	f.PutIntelligent(sheep2, base.NewVertex(10, 8))
	f.PutIntelligent(sheep3, base.NewVertex(15, 4))

	fcam := camera.NewFieldCam(dimCameraTiles, tilePixels)
	kbd := input.NewKeyboard()
	mouse := input.NewMouse(fcam.PositionCenter(), tilePixels)
	dev := input.NewJoinedDevice(kbd, mouse)
	go dev.ListenLoop(context.Background())

	g := Game{
		Field:    f,
		FieldCam: fcam,
		Me:       me,
		ButtonCh: dev.ButtonCh(),
	}
	return &g
}

func (g *Game) Update() error {
	// update the field
	if err := g.Field.Update(); err != nil {
		return err
	}

	// check inputs and do actions
	select {
	default:
		// no input
	case btn := <-g.ButtonCh:
		var moveAmount base.Vertex
		switch btn {
		case input.BtnUp:
			moveAmount = base.NewVertex(0, -1)
		case input.BtnDown:
			moveAmount = base.NewVertex(0, 1)
		case input.BtnLeft:
			moveAmount = base.NewVertex(-1, 0)
		case input.BtnRight:
			moveAmount = base.NewVertex(1, 0)
		case input.BtnA:
			// add sheep to randam location
			loc := base.NewVertex(-1, -1)
			for !g.Field.IsMovable(loc) {
				loc = base.NewVertex(rand.Intn(g.FieldCam.DimGrid.X), rand.Intn(g.FieldCam.DimGrid.Y))
			}
			sp := base.NewSheep()
			g.Field.PutIntelligent(sp, loc)
		}
		if moveAmount.X+moveAmount.Y != 0 {
			ch := g.Me.ToFieldCh()
			go func() {
				ch <- base.Action{
					Type:       base.ActMove,
					MoveAmount: moveAmount,
				}
			}()
		}
	}
	return nil
}

var (
	guide          = "Keyboard:\n  Arrow/WASD: Move\n  Space: Baa\nMouse:\n  Left: Move\n  Right: Baa"
	drawTime int64 = 0
)

func (g *Game) Draw(screen *ebiten.Image) {
	start := time.Now()

	g.FieldCam.DrawField(screen, g.Field, g.FieldCam.PositionTopLeft(g.Me.Loc()))

	g.drawDebugPrints(screen, start)
}

func (g *Game) drawDebugPrints(screen *ebiten.Image, started time.Time) {
	drawTime = drawTime / 60 * 59
	drawTime += time.Since(started).Nanoseconds() / 60
	stats := fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f\nDraw: %d us\n%s", ebiten.CurrentTPS(), ebiten.CurrentFPS(), drawTime/1000, guide)
	ebitenutil.DebugPrint(screen, stats)
	x, y := g.Layout(0, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Version: %s", BuildInfo.Version), x/100*75, y-30)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Build date: %s", BuildInfo.BuildDate.Format(time.RFC822)), x/100*75, y-15)

}

// Layout returns the screen resolution that is needed to draw the grid.
// Always returns same value for now.
func (g *Game) Layout(_, _ int) (int, int) {
	return g.FieldCam.ScreenResolution.X, g.FieldCam.ScreenResolution.Y
}
