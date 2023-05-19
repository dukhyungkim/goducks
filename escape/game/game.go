package game

import (
	"escape/game/construction"
	"escape/game/direction"
	"escape/game/item"
	"escape/game/player"
	"escape/msg"
	"escape/util"
	"fmt"
	"log"
	"strings"
)

const (
	xSize = 10
	ySize = 10
)

var rooms = [xSize][ySize]*construction.Room{}

func GetUserName() string {
	const minLen = 1
	const maxLen = 16

	for {
		fmt.Printf("이름을 입력해주세요. (길이: %d~%d) ", minLen, maxLen)
		input := util.GetUserInput()
		inputLen := len(input)
		if minLen <= inputLen && inputLen <= maxLen {
			return input
		}
		fmt.Println(msg.ErrWrongInput)
	}
}

func GetRooms() *[10][10]*construction.Room {
	return &rooms
}

func CanIGo(p player.Player, d direction.Direction) bool {
	cX, cY := player.CurrentPosition(p)
	fX, fY := direction.GetFutureXY(d, cX, cY)

	if xSize <= fX || ySize <= fY {
		return false
	}
	if fX < 0 || fY < 0 {
		return false
	}

	cRoom := rooms[cX][cY]
	fRoom := rooms[fX][fY]
	if fRoom == nil {
		return false
	}

	if cRoom.Door == nil {
		return true
	}

	if fRoom.Door != nil {
		switch cRoom.Door.State {
		case construction.Open, construction.Crashed:
			return true
		default:
			return false
		}
	}
	return true
}

func PrintCurrentStatus(p player.Player) {
	room := CurrentRoom(p)
	PrintDoorInfo(*room)
	PrintItemInfo(*room)
	PrintPath(p)
}

func PrintDoorInfo(room construction.Room) {
	door := room.Door
	if door == nil {
		return
	}
	fmt.Printf("%s쪽에 %s이 있습니다.\n", door.Direction, door.Name)
	PrintDoorStatus(*door)
}

func PrintDoorStatus(door construction.Door) {
	switch door.State {
	case construction.Open:
		fmt.Printf("%s이 열려있습니다. 지나갈 수 있습니다.\n", door.Name)
	case construction.Closed:
		fmt.Printf("%s이 닫혀있습니다. 지나갈 수 없습니다.\n", door.Name)
	case construction.Locked:
		fmt.Printf("%s이 잠겨있습니다. 열쇠가 필요합니다.\n", door.Name)
	case construction.Crashed:
		fmt.Printf("%s이 부숴져있습니다. 지나갈 수 있습니다.\n", door.Name)
	}
}

