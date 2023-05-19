package stage2

import (
	"escape/game"
	"escape/game/construction"
	"escape/game/direction"
	"escape/game/item"
	"escape/game/player"
	"fmt"
)

const (
	goalX = 0
	goalY = 7
)

func GoToStage(p *player.Player) {
	fmt.Println("스테이지2에 진입하셨습니다.")
	rooms := game.GetRooms()
	InitStage(rooms)
	x, y := StartPosition()
	player.SetPosition(p, x, y)

	for {
		game.PrintCurrentStatus(*p)
		command := game.GetCommand()
		game.HandleCommand(p, command)
		fmt.Println()
		if IsGoal(*p) {
			fmt.Println("축하합니다. 스테이지2를 클리어하셨습니다.")
			break
		}
	}
}

func InitStage(rooms *[10][10]*construction.Room) {
	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[i]); j++ {
			rooms[i][j] = nil
		}
	}
	rooms[0][3] = &construction.Room{Tools: []item.Tool{item.NewHammer()}}
	rooms[1][1] = &construction.Room{}
	rooms[1][2] = &construction.Room{}
	rooms[1][3] = &construction.Room{Door: construction.NewGlassDoor(direction.East)}
	rooms[2][3] = &construction.Room{Door: construction.NewGlassDoor(direction.West)}
	rooms[3][1] = &construction.Room{Door: construction.NewWoodDoor(direction.East)}
	rooms[3][2] = &construction.Room{}
	rooms[3][3] = &construction.Room{}
	rooms[3][4] = &construction.Room{}
	rooms[4][0] = &construction.Room{Tools: []item.Tool{item.NewKey()}}
	rooms[4][1] = &construction.Room{Door: construction.NewWoodDoor(direction.West)}
	rooms[4][4] = &construction.Room{Door: construction.NewWoodDoor(direction.North)}
	rooms[4][5] = &construction.Room{Door: construction.NewWoodDoor(direction.South)}
	rooms[5][5] = &construction.Room{}
	rooms[6][5] = &construction.Room{Door: construction.NewLockedDoor(direction.East)}
	rooms[7][5] = &construction.Room{Door: construction.NewLockedDoor(direction.West)}
}

func StartPosition() (x, y int) {
	return 1, 1
}

func IsGoal(p player.Player) bool {
	x, y := player.CurrentPosition(p)
	return goalX == x && goalY == y
}
