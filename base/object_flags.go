package base

var GetFlagSet map[ObjectType]FlagSet

const (
	f_BlockLand FlagSet = 1 << iota
	f_BlockSea
	f_BlockSky
	f_Character
	f_Me
	f_NPC
	f_Player
	f_ActionObj
)

const (
	F_All        FlagSet = 0xFFFFFFFFFFFFFFFF
	F_None       FlagSet = 0
	F_BlockAll           = f_BlockLand | f_BlockSky | f_BlockSea
	F_LandObject         = f_BlockLand | f_BlockSea
	F_SeaObject          = f_BlockLand | f_BlockSea
	F_SkyObject          = f_BlockSky
	F_Me                 = f_Me | f_Player | f_Character | F_LandObject
	F_NPC                = f_NPC | f_Character | f_ActionObj | F_LandObject
	F_Player             = f_Player | f_Character | f_ActionObj | F_LandObject
	F_Box                = f_ActionObj | F_LandObject
)

func init() {
	GetFlagSet = make(map[ObjectType]FlagSet)
	GetFlagSet[OBJ_Me] = F_Me
	GetFlagSet[OBJ_Err] = F_BlockAll
	GetFlagSet[OBJ_BG] = F_BlockAll
	GetFlagSet[OBJ_None] = F_None
	GetFlagSet[OBJ_Grass] = F_None
	GetFlagSet[OBJ_Tree] = F_LandObject
	GetFlagSet[OBJ_Box] = F_Box
}
