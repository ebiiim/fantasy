package base

var GetFlagSet map[ObjectType]FlagSet

const (
	fBlockLand FlagSet = 1 << iota
	fBlockSea
	fBlockSky
	fCharacter
	fMe
	fNPC
	fPlayer
	fActionObj
)

const (
	FNone       FlagSet = 0
	FAll                = ^FNone
	FBlockAll           = fBlockLand | fBlockSky | fBlockSea
	FLandObject         = fBlockLand | fBlockSea
	FSeaObject          = fBlockLand | fBlockSea
	FSkyObject          = fBlockSky
	FMe                 = fMe | fPlayer | fCharacter | FLandObject
	FNPC                = fNPC | fCharacter | fActionObj | FLandObject
	FPlayer             = fPlayer | fCharacter | fActionObj | FLandObject
	FBox                = fActionObj | FLandObject
)

func init() {
	GetFlagSet = initRegisterFlags()
}
