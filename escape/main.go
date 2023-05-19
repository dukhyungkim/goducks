package main

import (
	"escape/game"
	"escape/game/player"
	"escape/stages/stage1"
	"escape/stages/stage2"
	"fmt"
)

func main() {
	username := game.GetUserName()
	p := player.NewPlayer(username)
	fmt.Printf("%s님 환영합니다.\n\n", p.Name)

	stage1.GoToStage(&p)
	stage2.GoToStage(&p)
}
