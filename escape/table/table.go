package table

import (
	"math/rand"
)

type DropItem struct {
	Name  string
	Count int
	Range float32
}

type DropTable struct {
	list []DropItem
}

var (
	SquirrelDropTable DropTable
	RabbitDropTable   DropTable
	DeerDropTable     DropTable
	BoxDropTable      DropTable
)

func init() {
	SquirrelDropTable.list = calcDropTable(squirrelDropInfo)
	RabbitDropTable.list = calcDropTable(rabbitDropInfo)
	DeerDropTable.list = calcDropTable(deerDropInfo)
	BoxDropTable.list = calcDropTable(boxDropInfo)
}

func calcDropTable(dropInfo []dropInfo) []DropItem {
	var totalWeight float32
	for i := range dropInfo {
		totalWeight += dropInfo[i].Weight
	}

	var sum float32
	var dropItems []DropItem
	for _, info := range dropInfo {
		chance := (info.Weight / totalWeight) * 100
		sum += chance
		dropItem := DropItem{
			Name:  info.Name,
			Count: info.Count,
			Range: sum,
		}
		dropItems = append(dropItems, dropItem)
	}
	return dropItems
}

func (table *DropTable) GetItem() (string, int) {
	choice := rand.Float32() * (100 - 1)
	for _, item := range table.list {
		if choice <= item.Range {
			return item.Name, item.Count
		}
	}
	return "", 0
}
