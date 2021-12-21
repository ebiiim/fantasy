package img

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ebiiim/fantasy/base"
)


//go:embed assets/*
var assets embed.FS

var m map[base.Object]*ebiten.Image

func init() {
	m = make(map[base.Object]*ebiten.Image)
	initData()
}

func Get(obj base.Object) *ebiten.Image {
	v, ok := m[obj]
	if !ok {
		return m[base.ObjUndef]
	}
	return v
}

func load(file string) *ebiten.Image {
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
