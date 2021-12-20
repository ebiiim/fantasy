package game

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ebiiim/fantasy/base"
)

var (
	objSize          = 40
	ObjectPixels     = base.NewVertex(objSize, objSize)
	ScreenResolution = base.NewVertex(640, 480)
	DimCameraTiles   = ScreenResolution.Div(ObjectPixels) // (16,12)
)

type Camera struct {
	Center   base.Vertex
	DimTiles base.Vertex
	LeftTop  base.Vertex
}

func NewCamera(dim base.Vertex) *Camera {
	var c Camera
	c.DimTiles = dim
	return &c
}

func (c *Camera) Update() error {
	c.LeftTop = c.Center.Sub(base.NewVertex(c.DimTiles.X/2-1, c.DimTiles.Y/2-1))
	return nil
}

func (c *Camera) DrawMap(screen *ebiten.Image, m *base.Map) {
	for _, l := range m.Layers {
		c.DrawLayer(screen, l)
	}
}

func (c *Camera) DrawLayer(screen *ebiten.Image, l *base.Layer) {
	for y := 0; y < c.DimTiles.Y; y++ {
		for x := 0; x < c.DimTiles.X; x++ {
			loc := c.LeftTop.Add(base.NewVertex(x, y))
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
	if pos.IsOutside(c.DimTiles) {
		return
	}
	op := &ebiten.DrawImageOptions{}
	drawX := ObjectPixels.X * pos.X
	drawY := ObjectPixels.Y * pos.Y
	op.GeoM.Translate(float64(drawX), float64(drawY))

	oi, ok := Object2Image[obj.Type]
	if !ok {
		oi = Object2Image[base.OBJ_Err]
	}

	screen.DrawImage(oi, op)
}
