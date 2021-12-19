package game

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ebiiim/fantasy/base"
	"github.com/ebiiim/fantasy/img"
)

var (
	objSize   = 40
	DimObject = base.Vertex{X: objSize, Y: objSize}
	DimScreen = base.Vertex{X: 640, Y: 480}
	DimGame   = DimScreen.Div(objSize) // (16,12)
)

type Camera struct {
	Center  base.Vertex
	DimGame base.Vertex
	LeftTop base.Vertex
}

func NewCamera(dim base.Vertex) *Camera {
	var c Camera
	c.DimGame = dim
	return &c
}

func (c *Camera) Update() error {
	c.LeftTop = c.Center.Sub(base.Vertex{X: c.DimGame.X/2 - 1, Y: c.DimGame.Y/2 - 1})
	return nil
}

func (c *Camera) DrawMap(screen *ebiten.Image, m *base.Map) {
	for _, l := range m.Layers {
		c.DrawLayer(screen, l)
	}
}

func (c *Camera) DrawLayer(screen *ebiten.Image, l *base.Layer) {
	for y := 0; y < c.DimGame.Y; y++ {
		for x := 0; x < c.DimGame.X; x++ {
			loc := c.LeftTop.Add(base.Vertex{x, y})
			if loc.IsOutside(l.Size) {
				c.DrawObject(screen, base.NewObject(base.OBJ_BG, loc))
			} else {
				c.DrawObject(screen, l.GetObject(loc))
			}
		}
	}
}

func (c *Camera) DrawObject(screen *ebiten.Image, obj *base.Object) {
	pos := obj.Loc.Sub(c.LeftTop)
	if pos.IsOutside(c.DimGame) {
		return
	}
	op := &ebiten.DrawImageOptions{}
	drawX := DimObject.X * pos.X
	drawY := DimObject.Y * pos.Y
	op.GeoM.Translate(float64(drawX), float64(drawY))

	oi, ok := img.Object2Image[obj.Type]
	if !ok {
		oi = img.Object2Image[base.OBJ_Err]
	}

	screen.DrawImage(oi, op)
}
