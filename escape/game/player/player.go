package player

import (
	"escape/game/direction"
	"escape/game/item"
	"fmt"
)

type Player struct {
	Name          string
	CurrentHealth int
	MaxHealth     int
	Attack        int
	Defense       int
	Inventory     Inventory
	Coordinates   direction.Coordinates
	Equipment     Equipment
}

type Inventory struct {
	Tools   [10]item.Tool
	Weapons [10]item.Weapon
	Armors  [10]item.Armor
}

type Equipment struct {
	Top       item.Armor
	Bottom    item.Armor
	Shoes     item.Armor
	LeftHand  item.Weapon
	RightHand item.Weapon
}

func NewPlayer(name string) Player {
	return Player{
		Name:          name,
		CurrentHealth: 30,
		MaxHealth:     30,
		Attack:        3,
		Defense:       0,
		Inventory:     Inventory{},
		Equipment:     Equipment{},
	}
}

func CurrentPosition(p Player) (x, y int) {
	return p.Coordinates.X, p.Coordinates.Y
}

func SetPosition(p *Player, x, y int) {
	p.Coordinates.X = x
	p.Coordinates.Y = y
}

func FindInventory(p Player, toolName string) (item.Tool, bool) {
	for _, tool := range p.Inventory.Tools {
		if tool.Name == toolName {
			return tool, true
		}
	}
	return item.Tool{}, false
}

func PutItemToInventory(p *Player, tool item.Tool) {
	for i := range p.Inventory.Tools {
		if p.Inventory.Tools[i].Name == "" {
			p.Inventory.Tools[i] = tool
			fmt.Printf("%s을(를) 소지품에 넣었습니다.\n", tool.Name)
			return
		}
	}
}

func PutWeaponToInventory(p *Player, weapon item.Weapon) {
	for i := range p.Inventory.Weapons {
		if p.Inventory.Weapons[i].Name == "" {
			p.Inventory.Weapons[i] = weapon
			fmt.Printf("%s을(를) 소지품에 넣었습니다.\n", weapon.Name)
			return
		}
	}
}

func PutArmorToInventory(p *Player, armor item.Armor) {
	for i := range p.Inventory.Armors {
		if p.Inventory.Armors[i].Name == "" {
			p.Inventory.Armors[i] = armor
			fmt.Printf("%s을(를) 소지품에 넣었습니다.\n", armor.Name)
			return
		}
	}
}

func RemoveTool(p *Player, tool item.Tool) {
	for i := range p.Inventory.Tools {
		if p.Inventory.Tools[i].Name == tool.Name {
			p.Inventory.Tools[i] = item.Tool{}
			fmt.Printf("%s가 소지품에서 없어졌습니다.\n", tool.Name)
			return
		}
	}
}

func RemoveWeapon(player *Player, weapon item.Weapon) {
	for i := range player.Inventory.Weapons {
		if player.Inventory.Weapons[i].Name == weapon.Name {
			player.Inventory.Weapons[i] = item.Weapon{}
			fmt.Printf("%s가 소지품에서 없어졌습니다.\n", weapon.Name)
			return
		}
	}
}

func RemoveArmor(p *Player, armor item.Armor) {
	for i := range p.Inventory.Armors {
		if p.Inventory.Armors[i].Name == armor.Name {
			p.Inventory.Armors[i] = item.Armor{}
			fmt.Printf("%s가 소지품에서 없어졌습니다.\n", armor.Name)
			return
		}
	}
}

func PrintInventory(player Player) {
	fmt.Print("갖고있는 물건들: ")
	for _, tool := range player.Inventory.Tools {
		if tool.Name == "" {
			continue
		}
		fmt.Print(tool.Name + " ")
	}
	for _, weapon := range player.Inventory.Weapons {
		if weapon.Name == "" {
			continue
		}
		fmt.Print(weapon.Name + " ")
	}
	for _, armor := range player.Inventory.Armors {
		if armor.Name == "" {
			continue
		}
		fmt.Print(armor.Name + " ")
	}
	fmt.Println()
}
