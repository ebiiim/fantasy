package base

const (
	OBJ_Err   ObjectType = iota // 0
	OBJ_None                    // 1
	OBJ_Base                    // 2
	OBJ_Grass                   // 3
	OBJ_Tree                    // 4
	OBJ_Box                     // 5

	OBJ_Me ObjectType = -1
	OBJ_BG ObjectType = -99999999
)
