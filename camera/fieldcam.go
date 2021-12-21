package camera

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ebiiim/fantasy/base"
	"github.com/ebiiim/fantasy/field"
	"github.com/ebiiim/fantasy/img"
)

type FieldCam struct {
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

func NewFieldCam(dimTiles, tilePixels base.Vertex) *FieldCam {
	c := FieldCam{
		DimTiles:         dimTiles,
		TilePixels:       tilePixels,
		ScreenResolution: dimTiles.Mul(tilePixels),
	}
	return &c
}

func (c *FieldCam) CameraCenterTile() base.Vertex {
	return base.NewVertex(c.DimTiles.X/2-1, c.DimTiles.Y/2-1)
}

// CalcTopLeft calcs screen top left tile's location in map.
func (c *FieldCam) CalcTopLeft(locCenter base.Vertex) base.Vertex {
	return locCenter.Sub(c.CameraCenterTile())
}

func (c *FieldCam) DrawField(screen *ebiten.Image, f *field.Field, topLeft base.Vertex) {
	c.DrawMap(screen, f.Map, topLeft)
}

func (c *FieldCam) DrawMap(screen *ebiten.Image, m *base.Map, topLeft base.Vertex) {
	for _, l := range m.Layers {
		c.DrawLayer(screen, l, topLeft)
	}
}

func (c *FieldCam) DrawLayer(screen *ebiten.Image, l *base.Layer, topLeft base.Vertex) {
	for y := 0; y < c.DimTiles.Y; y++ {
		for x := 0; x < c.DimTiles.X; x++ {
			pos := base.NewVertex(x, y)
			loc := topLeft.Add(pos)
			if loc.IsOutside(l.Dimension) {
				c.DrawObject(screen, base.ObjBG, pos)
			} else {
				c.DrawObject(screen, l.GetObject(loc), pos)
			}
		}
	}
}

func (c *FieldCam) DrawObject(screen *ebiten.Image, obj base.Object, pos base.Vertex) {
	if pos.IsOutside(c.DimTiles) {
		return
	}
	drawX := c.TilePixels.X * pos.X
	drawY := c.TilePixels.X * pos.Y
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(drawX), float64(drawY))
	screen.DrawImage(img.Get(obj), op)
}
