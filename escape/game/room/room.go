package room

import (
	"escape/game/door"
	"escape/game/item"
	"escape/game/monster"
)

type Room struct {
	Door    *door.Door
	Weapons []*item.Weapon
	Armors  []*item.Armor
	Tools   []*item.Tool
	Monster *monster.Monster
	ItemBox *item.Box
}

func RemoveWeapon(room *Room, weapon item.Weapon) {
	for i, t := range room.Weapons {
		if t == nil {
			continue
		}
		if t.Name == weapon.Name {
			room.Weapons = append(room.Weapons[:i], room.Weapons[i+1:]...)
			return
		}
	}
}

func RemoveArmor(room *Room, armor item.Armor) {
	for i, t := range room.Armors {
		if t == nil {
			continue
		}
		if t.Name == armor.Name {
			room.Armors = append(room.Armors[:i], room.Armors[i+1:]...)
			return
		}
	}
}

func RemoveTool(room *Room, tool item.Tool) {
	for i, t := range room.Tools {
		if t == nil {
			continue
		}
		if t.Name == tool.Name {
			room.Tools = append(room.Tools[:i], room.Tools[i+1:]...)
			return
		}
	}
}
