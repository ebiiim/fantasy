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
	FieldCam *camera.FieldCam
	Field    *field.Field
	Me       base.Object
	MyLoc    base.Vertex
	ActionCh <-chan base.Action
}

func NewGame() *Game {

	m := base.NewMap(base.Map01, base.Map01Dim, []*base.Layer{
		base.NewLayer(base.Layer01Base, base.Map01Dim, base.LoadLayerFromStr(base.Layer01BaseStr)),
		base.NewLayer(base.Layer01Objects, base.Map01Dim, base.LoadLayerFromStr(base.Layer01ObjectsStr)),
	})
	f := field.NewField(m)

	kbd := input.NewKeyboard()
	camCenter := base.NewVertex(7, 5) // HACK: need calc camera center tile
	mouse := input.NewMouse(camCenter, tilePixels)
	dev := input.NewJoinedDevice(kbd, mouse)
	go dev.ListenLoop(context.Background())

	g := Game{
		FieldCam: camera.NewFieldCam(dimCameraTiles, tilePixels),
		Field:    f,
		Me:       base.ObjMe,
		MyLoc:    base.NewVertex(6, 5),
		ActionCh: dev.ActionCh(),
	}
	return &g
}

func (g *Game) Update() error {
	if err := g.Field.Update(); err != nil {
		return err
	}
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

	g.FieldCam.DrawField(screen, g.Field, g.FieldCam.CalcTopLeft(g.MyLoc))
	g.FieldCam.DrawObject(screen, g.Me, g.FieldCam.CameraCenterTile())

	drawTime = drawTime / 60 * 59
	drawTime += time.Since(start).Nanoseconds() / 60
	stats := fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f\nDraw: %d us\n%s", ebiten.CurrentTPS(), ebiten.CurrentFPS(), drawTime/1000, guide)

	// debug prints
	ebitenutil.DebugPrint(screen, stats)
	x, y := g.Layout(0, 0)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Version: %s", BuildInfo.Version), x/100*74, y-30)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Build date: %s", BuildInfo.BuildDate.Format("Jan 02 2006 15:04:05")), x/100*74, y-15)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.FieldCam.ScreenResolution.X, g.FieldCam.ScreenResolution.Y
}
