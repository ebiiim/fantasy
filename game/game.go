package game

import (
	"context"
	"fmt"
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
		GoVersion string
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

	me := base.NewMe(base.NewObject(base.ObjMe, base.NewVertex(6, 5)))
	sheep1 := base.NewSheep(base.NewObject(base.ObjSheep, base.NewVertex(8, 5)))
	sheep2 := base.NewSheep(base.NewObject(base.ObjSheep, base.NewVertex(9, 10)))
	f.AddIntelligent(me)
	f.AddIntelligent(sheep1)
	f.AddIntelligent(sheep2)

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
		movedLoc := base.NewVertex(-1, -1)
		switch btn {
		case input.BtnUp:
			movedLoc = base.NewVertex(g.Me.Obj().Loc.X, g.Me.Obj().Loc.Y-1)
		case input.BtnDown:
			movedLoc = base.NewVertex(g.Me.Obj().Loc.X, g.Me.Obj().Loc.Y+1)
		case input.BtnLeft:
			movedLoc = base.NewVertex(g.Me.Obj().Loc.X-1, g.Me.Obj().Loc.Y)
		case input.BtnRight:
			movedLoc = base.NewVertex(g.Me.Obj().Loc.X+1, g.Me.Obj().Loc.Y)
		}
		if movedLoc.X != -1 {
			g.Me.SendMe(base.Action{
				Type:    base.ActMove,
				MoveLoc: movedLoc,
			})
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

	g.FieldCam.DrawField(screen, g.Field, g.FieldCam.PositionTopLeft(g.Me.Obj().Loc))

	for _, intelli := range g.Field.Intelligents {
		g.FieldCam.DrawObject(screen, intelli.Obj(), intelli.Obj().Loc.Sub(g.FieldCam.PositionTopLeft(g.Me.Obj().Loc)))
	}

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
