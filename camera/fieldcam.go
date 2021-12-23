package camera

import (
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ebiiim/fantasy/base"
	"github.com/ebiiim/fantasy/field"
	"github.com/ebiiim/fantasy/img"
)

// FieldCam represents the field renderer.
type FieldCam struct {
	// DimGrid represents dimension of the grid.
	// E.g., {16,12}
	DimGrid base.Vertex

	// TilePixels represents tile size.
	// E.g., {40,40}
	TilePixels base.Vertex

	// ScreenResolution represents the screen size that is needed to draw the whole grid.
	// E.g., {640,480} <- {16*40,12*40}
	ScreenResolution base.Vertex
}

// NewFieldCam initializes FieldCam.
func NewFieldCam(dimTiles, tilePixels base.Vertex) *FieldCam {
	c := FieldCam{
		DimGrid:          dimTiles,
		TilePixels:       tilePixels,
		ScreenResolution: dimTiles.Mul(tilePixels),
	}
	return &c
}

// PositionCenter returns the center position of the grid.
func (c *FieldCam) PositionCenter() base.Vertex {
	return base.NewVertex(c.DimGrid.X/2-1, c.DimGrid.Y/2-1)
}

// PositionTopLeft returns the top left position of the grid.
func (c *FieldCam) PositionTopLeft(locCenter base.Vertex) base.Vertex {
	return locCenter.Sub(c.PositionCenter())
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
	for y := 0; y < c.DimGrid.Y; y++ {
		for x := 0; x < c.DimGrid.X; x++ {
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
	if pos.IsOutside(c.DimGrid) {
		return
	}
	drawX := c.TilePixels.X * pos.X
	drawY := c.TilePixels.X * pos.Y
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(drawX), float64(drawY))
	screen.DrawImage(img.Get(obj), op)
}
