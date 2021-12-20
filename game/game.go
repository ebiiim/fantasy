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
	tilePixels     = base.NewVertex(40, 40)
	dimCameraTiles = base.NewVertex(16, 12)
)

var _ ebiten.Game = (*Game)(nil)

type Game struct {
	MapCam   *camera.MapCam
	Map      *base.Map
	Me       *base.Object
	ActionCh <-chan base.Action
}

func NewGame() *Game {

	l001 := []*base.Layer{
		base.NewLayer(base.L_001_a, base.M_001_dim, base.LoadLayerFromStr(base.L_001_a_str)),
		base.NewLayer(base.L_001_b, base.M_001_dim, base.LoadLayerFromStr(base.L_001_b_str)),
	}
	m001 := base.NewMap(base.M_001, base.M_001_dim, l001)

	kbd := input.NewKeyboard()
	camCenter := base.NewVertex(7, 5) // HACK: need calc camera center tile
	mouse := input.NewMouse(camCenter, tilePixels)
	dev := input.NewJoinedDevice(kbd, mouse)
	go dev.ListenLoop(context.Background())

	g := Game{
		MapCam:   camera.NewMapCamera(dimCameraTiles, tilePixels),
		Map:      m001,
		Me:       base.NewObject(base.OBJ_Me, base.NewVertex(6, 5)),
		ActionCh: dev.ActionCh(),
	}
	return &g
}

func (g *Game) Update() error {
	if err := g.Map.Update(); err != nil {
		return err
	}
	select {
	default:
		// no input
	case in := <-g.ActionCh:
		switch in {
		case base.ACT_UP:
			if g.Map.IsMovable(base.NewVertex(g.Me.Loc.X, g.Me.Loc.Y-1)) {
				g.Me.Loc.Y -= 1
			}
		case base.ACT_DOWN:
			if g.Map.IsMovable(base.NewVertex(g.Me.Loc.X, g.Me.Loc.Y+1)) {
				g.Me.Loc.Y += 1
			}
		case base.ACT_LEFT:
			if g.Map.IsMovable(base.NewVertex(g.Me.Loc.X-1, g.Me.Loc.Y)) {
				g.Me.Loc.X -= 1
			}
		case base.ACT_RIGHT:
			if g.Map.IsMovable(base.NewVertex(g.Me.Loc.X+1, g.Me.Loc.Y)) {
				g.Me.Loc.X += 1
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

	g.MapCam.DrawMap(screen, g.Map, g.MapCam.CalcTopLeft(g.Me.Loc))
	g.MapCam.DrawObject(screen, g.Me, g.MapCam.CalcTopLeft(g.Me.Loc))

	drawTime = drawTime / 60 * 59
	drawTime += time.Since(start).Nanoseconds() / 60
	stats := fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f\nDraw: %d us\n%s", ebiten.CurrentTPS(), ebiten.CurrentFPS(), drawTime/1000, guide)
	ebitenutil.DebugPrint(screen, stats)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.MapCam.ScreenResolution.X, g.MapCam.ScreenResolution.Y
}
