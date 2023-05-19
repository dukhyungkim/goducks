package item

type Tool struct {
	Name           string
	Type           ToolType
	HealthRecovery int
}

type ToolType int8

const (
	Hammer ToolType = iota + 1
	Key
	Potion
)

func NewHammer() Tool {
	return Tool{Name: "망치", Type: Hammer}
}

func NewKey() Tool {
	return Tool{Name: "열쇠", Type: Key}
}

func NewPotion() Tool {
	return Tool{Name: "회복약", Type: Potion, HealthRecovery: 30}
}
