package main

import "fmt"

type Door string

const (
	NoDoor     Door = ""
	GlassDoor  Door = "유리문"
	WoodDoor   Door = "나무문"
	LockedDoor Door = "잠긴문"
)

type DoorState int

const (
	None DoorState = iota
	Open
	Closed
	Locked
	Crashed
)

var (
	doors      = [MaxXSize][MaxYSize]Door{}
	doorStates = [MaxXSize][MaxYSize]DoorState{}
)

func InitDoors() {
	doors[1][3] = GlassDoor
	doors[2][3] = GlassDoor
	doorStates[1][3] = Closed
	doorStates[2][3] = Closed

	doors[5][4] = WoodDoor
	doors[5][5] = WoodDoor
	doorStates[5][4] = Closed
	doorStates[5][5] = Closed

	doors[7][5] = LockedDoor
	doors[8][5] = LockedDoor
	doorStates[7][5] = Locked
	doorStates[8][5] = Locked
}

func GetDoor(x, y int) Door {
	return doors[x][y]
}

func PrintDoorInfo() {
	x, y := GetCurrentXY()
	door := GetDoor(x, y)
	if door != NoDoor {
		direction := GetDoorDirection(door, x, y)
		fmt.Printf("%s쪽에 %s이 있습니다.\n", direction, door)
		PrintDoorStatus(door, x, y)
	}
}

func GetDoorDirection(door Door, x, y int) Direction {
	for _, direction := range []Direction{
		East, West, South, North,
	} {
		fX, fY := GetFutureXY(direction, x, y)
		if fX >= MaxXSize || fY >= MaxYSize {
			return NoDirection
		}
		if GetDoor(fX, fY) == door {
			return direction
		}
	}
	return NoDirection
}

func OpenDoor(door Door, x, y int) {
	switch GetDoorStatus(x, y) {
	case Closed:
		ChangeDoorState(door, x, y)
	case Open:
		fmt.Printf("%s은 이미 열려있습니다.\n", door)
	default:
		PrintCannot()
	}
}

func CloseDoor(door Door, x, y int) {
	switch GetDoorStatus(x, y) {
	case Open:
		ChangeDoorState(door, x, y)
	case Closed:
		fmt.Printf("%s은 이미 닫혀있습니다.\n", door)
	default:
		PrintCannot()
	}
}

func ChangeDoorState(door Door, x, y int) {
	direction := GetDoorDirection(door, x, y)
	fX, fY := GetFutureXY(direction, x, y)

	switch door {
	case GlassDoor:
		doorStates[x][y] = Crashed
		doorStates[fX][fY] = Crashed
	case WoodDoor:
		if doorStates[x][y] == Closed {
			doorStates[x][y] = Open
			doorStates[fX][fY] = Open
		} else {
			doorStates[x][y] = Closed
			doorStates[fX][fY] = Closed
		}
	case LockedDoor:
		if doorStates[x][y] == Locked {
			doorStates[x][y] = Closed
			doorStates[fX][fY] = Closed
			return
		}
		if doorStates[x][y] == Closed {
			doorStates[x][y] = Open
			doorStates[fX][fY] = Open
		} else {
			doorStates[x][y] = Closed
			doorStates[fX][fY] = Closed
		}
	}
}

func GetDoorStatus(x, y int) DoorState {
	if doors[x][y] == NoDoor {
		return None
	}
	return doorStates[x][y]
}

func IsDoor(s string) bool {
	switch Door(s) {
	case GlassDoor, WoodDoor, LockedDoor:
		return true
	default:
		return false
	}
}

func IsDoorExist(x, y int) bool {

	return doors[x][y] != NoDoor
}

func PrintDoorStatus(door Door, x int, y int) {
	if door == NoDoor {
		return
	}

	if doors[x][y] != door {
		return
	}
	switch doorStates[x][y] {
	case Open:
		fmt.Printf("%s이 열려있습니다. 지나갈 수 있습니다.\n", door)
	case Closed:
		fmt.Printf("%s이 닫혀있습니다. 지나갈 수 없습니다.\n", door)
	case Locked:
		fmt.Printf("%s이 잠겨있습니다. 열쇠가 필요합니다.\n", door)
	case Crashed:
		fmt.Printf("%s이 부숴져있습니다. 지나갈 수 있습니다.\n", door)

	}
}
