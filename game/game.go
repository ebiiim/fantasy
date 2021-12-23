package game

import (
	"context"
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/ebiiim/fantasy/base"
	"github.com/ebiiim/fantasy/camera"
	"github.com/ebiiim/fantasy/field"
	"github.com/ebiiim/fantasy/input"
)

var (
	// values decided at compile time and injected by main function
	BuildInfo struct {
		Version   string
		BuildDate time.Time
		GoVersion string
	}
)

var (
	tilePixels     = base.NewVertex(40, 40)
	dimCameraTiles = base.NewVertex(16, 12)
)

var _ ebiten.Game = (*Game)(nil)

type Game struct {
	Field    *field.Field
	FieldCam *camera.FieldCam
	Me       base.Object
	MyLoc    base.Vertex
	ActionCh <-chan base.Action
}

func NewGame() *Game {

	m := base.MustLoadMap("assets/map01.yaml")
	f := field.NewField(m)

	fcam := camera.NewFieldCam(dimCameraTiles, tilePixels)
	kbd := input.NewKeyboard()
	mouse := input.NewMouse(fcam.PositionCenter(), tilePixels)
	dev := input.NewJoinedDevice(kbd, mouse)
	go dev.ListenLoop(context.Background())

	g := Game{
		Field:    f,
		FieldCam: fcam,
		Me:       base.ObjMe,
		MyLoc:    base.NewVertex(6, 5),
		ActionCh: dev.ActionCh(),
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
	case in := <-g.ActionCh:
		switch in {
		case base.ActUp:
			if g.Field.IsMovable(base.NewVertex(g.MyLoc.X, g.MyLoc.Y-1)) {
				g.MyLoc.Y -= 1
			}
		case base.ActDown:
			if g.Field.IsMovable(base.NewVertex(g.MyLoc.X, g.MyLoc.Y+1)) {
				g.MyLoc.Y += 1
			}
		case base.ActLeft:
			if g.Field.IsMovable(base.NewVertex(g.MyLoc.X-1, g.MyLoc.Y)) {
				g.MyLoc.X -= 1
			}
		case base.ActRight:
			if g.Field.IsMovable(base.NewVertex(g.MyLoc.X+1, g.MyLoc.Y)) {
				g.MyLoc.X += 1
			}
		}
	}
	return nil
}

var (
	guide          = "Keyboard:\n  Arrow/WASD: Move\nMouse:\n  Left: Move"
	drawTime int64 = 0
)

func (g *Game) Draw(screen *ebiten.Image) {
	start := time.Now()

	g.FieldCam.DrawField(screen, g.Field, g.FieldCam.PositionTopLeft(g.MyLoc))
	g.FieldCam.DrawObject(screen, g.Me, g.FieldCam.PositionCenter())

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
