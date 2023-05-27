package direct

import (
	"fmt"
)

type Direction string

const (
	East  Direction = "동"
	West  Direction = "서"
	South Direction = "남"
	North Direction = "북"
)

type Coordinates struct {
	X int
	Y int
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
		panic(fmt.Sprintf("wrong direction: %v", d))
	}
}

func GetAgainstDirect(d Direction) Direction {
	switch d {
	case East:
		return West
	case West:
		return East
	case South:
		return North
	case North:
		return South
	default:
		panic(fmt.Sprintf("wrong direction: %v", d))
	}
}
