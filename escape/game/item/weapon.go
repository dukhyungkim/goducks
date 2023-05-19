package item

type Weapon struct {
	Name    string
	Type    WeaponType
	Attack  int
	Defense int
}

type WeaponType string

const (
	OneHandSword WeaponType = "한손검"
	Shield       WeaponType = "한손검"
)

func NewWoodSword() Weapon {
	return Weapon{Name: "목검", Type: OneHandSword, Attack: 5}
}

func NewIronSword() Weapon {
	return Weapon{Name: "철검", Type: OneHandSword, Attack: 10}
}

func NewWoodShield() Weapon {
	return Weapon{Name: "나무 방패", Type: Shield, Defense: 10}
}
