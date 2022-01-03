package camera

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ebiiim/fantasy/base"
)

//go:embed assets/*
var assets embed.FS

var m map[base.ObjectType]*ebiten.Image

func init() {
	m = make(map[base.ObjectType]*ebiten.Image)
	initData()
}

// GetImg returns the image for the giving ObjectType.
// Returns ObjUndef if no entry is found for `obj`.
func GetImg(obj base.ObjectType) *ebiten.Image {
	v, ok := m[obj]
	if !ok {
		return m[base.ObjUndef]
	}
	return v
}

// loadImg loads images and panics if it fails.
func loadImg(file string) *ebiten.Image {
	f, err := assets.Open(file)
	if err != nil {
		panic(err)
	}
	im, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(im)
}
