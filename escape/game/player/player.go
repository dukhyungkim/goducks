package player

import (
	"escape/game/direct"
	"escape/game/item"
	"escape/util/hangul"
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

func (p *Player) CurrentPosition() (x, y int) {
	return p.Coordinates.X, p.Coordinates.Y
}

func (p *Player) SetPosition(x, y int) {
	p.Coordinates.X = x
	p.Coordinates.Y = y
}

func (p *Player) FindInventory(toolName string) (item.Tool, bool) {
	for _, tool := range p.Inventory.Tools {
		if tool.Name == toolName {
			return tool, true
		}
	}
	return item.Tool{}, false
}

func (p *Player) AddWeaponToInventory(weapon item.Weapon) {
	for i := range p.Inventory.Weapons {
		if p.Inventory.Weapons[i].Name == "" {
			p.Inventory.Weapons[i] = weapon
			fmt.Printf("%s 소지품에 넣었습니다.\n", hangul.WithJosa(weapon.Name, hangul.EulLul))
			return
		}
	}
}

func (p *Player) AddArmorToInventory(armor item.Armor) {
	for i := range p.Inventory.Armors {
		if p.Inventory.Armors[i].Name == "" {
			p.Inventory.Armors[i] = armor
			fmt.Printf("%s 소지품에 넣었습니다.\n", hangul.WithJosa(armor.Name, hangul.EulLul))
			return
		}
	}
}

func (p *Player) AddToolToInventory(tool item.Tool) {
	for i := range p.Inventory.Tools {
		if p.Inventory.Tools[i].Name == "" {
			p.Inventory.Tools[i] = tool
			fmt.Printf("%s 소지품에 넣었습니다.\n", hangul.WithJosa(tool.Name, hangul.EulLul))
			return
		}
	}
}

func (p *Player) RemoveToolFromInventory(tool item.Tool) {
	for i := range p.Inventory.Tools {
		if p.Inventory.Tools[i].Name == tool.Name {
			p.Inventory.Tools[i] = item.Tool{}
			fmt.Printf("%s 소지품에서 없어졌습니다.\n", hangul.WithJosa(tool.Name, hangul.LeeGa))
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

func (p *Player) PrintInventory() {
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

func (p *Player) PrintEquipments() {
	fmt.Println("=== 착용중인 장비 ===")
	fmt.Printf("[ 상의]: %s\n", p.Equipment.Top.Name)
	fmt.Printf("[ 하의]: %s\n", p.Equipment.Bottom.Name)
	fmt.Printf("[   발]: %s\n", p.Equipment.Shoes.Name)
	fmt.Printf("[ 왼손]: %s\n", p.Equipment.LeftHand.Name)
	fmt.Printf("[오른손]: %s\n", p.Equipment.RightHand.Name)
}

func (p *Player) PrintStatus() {
	fmt.Println("=== 플레이어 정보 ===")
	fmt.Printf("[이름]: %s\n", p.Name)
	fmt.Printf("[체력]: %d / %d\n", p.CurrentHealth, p.MaxHealth)
	fmt.Printf("[오른손 공격력]: %d\n", p.Equipment.RightHand.Attack+p.Attack)
	fmt.Printf("[왼손 공격력]: %d\n", p.Equipment.LeftHand.Attack+p.Attack)
	fmt.Printf("[방어력]: %d\n", p.Defense)
	fmt.Printf("[위치]: (%d, %d)\n", p.Coordinates.X, p.Coordinates.Y)
}

func (p *Player) EquipWeapon(weapon item.Weapon) {
	if p.Equipment.RightHand.Name == "" {
		p.Equipment.RightHand = weapon
		RemoveWeaponFromInventory(p, weapon)
		fmt.Printf("%s 오른손에 쥐었습니다.", hangul.WithJosa(weapon.Name, hangul.EulLul))
		return
	}
	if p.Equipment.LeftHand.Name == "" {
		p.Equipment.LeftHand = weapon
		RemoveWeaponFromInventory(p, weapon)
		fmt.Printf("%s 왼손에 쥐었습니다.", hangul.WithJosa(weapon.Name, hangul.EulLul))
		return
	}
	fmt.Println("더이상 착용할 수 없습니다.")
}

func (p *Player) UnEquipWeapon(weaponName string) {
	if p.Equipment.RightHand.Name == weaponName {
		p.Equipment.RightHand = item.Weapon{}
		fmt.Printf("오른손의 %s 내려놓았습니다.", hangul.WithJosa(weaponName, hangul.EulLul))
		return
	}
	if p.Equipment.LeftHand.Name == weaponName {
		p.Equipment.LeftHand = item.Weapon{}
		fmt.Printf("왼손의 %s 내려놓았습니다.", hangul.WithJosa(weaponName, hangul.EulLul))
		return
	}
	fmt.Println("그런 장비는 없습니다.")
}

func (p *Player) EquipArmor(armor item.Armor) {
	switch armor.Type {
	case item.Top:
		if p.Equipment.Top.Name != "" {
			fmt.Printf("이미 %s 입고있습니다.\n", hangul.WithJosa(p.Equipment.Top.Name, hangul.EulLul))
			return
		}
		p.Equipment.Top = armor
		p.Defense += armor.Defense
		RemoveArmorFromInventory(p, armor)
		fmt.Printf("%s 입었습니다.\n", hangul.WithJosa(armor.Name, hangul.EulLul))
	case item.Bottom:
		if p.Equipment.Bottom.Name != "" {
			fmt.Printf("이미 %s 입고있습니다.\n", hangul.WithJosa(p.Equipment.Bottom.Name, hangul.EulLul))
			return
		}
		p.Equipment.Bottom = armor
		p.Defense += armor.Defense
		RemoveArmorFromInventory(p, armor)
		fmt.Printf("%s 입었습니다.\n", hangul.WithJosa(armor.Name, hangul.EulLul))
	case item.Shoes:
		if p.Equipment.Shoes.Name != "" {
			fmt.Printf("이미 %s 입고있습니다.\n", hangul.WithJosa(p.Equipment.Shoes.Name, hangul.EulLul))
			return
		}
		p.Equipment.Shoes = armor
		p.Defense += armor.Defense
		RemoveArmorFromInventory(p, armor)
		fmt.Printf("%s 신었습니다.\n", hangul.WithJosa(armor.Name, hangul.EulLul))
	}
}

func (p *Player) UnEquipArmor(armorName string) {
	switch armorName {
	case p.Equipment.Top.Name:
		armor := p.Equipment.Top
		p.Equipment.Top = item.Armor{}
		p.Defense -= armor.Defense
		p.AddArmorToInventory(armor)
		fmt.Printf("%s 벗었습니다.\n", hangul.WithJosa(armor.Name, hangul.EulLul))
		return
	case p.Equipment.Bottom.Name:
		armor := p.Equipment.Bottom
		p.Equipment.Bottom = item.Armor{}
		p.Defense -= armor.Defense
		p.AddArmorToInventory(armor)
		fmt.Printf("%s 벗었습니다.\n", hangul.WithJosa(armor.Name, hangul.EulLul))
		return
	case p.Equipment.Shoes.Name:
		armor := p.Equipment.Shoes
		p.Equipment.Shoes = item.Armor{}
		p.Defense -= armor.Defense
		p.AddArmorToInventory(armor)
		fmt.Printf("%s 벗었습니다.\n", hangul.WithJosa(armor.Name, hangul.EulLul))
		return
	}
	fmt.Println("그런 장비는 없습니다.")
}

func (p *Player) FindTool(toolName string) (item.Tool, bool) {
	for _, tool := range p.Inventory.Tools {
		if tool.Name == toolName {
			return tool, true
		}
	}
	return item.Tool{}, false
}
