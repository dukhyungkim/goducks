package construction

import (
	"escape/game/item"
	"escape/game/monster"
)

type Room struct {
	Door    *Door
	Tools   []item.Tool
	Monster *monster.Monster
	ItemBox *item.Box
}
