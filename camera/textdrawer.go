package camera

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/ebiiim/fantasy/base"
)

type TextDrawer struct {
	FontRegular *font.Face
	FontBold    *font.Face
}

func NewTextDrawer(fontSize float64) (*TextDrawer, error) {
	regular, err := opentype.NewFace(fontPixelMplusRegular, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, err
	}
	bold, err := opentype.NewFace(fontPixelMplusBold, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, err
	}
	return &TextDrawer{FontRegular: &regular, FontBold: &bold}, nil
}

func (d *TextDrawer) Draw(screen *ebiten.Image, txt string, pixelPos base.Vertex, clr color.Color, bold, shadow bool) {
	fontFace := d.FontRegular
	if bold {
		fontFace = d.FontBold
	}
	if shadow {
		r, g, b, a := clr.RGBA()
		rev := color.RGBA{
			R: uint8((0xffff - r) >> 8),
			G: uint8((0xffff - g) >> 8),
			B: uint8((0xffff - b) >> 8),
			A: uint8(a),
		}
		text.Draw(screen, txt, *fontFace, pixelPos.X+1, pixelPos.Y+1, rev)
	}
	text.Draw(screen, txt, *fontFace, pixelPos.X, pixelPos.Y, clr)
}
