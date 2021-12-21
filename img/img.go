package img

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ebiiim/fantasy/base"
)

var m map[base.Object]*ebiten.Image

func Get(obj base.Object) *ebiten.Image {
	v, ok := m[obj]
	if !ok {
		return m[base.ObjUndef]
	}
	return v
}
