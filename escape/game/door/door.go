package door

import (
	"escape/game/direct"
	"escape/msg"
	"fmt"
)

type Door struct {
	Name      string
	Type      DoorType
	State     DoorState
	Direction direct.Direction
}

type DoorType int8

const (
	Glass DoorType = iota + 1
	Wood
	Lock
)

type DoorState int8

const (
	Open = iota + 1
	Closed
	Locked
	Crashed
)

func NewGlassDoor(d direct.Direction) *Door {
	return &Door{Name: "유리문", Type: Glass, State: Closed, Direction: d}
}

func NewWoodDoor(d direct.Direction) *Door {
	return &Door{Name: "나무문", Type: Wood, State: Closed, Direction: d}
}

func NewLockedDoor(d direct.Direction) *Door {
	return &Door{Name: "잠긴문", Type: Lock, State: Locked, Direction: d}
}

func OpenDoor(door *Door) {
	door.State = Open
}

func CloseDoor(door *Door) {
	door.State = Closed
}

func CrashDoor(door *Door) {
	switch door.Type {
	case Glass:
		door.State = Crashed
	default:
		fmt.Println(msg.ErrCannot)
	}
}

func UnlockDoor(door *Door) {
	switch door.Type {
	case Locked:
		door.State = Closed
	default:
		fmt.Println(msg.ErrCannot)
	}
}
