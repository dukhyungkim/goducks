package monster

type Monster struct {
	Name          string
	CurrentHealth int
	MaxHealth     int
	Attack        int
	Defense       int
}

func NewSquirrel() Monster {
	return Monster{
		Name:          "다람쥐",
		CurrentHealth: 50,
		MaxHealth:     50,
		Attack:        5,
		Defense:       0,
	}
}

func NewRabbit() Monster {
	return Monster{
		Name:          "토끼",
		CurrentHealth: 70,
		MaxHealth:     70,
		Attack:        7,
		Defense:       3,
	}
}

func NewDeer() Monster {
	return Monster{
		Name:          "사슴",
		CurrentHealth: 100,
		MaxHealth:     100,
		Attack:        10,
		Defense:       5,
	}
}
