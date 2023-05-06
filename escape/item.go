package main

import "fmt"

type Item string

const (
	NoItem Item = ""
	Key    Item = "열쇠"
	Hammer Item = "망치"
)

var items = [MaxXSize][MaxYSize]Item{}

func InitItems() {
	items[0][3] = Hammer
	items[5][0] = Key
}

func UseItem(item Item, door Door) {
	if item == NoItem || door == NoDoor {
		PrintWrongInput()
		return
	}

	switch item {
	case Hammer:
		if door != GlassDoor {
			PrintWrongUsage(item, door)
			return
		}
	case Key:
		if door != LockedDoor {
			PrintWrongInput()
			return
		}
	}
	x, y := GetCurrentXY()
	ChangeDoorState(door, x, y)
	RemoveItem(item)
}

func PrintItemInfo() {
	x, y := GetCurrentXY()
	if item := items[x][y]; item != NoItem {
		fmt.Println("떨어진 아이템:", item)
	}
}

func IsItem(s string) bool {
	switch Item(s) {
	case Hammer, Key:
		return true
	default:
		return false
	}
}

func RemoveItemInRoom(item Item) {
	x, y := GetCurrentXY()
	if items[x][y] == item {
		items[x][y] = NoItem
	}
}
