package game

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/ebiiim/fantasy/base"
)

var Object2Image map[base.ObjectType]*ebiten.Image
var object2ImageFile map[base.ObjectType]string

func init() {
	log.Println("load images")

	object2ImageFile = make(map[base.ObjectType]string)
	
	object2ImageFile[base.OBJ_Err] = "assets/img/err.png"
	object2ImageFile[base.OBJ_Me] = "assets/img/me.png"
	object2ImageFile[base.OBJ_BG] = "assets/img/bg.png"
	object2ImageFile[base.OBJ_None] = "assets/img/none.png"
	object2ImageFile[base.OBJ_Base] = "assets/img/base.png"
	object2ImageFile[base.OBJ_Grass] = "assets/img/grass.png"
	object2ImageFile[base.OBJ_Tree] = "assets/img/tree.png"
	object2ImageFile[base.OBJ_Box] = "assets/img/box.png"

	Object2Image = make(map[base.ObjectType]*ebiten.Image)
	for k, v := range object2ImageFile {
		img, _, err := ebitenutil.NewImageFromFile(v)
		if err != nil {
			log.Fatal(err)
		}
		Object2Image[k] = img
	}
}
