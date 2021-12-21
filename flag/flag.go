package flag

import "github.com/ebiiim/fantasy/base"

type Flag uint

const (
	blockLand Flag = 1 << iota
	blockSea
	blockSky
	character
	me
	npc
	player
	actionObj
)

const (
	None       Flag = 0
	All             = ^None
	BlockAll        = blockLand | blockSky | blockSea
	LandObject      = blockLand | blockSea
	SeaObject       = blockLand | blockSea
	SkyObject       = blockSky
	Me              = me | player | character
	Player          = player | character | actionObj
	NPC             = npc | character | actionObj | LandObject
	Box             = actionObj | LandObject
)

var m map[base.Object]Flag

func init() {
	m = make(map[base.Object]Flag)
	initData()
}

func Get(obj base.Object) Flag {
	v, ok := m[obj]
	if !ok {
		return m[base.ObjUndef]
	}
	return v
}
