package camera

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/ebiiim/fantasy/base"
)

var GetImage map[base.ObjectType]*ebiten.Image

func init() {
	GetImage = initRegisterImgs()
}
