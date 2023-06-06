package stage2

import (
	"escape/game"
	"escape/game/direct"
	"escape/game/door"
	"escape/game/item"
	"escape/game/monster"
	"escape/game/player"
	"escape/game/room"
	"fmt"
)

func GoToStage(p *player.Player) {
	fmt.Println("스테이지2에 진입하셨습니다.")
	rooms := game.GetRooms()
	InitStage(rooms)
	x, y := StartPosition()
	p.SetPosition(x, y)

	for {
		if p.Mode == player.InBattle {
			command := game.GetBattleCommand()
			game.HandleBattleCommand(p, command)
			continue
		}
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

func InitStage(rooms *[10][10]*room.Room) {
	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[i]); j++ {
			rooms[i][j] = nil
		}
	}

	rooms[0][7] = &room.Room{Door: door.NewLockedDoor(direct.East)}

	rooms[1][3] = &room.Room{Monster: monster.NewDeer(), Door: door.NewGlassDoor(direct.North)}
	rooms[1][4] = &room.Room{Door: door.NewGlassDoor(direct.South)}
	rooms[1][5] = &room.Room{}
	rooms[1][6] = &room.Room{}
	rooms[1][7] = &room.Room{Door: door.NewLockedDoor(direct.West)}

	rooms[2][3] = &room.Room{}

	rooms[3][3] = &room.Room{Door: door.NewGlassDoor(direct.East)}
	rooms[3][6] = &room.Room{ItemBox: item.NewItemBox(), Door: door.NewWoodDoor(direct.East)}

	rooms[4][3] = &room.Room{Monster: monster.NewRabbit(), Tools: []*item.Tool{item.NewHammer()}, Door: door.NewGlassDoor(direct.West)}
	rooms[4][4] = &room.Room{}
	rooms[4][5] = &room.Room{}
	rooms[4][6] = &room.Room{Monster: monster.NewSquirrel(), Door: door.NewWoodDoor(direct.West)}
	rooms[4][7] = &room.Room{}
	rooms[4][8] = &room.Room{}
	rooms[4][9] = &room.Room{ItemBox: item.NewItemBox(), Tools: []*item.Tool{item.NewHammer()}}

	rooms[5][6] = &room.Room{}

	rooms[6][0] = &room.Room{ItemBox: item.NewItemBox(), Door: door.NewWoodDoor(direct.North)}
	rooms[6][1] = &room.Room{Door: door.NewWoodDoor(direct.South)}
	rooms[6][2] = &room.Room{Tools: []*item.Tool{item.NewHammer()}}
	rooms[6][3] = &room.Room{}
	rooms[6][4] = &room.Room{}
	rooms[6][5] = &room.Room{Monster: monster.NewSquirrel()}
	rooms[6][6] = &room.Room{Tools: []*item.Tool{item.NewPotion()}}

	rooms[7][2] = &room.Room{}
	rooms[7][6] = &room.Room{}

	rooms[8][2] = &room.Room{}
	rooms[8][6] = &room.Room{Door: door.NewGlassDoor(direct.East)}

	rooms[9][0] = &room.Room{}
	rooms[9][1] = &room.Room{}
	rooms[9][2] = &room.Room{Weapons: []*item.Weapon{item.NewWoodSword()}}
	rooms[9][6] = &room.Room{ItemBox: item.NewItemBox(), Door: door.NewGlassDoor(direct.West)}
}

func StartPosition() (x, y int) {
	const startX = 9
	const startY = 0

	return startX, startY
}

func IsGoal(p player.Player) bool {
	const goalX = 0
	const goalY = 7

	x, y := p.CurrentPosition()
	return goalX == x && goalY == y
}
