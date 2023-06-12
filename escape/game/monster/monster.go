package monster

import (
	"escape/game/player"
	"escape/table"
	"escape/util/hangul"
	"fmt"
	"os"
)

type Monster struct {
	Name          string
	CurrentHealth int
	MaxHealth     int
	Attack        int
	Defense       int
	dropTable     table.DropTable
}

func NewSquirrel() *Monster {
	return &Monster{
		Name:          "다람쥐",
		CurrentHealth: 50,
		MaxHealth:     50,
		Attack:        5,
		Defense:       0,
		dropTable:     table.SquirrelDropTable,
	}
}

func NewRabbit() *Monster {
	return &Monster{
		Name:          "토끼",
		CurrentHealth: 70,
		MaxHealth:     70,
		Attack:        7,
		Defense:       3,
		dropTable:     table.RabbitDropTable,
	}
}

func NewDeer() *Monster {
	return &Monster{
		Name:          "사슴",
		CurrentHealth: 100,
		MaxHealth:     100,
		Attack:        10,
		Defense:       5,
		dropTable:     table.DeerDropTable,
	}
}

func (m *Monster) AttackPlayer(p *player.Player) {
	fmt.Printf("%s 당신을 공격합니다.\n", hangul.WithJosa(m.Name, hangul.LeeGa))
	fmt.Printf("%s 당신에게 %d의 피해를 입혔습니다.\n", hangul.WithJosa(m.Name, hangul.LeeGa), m.Attack)
	p.CurrentHealth -= m.Attack + p.Defense
	if p.CurrentHealth <= 0 {
		fmt.Println("사망하였습니다.")
		os.Exit(44)
	}
}

func (m *Monster) DropItem() (itemName string, itemCount int) {
	return m.dropTable.GetItem()
}
