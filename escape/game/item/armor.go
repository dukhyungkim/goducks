package item

type Armor struct {
	Name    string
	Type    ArmorType
	Defense int
}

type ArmorType string

const (
	Top    ArmorType = "상의"
	Bottom ArmorType = "하의"
	Foot   ArmorType = "신발"
)

func NewLeatherHat() Armor {
	return Armor{Name: "가죽 옷", Type: Top, Defense: 6}
}

func NewLeatherPants() Armor {
	return Armor{Name: "가죽 바지", Type: Bottom, Defense: 4}
}

func NewLeatherShoes() Armor {
	return Armor{Name: "가죽 신발", Type: Foot, Defense: 3}
}
