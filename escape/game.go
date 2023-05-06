package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	MaxXSize = 9
	MaxYSize = 6

	GoalX = 8
	GoalY = 5
)

type Direction string

const (
	NoDirection Direction = ""
	East        Direction = "동"
	West        Direction = "서"
	South       Direction = "남"
	North       Direction = "북"
)

func InitGame() {
	InitRooms()
	InitDoors()
	InitItems()
	InitPlayer()
}

func PrintCurrentStatus() {
	PrintDoorInfo()
	PrintItemInfo()
	PrintPath()
}

func GetUserInput() string {
	fmt.Print(">> ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func HandleUserInput(text string) {
	var d Direction
	switch text {
	case string(East), "ㄷ", "e":
		d = East
	case string(West), "ㅅ", "t":
		d = West
	case string(South), "ㄴ", "s":
		d = South
	case string(North), "ㅂ", "q":
		d = North
	case "소지품", "인벤", "인벤토리", "twv", "dq":
		PrintInventory()
		return
	default:
		handleLongCommand(text)
		return
	}

	if !CanIGo(d) {
		fmt.Println("이동할 수 없는 곳입니다.")
		return
	}

	x, y := GetCurrentXY()
	fX, fY := GetFutureXY(d, x, y)
	MovePlayer(fX, fY)
}

func IsGoal() bool {
	x, y := GetCurrentXY()
	return GoalX == x && GoalY == y
}

func handleLongCommand(text string) {
	tokens := strings.Split(text, " ")
	switch len(tokens) {
	case 2:
		handleTwoWordsCommand(tokens)
	case 3:
		handleThreeWordsCommand(tokens)
	default:
		PrintWrongInput()
	}
}

func handleTwoWordsCommand(tokens []string) {
	switch {
	case IsDoor(tokens[0]):
		door := Door(tokens[0])
		handleDoorCommand(door, tokens[1])
	case IsItem(tokens[0]):
		item := Item(tokens[0])
		handleItemCommand(item, tokens[1])
	default:
		PrintWrongInput()
	}
}

func handleDoorCommand(door Door, command string) {
	x, y := GetCurrentXY()
	switch command {
	case "보", "보다", "본다":
		PrintDoorStatus(door, x, y)
	case "열", "열다", "연다":
		switch GetDoorStatus(x, y) {
		case Open, Closed:
			ChangeDoorState(door, x, y)
		default:
			PrintCannot()
		}
	default:
		PrintWrongInput()
	}
}

func handleItemCommand(item Item, command string) {
	switch command {
	case "줍", "주워", "줍다":
		PutItemToInventory(item)
		RemoveItemInRoom(item)
	default:
		PrintWrongInput()
	}
}

func handleThreeWordsCommand(tokens []string) {
	command := tokens[2]
	switch command {
	case "사용":
		item := Item(tokens[0])
		door := Door(tokens[1])
		UseItem(item, door)
	default:
		PrintWrongInput()
	}
}

func PrintPath() {
	fmt.Print("이동가능한 경로: ")

	if CanIGo(East) {
		fmt.Printf("%s ", East)
	}
	if CanIGo(West) {
		fmt.Printf("%s ", West)
	}
	if CanIGo(South) {
		fmt.Printf("%s ", South)
	}
	if CanIGo(North) {
		fmt.Printf("%s ", North)
	}
	fmt.Println()
}

func CanIGo(d Direction) bool {
	currentX, currentY := GetCurrentXY()
	futureX, futureY := GetFutureXY(d, currentX, currentY)

	if futureX >= MaxXSize || futureY >= MaxYSize {
		return false
	}
	if futureX < 0 || futureY < 0 {
		return false
	}

	if IsDoorExist(futureX, futureY) {
		if IsDoorExist(currentX, currentY) {
			doorState := GetDoorStatus(futureX, futureY)
			switch doorState {
			case Open, Crashed:
				return true
			default:
				return false
			}
		}
	}
	return rooms[futureX][futureY]
}

func GetFutureXY(d Direction, x, y int) (futureX, futureY int) {
	switch d {
	case East:
		return x + 1, y
	case West:
		return x - 1, y
	case South:
		return x, y - 1
	case North:
		return x, y + 1
	default:
		return x, y
	}
}

func PrintWrongInput() {
	fmt.Println("잘못 입력하셨습니다.")
}

func PrintCannot() {
	fmt.Println("그것은 할 수 없습니다.")
}

func PrintWrongUsage(item Item, door Door) {
	fmt.Printf("%s는 %s에 사용할 수 없습니다.\n", item, door)
}
