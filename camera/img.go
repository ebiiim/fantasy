package camera

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ebiiim/fantasy/base"
	"github.com/ebiiim/fantasy/log"
)

var lg = log.NewLogger("Img")

//go:embed assets/*
var assets embed.FS

var m map[base.ObjectType]map[base.Posture]*ebiten.Image

func init() {
	m = make(map[base.ObjectType]map[base.Posture]*ebiten.Image)
	initData()
}

// GetImg returns the image for the giving ObjectType and Posture.
// Returns ObjUndef if no entry is found for `obj`.
// Returns image of PosNone if no entry is found for `posture`.
func GetImg(obj base.ObjectType, posture base.Posture) *ebiten.Image {
	v, ok := m[obj]
	if !ok {
		lg.Warn(log.TypeInternal, "GetImg", "", "no img for obj=%v", obj)
		return m[base.ObjUndef][base.PosNone]
	}

	img, ok := v[posture]
	if ok {
		return img
	}

	imgPosNone, ok := v[base.PosNone]
	if ok {
		return imgPosNone
	}

	lg.Warn(log.TypeInternal, "GetImg", "", "no img for obj=%v pos=%v", obj, posture)
	return m[base.ObjUndef][base.PosNone]
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
