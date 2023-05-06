package main

import "fmt"

var (
	playerX         = 0
	playerY         = 0
	playerInventory [2]Item
)

func InitPlayer() {
	playerX = 1
	playerY = 1
}

func MovePlayer(x, y int) {
	playerX = x
	playerY = y
}

func GetCurrentXY() (x, y int) {
	return playerX, playerY
}

func PutItemToInventory(item Item) {
	for i := range playerInventory {
		if playerInventory[i] == NoItem {
			playerInventory[i] = item
			fmt.Printf("%s를 소지품에 넣었습니다.\n", item)
			return
		}
	}
}

func RemoveItem(item Item) {
	for i := range playerInventory {
		if playerInventory[i] == item {
			playerInventory[i] = NoItem
			fmt.Printf("%s가 소지품에서 없어졌습니다.\n", item)
			return
		}
	}
}

func PrintInventory() {
	fmt.Print("갖고있는 물건들: ")
	for _, item := range playerInventory {
		if item == NoItem {
			continue
		}
		fmt.Print(item + " ")
	}
}