func PrintItemInfo(room construction.Room) {
	if len(room.Tools) == 0 {
		return
	}
	fmt.Print("떨어진 아이템: ")
	for i, tool := range room.Tools {
		fmt.Print(tool.Name)
		if i != len(room.Tools)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()
}

func PrintPath(p player.Player) {
	fmt.Print("이동가능한 경로: ")
	for _, d := range []direction.Direction{
		direction.East, direction.West, direction.South, direction.North,
	} {
		if CanIGo(p, d) {
			fmt.Printf("%s ", d)
		}
	}
	fmt.Println()
}

func GetCommand() string {
	fmt.Print(">> ")
	return util.GetUserInput()
}

func HandleCommand(p *player.Player, text string) {
	switch text {
	case string(direction.East), "ㄷ", "e":
		moveTo(p, direction.East)
	case string(direction.West), "ㅅ", "t":
		moveTo(p, direction.West)
	case string(direction.South), "ㄴ", "s":
		moveTo(p, direction.South)
	case string(direction.North), "ㅂ", "q":
		moveTo(p, direction.North)
	case "소지품", "인벤", "인벤토리", "twv", "dq":
		player.PrintInventory(*p)
	default:
		handleLongCommand(p, text)
	}
}

func moveTo(p *player.Player, d direction.Direction) {
	if !CanIGo(*p, d) {
		fmt.Println("이동할 수 없는 곳입니다.")
		return
	}
	x, y := player.CurrentPosition(*p)
	fX, fY := direction.GetFutureXY(d, x, y)
	player.SetPosition(p, fX, fY)
}

func handleLongCommand(p *player.Player, text string) {
	tokens := strings.Split(text, " ")
	switch len(tokens) {
	case 2:
		handleTwoWordsCommand(p, tokens)
	case 3:
		handleThreeWordsCommand(p, tokens)
	default:
		fmt.Println(msg.ErrWrongInput)
	}
}

func handleTwoWordsCommand(p *player.Player, tokens []string) {
	room := CurrentRoom(*p)

	target := tokens[0]
	command := tokens[1]

	if room.Door != nil {
		if room.Door.Name == target {
			handleDoorCommand(p, room.Door, command)
			return
		}
	}

	for _, tool := range room.Tools {
		if tool.Name == target {
			handleItemCommand(p, tool, command)
			return
		}
	}

	fmt.Println(msg.ErrWrongInput)
}

func handleDoorCommand(p *player.Player, door *construction.Door, command string) {
	switch command {
	case "보", "보다", "본다":
		PrintDoorStatus(*door)
	case "열", "열다", "연다":
		switch door.State {
		case construction.Open:
			fmt.Printf("%s은 이미 열려있습니다.\n", door.Name)
			return
		case construction.Closed:
			construction.OpenDoor(door)
			nextDoor := findNextDoor(p, *door)
			construction.OpenDoor(nextDoor)
			fmt.Printf("%s을(를) 열었습니다.\n", door.Name)
		default:
			fmt.Println(msg.ErrCannot)
			return
		}
	case "닫", "닫다", "닫는다":
		switch door.State {
		case construction.Closed:
			fmt.Printf("%s은 이미 닫혀있습니다.\n", door.Name)
			return
		case construction.Open:
			construction.CloseDoor(door)
			nextDoor := findNextDoor(p, *door)
			construction.CloseDoor(nextDoor)
			fmt.Printf("%s을(를) 닫았습니다.\n", door.Name)
		default:
			fmt.Println(msg.ErrCannot)
			return
		}
		if door.State == construction.Closed {
			fmt.Printf("%s은 이미 닫혀있습니다.\n", door.Name)
			return
		}
	default:
		fmt.Println(msg.ErrWrongInput)
	}
}

func handleItemCommand(p *player.Player, item item.Tool, command string) {
	room := CurrentRoom(*p)
	switch command {
	case "줍", "주워", "줍다":
		player.PutItemToInventory(p, item)
		RemoveTool(room, item)
	default:
		fmt.Println(msg.ErrWrongInput)
	}
}

func handleThreeWordsCommand(p *player.Player, tokens []string) {
	command := tokens[2]
	switch command {
	case "사용":
		tool, ok := player.FindInventory(*p, tokens[0])
		if !ok {
			fmt.Println(msg.ErrNotHave)
			return
		}
		room := CurrentRoom(*p)
		if room.Door.Name != tokens[1] {
			fmt.Println(msg.ErrNotFound)
			return
		}
		UseTool(p, tool, room.Door)
	default:
		fmt.Println(msg.ErrWrongInput)
	}
}

func CurrentRoom(p player.Player) *construction.Room {
	x, y := player.CurrentPosition(p)
	return rooms[x][y]
}

func RemoveTool(room *construction.Room, tool item.Tool) {
	for i, t := range room.Tools {
		if t.Name == tool.Name {
			room.Tools[i] = item.Tool{}
		}
	}
}

func UseTool(p *player.Player, tool item.Tool, door *construction.Door) {
	switch tool.Type {
	case item.Hammer:
		if door.Type == construction.Glass {
			construction.CrashDoor(door)
			nextDoor := findNextDoor(p, *door)
			construction.CrashDoor(nextDoor)
			fmt.Printf("%s로 %s을 부쉈습니다.\n", tool.Name, door.Name)
			player.RemoveTool(p, tool)
			return
		}
	case item.Key:
		if door.Type == construction.Locked {
			construction.UnlockDoor(door)
			nextDoor := findNextDoor(p, *door)
			construction.UnlockDoor(nextDoor)
			fmt.Printf("%s로 %s을 열었습니다.\n", tool.Name, door.Name)
			player.RemoveTool(p, tool)
			return
		}
	}
	fmt.Println(msg.ErrCannot)
}

func findNextDoor(p *player.Player, door construction.Door) *construction.Door {
	x, y := player.CurrentPosition(*p)
	fX, fY := direction.GetFutureXY(door.Direction, x, y)
	nextDoor := rooms[fX][fY].Door
	if nextDoor != nil && nextDoor.Name == door.Name {
		return nextDoor
	}
	log.Panicln("cannot find next door")
	return nil
}
