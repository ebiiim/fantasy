package camera

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ebiiim/fantasy/base"
)

type MapCam struct {
	// DimTiles represents tiles to draw.
	// e.g., {16,12}
	DimTiles base.Vertex

	// TilePixels represents tile size.
	// e.g., {40,40}
	TilePixels base.Vertex

	// ScreenResolution represents screen size needed to draw tiles.
	// e.g., {640,480} <- {16*40,12*40}
	ScreenResolution base.Vertex
}

func NewMapCamera(dimTiles, tilePixels base.Vertex) *MapCam {
	c := MapCam{
		DimTiles:         dimTiles,
		TilePixels:       tilePixels,
		ScreenResolution: dimTiles.Mul(tilePixels),
	}
	return &c
}

// CalcTopLeft calcs screen top left tile's location in map.
func (c *MapCam) CalcTopLeft(locCenter base.Vertex) base.Vertex {
	return locCenter.Sub(base.NewVertex(c.DimTiles.X/2-1, c.DimTiles.Y/2-1))
}

func (c *MapCam) DrawMap(screen *ebiten.Image, m *base.Map, topLeft base.Vertex) {
	for _, l := range m.Layers {
		c.DrawLayer(screen, l, topLeft)
	}
}

func (c *MapCam) DrawLayer(screen *ebiten.Image, l *base.Layer, topLeft base.Vertex) {
	for y := 0; y < c.DimTiles.Y; y++ {
		for x := 0; x < c.DimTiles.X; x++ {
			loc := topLeft.Add(base.NewVertex(x, y))
			if loc.IsOutside(l.Size) {
				c.DrawObject(screen, base.NewObject(base.ObjBG, loc), topLeft)
			} else {
				c.DrawObject(screen, l.GetObject(loc), topLeft)
			}
		}
	}
}

func (c *MapCam) DrawObject(screen *ebiten.Image, obj *base.Object, locTopLeft base.Vertex) {
	pos := obj.Loc.Sub(locTopLeft)
	if pos.IsOutside(c.DimTiles) {
		return
	}
	op := &ebiten.DrawImageOptions{}
	drawX := c.TilePixels.X * pos.X
	drawY := c.TilePixels.X * pos.Y
	op.GeoM.Translate(float64(drawX), float64(drawY))

	oi, ok := GetImage[obj.Type]
	if !ok {
		oi = GetImage[base.ObjUndef]
	}

	screen.DrawImage(oi, op)
}
