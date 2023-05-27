package player

import (
	"escape/game/direct"
	"escape/game/item"
	"fmt"
	"strings"
)

type PlayerMode int

const (
	InNormal PlayerMode = iota
	InBattle
)

type Player struct {
	Name          string
	CurrentHealth int
	MaxHealth     int
	Attack        int
	Defense       int
	Mode          PlayerMode
	Inventory     Inventory
	Coordinates   direct.Coordinates
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
		CurrentHealth: 50,
		MaxHealth:     50,
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

func AddWeaponToInventory(p *Player, weapon item.Weapon) {
	for i := range p.Inventory.Weapons {
		if p.Inventory.Weapons[i].Name == "" {
			p.Inventory.Weapons[i] = weapon
			fmt.Printf("%s을(를) 소지품에 넣었습니다.\n", weapon.Name)
			return
		}
	}
}

func AddArmorToInventory(p *Player, armor item.Armor) {
	for i := range p.Inventory.Armors {
		if p.Inventory.Armors[i].Name == "" {
			p.Inventory.Armors[i] = armor
			fmt.Printf("%s을(를) 소지품에 넣었습니다.\n", armor.Name)
			return
		}
	}
}

func AddToolToInventory(p *Player, tool item.Tool) {
	for i := range p.Inventory.Tools {
		if p.Inventory.Tools[i].Name == "" {
			p.Inventory.Tools[i] = tool
			fmt.Printf("%s을(를) 소지품에 넣었습니다.\n", tool.Name)
			return
		}
	}
}

func RemoveToolFromInventory(p *Player, tool item.Tool) {
	for i := range p.Inventory.Tools {
		if p.Inventory.Tools[i].Name == tool.Name {
			p.Inventory.Tools[i] = item.Tool{}
			fmt.Printf("%s(이)가 소지품에서 없어졌습니다.\n", tool.Name)
			return
		}
	}
}

func RemoveWeaponFromInventory(player *Player, weapon item.Weapon) {
	for i := range player.Inventory.Weapons {
		if player.Inventory.Weapons[i].Name == weapon.Name {
			player.Inventory.Weapons[i] = item.Weapon{}
			return
		}
	}
}

func RemoveArmorFromInventory(p *Player, armor item.Armor) {
	for i := range p.Inventory.Armors {
		if p.Inventory.Armors[i].Name == armor.Name {
			p.Inventory.Armors[i] = item.Armor{}
			return
		}
	}
}

func PrintInventory(p Player) {
	weapons := p.Inventory.Weapons
	armors := p.Inventory.Armors
	tools := p.Inventory.Tools

	var outputs []string

	for _, weapon := range weapons {
		if weapon.Name == "" {
			continue
		}
		outputs = append(outputs, weapon.Name)
	}

	for _, armor := range armors {
		if armor.Name == "" {
			continue
		}
		outputs = append(outputs, armor.Name)
	}

	for _, tool := range tools {
		if tool.Name == "" {
			continue
		}
		outputs = append(outputs, tool.Name)
	}

	fmt.Println("소지품:", strings.Join(outputs, ", "))
}

func PrintEquipments(player Player) {
	fmt.Println("=== 착용중인 장비 ===")
	fmt.Printf("[ 상의]: %s\n", player.Equipment.Top.Name)
	fmt.Printf("[ 하의]: %s\n", player.Equipment.Bottom.Name)
	fmt.Printf("[   발]: %s\n", player.Equipment.Shoes.Name)
	fmt.Printf("[ 왼손]: %s\n", player.Equipment.LeftHand.Name)
	fmt.Printf("[오른손]: %s\n", player.Equipment.RightHand.Name)
}

func PrintStatus(player Player) {
	fmt.Println("=== 플레이어 정보 ===")
	fmt.Printf("[이름]: %s\n", player.Name)
	fmt.Printf("[체력]: %d / %d\n", player.CurrentHealth, player.MaxHealth)
	fmt.Printf("[오른손 공격력]: %d\n", player.Equipment.RightHand.Attack+player.Attack)
	fmt.Printf("[왼손 공격력]: %d\n", player.Equipment.LeftHand.Attack+player.Attack)
	fmt.Printf("[방어력]: %d\n", player.Defense)
	fmt.Printf("[위치]: (%d, %d)\n", player.Coordinates.X, player.Coordinates.Y)
}

func EquipWeapon(player *Player, weapon item.Weapon) {
	if player.Equipment.RightHand.Name == "" {
		player.Equipment.RightHand = weapon
		RemoveWeaponFromInventory(player, weapon)
		fmt.Printf("%s을(를) 오른손에 쥐었습니다.", weapon.Name)
		return
	}
	if player.Equipment.LeftHand.Name == "" {
		player.Equipment.LeftHand = weapon
		RemoveWeaponFromInventory(player, weapon)
		fmt.Printf("%s을(를) 왼손에 쥐었습니다.", weapon.Name)
		return
	}
	fmt.Println("더이상 착용할 수 없습니다.")
}

func UnEquipWeapon(player *Player, weaponName string) {
	if player.Equipment.RightHand.Name == weaponName {
		player.Equipment.RightHand = item.Weapon{}
		fmt.Printf("오른손의 %s을(를) 내려놓았습니다.", weaponName)
		return
	}
	if player.Equipment.LeftHand.Name == weaponName {
		player.Equipment.LeftHand = item.Weapon{}
		fmt.Printf("왼손의 %s을(를) 내려놓았습니다.", weaponName)
		return
	}
	fmt.Println("그런 장비는 없습니다.")
}

func EquipArmor(player *Player, armor item.Armor) {
	switch armor.Type {
	case item.Top:
		if player.Equipment.Top.Name != "" {
			fmt.Printf("이미 %s을(를) 입고있습니다.\n", player.Equipment.Top.Name)
			return
		}
		player.Equipment.Top = armor
		player.Defense += armor.Defense
		RemoveArmorFromInventory(player, armor)
		fmt.Printf("%s을(를) 입었습니다.\n", armor.Name)
	case item.Bottom:
		if player.Equipment.Bottom.Name != "" {
			fmt.Printf("이미 %s을(를) 입고있습니다.\n", player.Equipment.Bottom.Name)
			return
		}
		player.Equipment.Bottom = armor
		player.Defense += armor.Defense
		RemoveArmorFromInventory(player, armor)
		fmt.Printf("%s을(를) 입었습니다.\n", armor.Name)
	case item.Shoes:
		if player.Equipment.Shoes.Name != "" {
			fmt.Printf("이미 %s을(를) 입고있습니다.\n", player.Equipment.Shoes.Name)
			return
		}
		player.Equipment.Shoes = armor
		player.Defense += armor.Defense
		RemoveArmorFromInventory(player, armor)
		fmt.Printf("%s을(를) 신었습니다.\n", armor.Name)
	}
}

func UnEquipArmor(player *Player, armorName string) {
	switch armorName {
	case player.Equipment.Top.Name:
		armor := player.Equipment.Top
		player.Equipment.Top = item.Armor{}
		player.Defense -= armor.Defense
		AddArmorToInventory(player, armor)
		fmt.Printf("%s을(를) 벗었습니다.\n", armor.Name)
		return
	case player.Equipment.Bottom.Name:
		armor := player.Equipment.Bottom
		player.Equipment.Bottom = item.Armor{}
		player.Defense -= armor.Defense
		AddArmorToInventory(player, armor)
		fmt.Printf("%s을(를) 벗었습니다.\n", armor.Name)
		return
	case player.Equipment.Shoes.Name:
		armor := player.Equipment.Shoes
		player.Equipment.Shoes = item.Armor{}
		player.Defense -= armor.Defense
		AddArmorToInventory(player, armor)
		fmt.Printf("%s을(를) 벗었습니다.\n", armor.Name)
		return
	}
	fmt.Println("그런 장비는 없습니다.")
}

func FindTool(p Player, toolName string) (item.Tool, bool) {
	for _, tool := range p.Inventory.Tools {
		if tool.Name == toolName {
			return tool, true
		}
	}
	return item.Tool{}, false
}
