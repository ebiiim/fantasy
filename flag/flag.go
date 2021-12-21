package flag

import "github.com/ebiiim/fantasy/base"

type Flag uint

const (
	BlockLand Flag = 1 << iota
	BlockSea
	BlockSky
	IsHigh
	IsCharacter
	IsMe
	IsPlayer
	IsNPC
	HasAction
)

const (
	None       Flag = 0
	All             = ^None
	BlockAll        = BlockLand | BlockSky | BlockSea
	Land            = BlockSea
	Sea             = BlockLand
	Sky             = IsHigh
	LandObject      = Land | BlockLand
	SeaObject       = Sea | BlockSea
	SkyObject       = Sky | BlockSky
	Me              = IsMe | IsPlayer | IsCharacter
	Player          = IsPlayer | IsCharacter | HasAction
	NPC             = IsNPC | IsCharacter | HasAction | LandObject
	ItemBox         = HasAction | LandObject
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
