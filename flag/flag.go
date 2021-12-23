package flag

import "github.com/ebiiim/fantasy/base"

type Flag uint64

const (
	Undef Flag = 1 << iota

	// E.g., meadow: TrrainLand | TrrainSky
	//       water:  TrrainSea  | TrrainSky
	TerrainLand
	TerrainSea
	TerrainSky

	// Movable objects have this kind of attribute.
	// E.g., dragon: CanOnLand | CanInSea(?) | CanInSky
	// E.g., fish: CanInSea
	CanOnLand
	CanInSea
	CanInSky

	// Each Object sets where it is.
	IsOnLand
	IsInSea
	IsInSky

	IsBlockingObject

	IsCharacter
	IsMe
	IsPlayer
	IsNPC

	HasAction
)

const (
	None Flag = 0
	All       = ^None

	BlockAll  = IsBlockingObject | IsOnLand | IsInSea | IsInSky
	BlockLand = IsBlockingObject | IsOnLand
	BlockSea  = IsBlockingObject | IsInSea
	BlockSky  = IsBlockingObject | IsInSky

	Me      = IsCharacter | IsPlayer | IsMe | CanOnLand
	Player  = IsCharacter | IsPlayer | HasAction | CanOnLand
	NPC     = IsCharacter | IsNPC | HasAction | IsBlockingObject | CanOnLand
	ItemBox = HasAction | IsBlockingObject
)

var m map[base.Object]Flag

func init() {
	m = make(map[base.Object]Flag)
	initData()
}

// Get returns the flag for the giving Object.
// Returns Undef if no flag is found for `obj`.
func Get(obj base.Object) Flag {
	v, ok := m[obj]
	if !ok {
		return Undef
	}
	return v
}

func AnyOf(fs ...Flag) Flag {
	f0 := None
	for _, f := range fs {
		f0 |= f
	}
	return f0
}

func AllOf(fs ...Flag) Flag {
	f0 := All
	for _, f := range fs {
		f0 &= f
	}
	return f0
}

func (f0 Flag) Has(fs ...Flag) bool {
	merged := AnyOf(fs...)
	return f0&merged == merged
}

func (f0 Flag) Excepts(fs ...Flag) bool {
	merged := AnyOf(fs...)
	return ^f0&merged == merged
}
