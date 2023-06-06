package item

import (
	"escape/table"
	"fmt"
)

type Box struct {
	dropTable table.DropTable
}

func NewItemBox() *Box {
	return &Box{
		dropTable: table.BoxDropTable,
	}
}

func (b *Box) Open() (*Weapon, *Armor, []*Tool) {
	itemName, itemCount := b.dropTable.GetItem()
	if itemName == "" && itemCount == 0 {
		fmt.Println("상자에서 아무것도 나오지 않았습니다.")
		return nil, nil, nil
	}

	fmt.Printf("상자에서 %s이(가) %d개 나왔습니다.\n", itemName, itemCount)
	switch itemName {
	case "목검":
		return NewWoodSword(), nil, nil
	case "철검":
		return NewIronSword(), nil, nil
	case "가죽옷":
		return nil, NewLeatherShirt(), nil
	case "가죽바지":
		return nil, NewLeatherPants(), nil
	case "가죽신발":
		return nil, NewLeatherShoes(), nil
	case "회복약":
		var tools []*Tool
		for i := 0; i < itemCount; i++ {
			tools = append(tools, NewPotion())
		}
		return nil, nil, tools
	default:
		panic(fmt.Sprintf("unknown item: %s\n", itemName))
	}
}
