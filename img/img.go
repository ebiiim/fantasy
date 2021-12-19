package img

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/ebiiim/fantasy/base"
)

var Object2Image map[base.ObjectType]*ebiten.Image

func init() {
	log.Println("load images")
	var (
		Me    *ebiten.Image
		Err   *ebiten.Image
		None  *ebiten.Image
		Base  *ebiten.Image
		Grass *ebiten.Image
		Tree  *ebiten.Image
		Box   *ebiten.Image
	)
	var err error
	Me, _, err = ebitenutil.NewImageFromFile("assets/img/me.png")
	if err != nil {
		log.Fatal(err)
	}
	Err, _, err = ebitenutil.NewImageFromFile("assets/img/err.png")
	if err != nil {
		log.Fatal(err)
	}
	None, _, err = ebitenutil.NewImageFromFile("assets/img/none.png")
	if err != nil {
		log.Fatal(err)
	}
	Grass, _, err = ebitenutil.NewImageFromFile("assets/img/grass.png")
	if err != nil {
		log.Fatal(err)
	}
	Tree, _, err = ebitenutil.NewImageFromFile("assets/img/tree.png")
	if err != nil {
		log.Fatal(err)
	}
	Base, _, err = ebitenutil.NewImageFromFile("assets/img/base.png")
	if err != nil {
		log.Fatal(err)
	}
	Box, _, err = ebitenutil.NewImageFromFile("assets/img/box.png")
	if err != nil {
		log.Fatal(err)
	}

	Object2Image = make(map[base.ObjectType]*ebiten.Image)
	Object2Image[base.OBJ_Err] = Err
	Object2Image[base.OBJ_Me] = Me
	Object2Image[base.OBJ_None] = None
	Object2Image[base.OBJ_Base] = Base
	Object2Image[base.OBJ_Grass] = Grass
	Object2Image[base.OBJ_Tree] = Tree
	Object2Image[base.OBJ_Box] = Box
}
