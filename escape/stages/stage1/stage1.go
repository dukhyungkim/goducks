package stage1

import (
	"escape/game"
	"escape/game/direct"
	"escape/game/door"
	"escape/game/item"
	"escape/game/player"
	"escape/game/room"
	"fmt"
)

func GoToStage(p *player.Player) {
	fmt.Println("스테이지1에 진입하셨습니다.")
	rooms := game.GetRooms()
	InitStage(rooms)
	x, y := StartPosition()
	p.SetPosition(x, y)

	for {
		game.PrintCurrentStatus(*p)
		command := game.GetCommand()
		game.HandleCommand(p, command)
		fmt.Println()
		if IsGoal(*p) {
			fmt.Println("축하합니다. 스테이지1을 클리어하셨습니다.")
			break
		}
	}
}

func InitStage(rooms *[10][10]*room.Room) {
	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[i]); j++ {
			rooms[i][j] = nil
		}
	}
	rooms[0][3] = &room.Room{Tools: []*item.Tool{item.NewHammer()}}
	rooms[1][1] = &room.Room{}
	rooms[1][2] = &room.Room{}
	rooms[1][3] = &room.Room{Door: door.NewGlassDoor(direct.East)}
	rooms[2][3] = &room.Room{Door: door.NewGlassDoor(direct.West)}
	rooms[3][1] = &room.Room{Door: door.NewWoodDoor(direct.East)}
	rooms[3][2] = &room.Room{}
	rooms[3][3] = &room.Room{}
	rooms[3][4] = &room.Room{}
	rooms[4][0] = &room.Room{Tools: []*item.Tool{item.NewKey()}}
	rooms[4][1] = &room.Room{Door: door.NewWoodDoor(direct.West)}
	rooms[4][4] = &room.Room{Door: door.NewWoodDoor(direct.North)}
	rooms[4][5] = &room.Room{Door: door.NewWoodDoor(direct.South)}
	rooms[5][5] = &room.Room{}
	rooms[6][5] = &room.Room{Door: door.NewLockedDoor(direct.East)}
	rooms[7][5] = &room.Room{Door: door.NewLockedDoor(direct.West)}
}

func StartPosition() (x, y int) {
	const startX = 1
	const startY = 1

	return startX, startY
}

func IsGoal(p player.Player) bool {
	const goalX = 7
	const goalY = 5

	x, y := p.CurrentPosition()
	return goalX == x && goalY == y
}
