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

func (r *Room) RemoveWeapon(weapon item.Weapon) {
	for i, t := range r.Weapons {
		if t == nil {
			continue
		}
		if t.Name == weapon.Name {
			r.Weapons = append(r.Weapons[:i], r.Weapons[i+1:]...)
			return
		}
	}
}

func (r *Room) RemoveArmor(armor item.Armor) {
	for i, t := range r.Armors {
		if t == nil {
			continue
		}
		if t.Name == armor.Name {
			r.Armors = append(r.Armors[:i], r.Armors[i+1:]...)
			return
		}
	}
}

func (r *Room) RemoveTool(tool item.Tool) {
	for i, t := range r.Tools {
		if t == nil {
			continue
		}
		if t.Name == tool.Name {
			r.Tools = append(r.Tools[:i], r.Tools[i+1:]...)
			return
		}
	}
}
