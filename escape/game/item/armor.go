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
	Shoes  ArmorType = "신발"
)

func NewLeatherShirt() *Armor {
	return &Armor{Name: "가죽옷", Type: Top, Defense: 6}
}

func NewLeatherPants() *Armor {
	return &Armor{Name: "가죽바지", Type: Bottom, Defense: 4}
}

func NewLeatherShoes() *Armor {
	return &Armor{Name: "가죽신발", Type: Shoes, Defense: 3}
}
