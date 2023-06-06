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

func (d *Door) Open() {
	d.State = Open
}

func (d *Door) Close() {
	d.State = Closed
}

func (d *Door) Crash() {
	switch d.Type {
	case Glass:
		d.State = Crashed
	default:
		fmt.Println(msg.ErrCannot)
	}
}

func (d *Door) Unlock() {
	switch d.Type {
	case Locked:
		d.State = Closed
	default:
		fmt.Println(msg.ErrCannot)
	}
}

func (d *Door) PrintStatus() {
	switch d.State {
	case Open:
		fmt.Printf("%s이 열려있습니다. 지나갈 수 있습니다.\n", d.Name)
	case Closed:
		fmt.Printf("%s이 닫혀있습니다. 지나갈 수 없습니다.\n", d.Name)
	case Locked:
		fmt.Printf("%s이 잠겨있습니다. 열쇠가 필요합니다.\n", d.Name)
	case Crashed:
		fmt.Printf("%s이 부숴져있습니다. 지나갈 수 있습니다.\n", d.Name)
	}
}
