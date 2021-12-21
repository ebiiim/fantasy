package camera

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ebiiim/fantasy/base"
)

//go:generate rm -rf ./assets
//go:generate cp -r ../assets ./assets
//go:embed assets/img/*
var assets embed.FS

func initRegisterImgs() map[base.ObjectType]*ebiten.Image {
	load := func(file string) *ebiten.Image {
		f, err := assets.Open(file)
		if err != nil {
			panic(err)
		}
		img, _, err := image.Decode(f)
		if err != nil {
			panic(err)
		}
		return ebiten.NewImageFromImage(img)
	}

	m := make(map[base.ObjectType]*ebiten.Image)

	m[base.ObjUndef] = load("assets/img/undef.png")
	m[base.ObjMe] = load("assets/img/me.png")
	m[base.ObjBG] = load("assets/img/bg.png")
	m[base.ObjNone] = load("assets/img/none.png")
	m[base.ObjBase] = load("assets/img/base.png")
	m[base.ObjGrass] = load("assets/img/grass.png")
	m[base.ObjTree] = load("assets/img/tree.png")
	m[base.ObjBox] = load("assets/img/box.png")

	return m
}
