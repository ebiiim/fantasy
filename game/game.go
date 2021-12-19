package game

import (
	"context"
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/ebiiim/fantasy/base"
)

var _ ebiten.Game = (*Game)(nil)

type Game struct {
	Camera  *Camera
	Map     *base.Map
	Me      *base.Object
	InputCh <-chan UserInput
}

func NewGame() *Game {
	l001 := []*base.Layer{
		base.NewLayer(base.L_001_a, base.M_001_dim, base.LoadLayerFromStr(base.L_001_a_str)),
		base.NewLayer(base.L_001_b, base.M_001_dim, base.LoadLayerFromStr(base.L_001_b_str)),
	}
	m001 := base.NewMap(base.M_001, base.M_001_dim, l001)

	inputCh := make(chan UserInput, 10000) // HACK

	centerTile := base.NewVertex(7, 5) // HACK: need calc camera center tile
	kbd := NewKBDInput()
	mouse := NewMouseInput(centerTile)
	go kbd.StartInputLoop(context.Background())
	go mouse.StartInputLoop(context.Background())
	go func() {
		kbdCh := kbd.UserInputCh()
		mouseCh := mouse.UserInputCh()
		for {
			select {
			case in := <-kbdCh:
				inputCh <- in
			case in := <-mouseCh:
				inputCh <- in
			}
		}
	}()

	g := Game{
		Camera:  NewCamera(DimCameraTiles),
		Map:     m001,
		Me:      base.NewObject(base.OBJ_Me, base.NewVertex(6, 5)),
		InputCh: inputCh,
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
	case in := <-g.InputCh:
		switch in {
		case IN_UP:
			if g.Map.IsMovable(base.NewVertex(g.Me.Loc.X, g.Me.Loc.Y-1)) {
				g.Me.Loc.Y -= 1
			}
		case IN_DOWN:
			if g.Map.IsMovable(base.NewVertex(g.Me.Loc.X, g.Me.Loc.Y+1)) {
				g.Me.Loc.Y += 1
			}
		case IN_LEFT:
			if g.Map.IsMovable(base.NewVertex(g.Me.Loc.X-1, g.Me.Loc.Y)) {
				g.Me.Loc.X -= 1
			}
		case IN_RIGHT:
			if g.Map.IsMovable(base.NewVertex(g.Me.Loc.X+1, g.Me.Loc.Y)) {
				g.Me.Loc.X += 1
			}
		}
	}
	g.Camera.Center = g.Me.Loc
	g.Camera.Update()
	return nil
}

var (
	guide          = "Keyboard:\n  Arrow/WASD: Move\nMouse:\n  Left: Move"
	drawTime int64 = 0
)

func (g *Game) Draw(screen *ebiten.Image) {
	start := time.Now()

	g.Camera.DrawMap(screen, g.Map)
	g.Camera.DrawObject(screen, g.Me)

	drawTime = drawTime / 60 * 59
	drawTime += time.Since(start).Nanoseconds() / 60
	stats := fmt.Sprintf("TPS: %0.2f\nFPS: %0.2f\nDraw: %d us\n%s", ebiten.CurrentTPS(), ebiten.CurrentFPS(), drawTime/1000, guide)
	ebitenutil.DebugPrint(screen, stats)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return ScreenResolution.X, ScreenResolution.Y
}
