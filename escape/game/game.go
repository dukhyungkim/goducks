package game

import (
	"escape/game/direct"
	"escape/game/door"
	"escape/game/item"
	"escape/game/monster"
	"escape/game/player"
	"escape/game/room"
	"escape/msg"
	"escape/util"
	"fmt"
	"log"
	"math/rand"
	"strings"
)

const (
	xSize = 10
	ySize = 10
)

var rooms = [xSize][ySize]*room.Room{}

func GetUserName() string {
	const minLen = 1
	const maxLen = 16

	for {
		fmt.Printf("이름을 입력해주세요. (길이: %d~%d) ", minLen, maxLen)
		input := util.GetUserInput()
		inputLen := len(input)
		if minLen <= inputLen && inputLen <= maxLen {
			return input
		}
		fmt.Println(msg.ErrWrongInput)
	}
}

func GetRooms() *[10][10]*room.Room {
	return &rooms
}

func CanIGo(p player.Player, d direct.Direction) bool {
	cX, cY := player.CurrentPosition(p)
	fX, fY := direct.GetFutureXY(d, cX, cY)

	if xSize <= fX || ySize <= fY {
		return false
	}
	if fX < 0 || fY < 0 {
		return false
	}

	cRoom := rooms[cX][cY]
	fRoom := rooms[fX][fY]
	if fRoom == nil {
		return false
	}

	if cRoom.Door == nil {
		return true
	}

	if fRoom.Door != nil {
		switch cRoom.Door.State {
		case door.Open, door.Crashed:
			return true
		default:
			return false
		}
	}
	return true
}

func PrintCurrentStatus(p player.Player) {
	r := currentRoom(p)
	printDoorInfo(*r)
	printItemInfo(*r)
	printBoxInfo(*r)
	printMonsterInfo(*r)
	printPath(p)
}

func printDoorInfo(r room.Room) {
	d := r.Door
	if d == nil {
		return
	}
	fmt.Printf("%s쪽에 %s이 있습니다.\n", d.Direction, d.Name)
	printDoorStatus(*d)
}

func printDoorStatus(d door.Door) {
	switch d.State {
	case door.Open:
		fmt.Printf("%s이 열려있습니다. 지나갈 수 있습니다.\n", d.Name)
	case door.Closed:
		fmt.Printf("%s이 닫혀있습니다. 지나갈 수 없습니다.\n", d.Name)
	case door.Locked:
		fmt.Printf("%s이 잠겨있습니다. 열쇠가 필요합니다.\n", d.Name)
	case door.Crashed:
		fmt.Printf("%s이 부숴져있습니다. 지나갈 수 있습니다.\n", d.Name)
	}
}

func printItemInfo(r room.Room) {
	if len(r.Weapons) == 0 && len(r.Armors) == 0 && len(r.Tools) == 0 {
		return
	}

	var outputs []string

	for _, weapon := range r.Weapons {
		outputs = append(outputs, weapon.Name)
	}

	for _, armor := range r.Armors {
		outputs = append(outputs, armor.Name)
	}

	for _, tool := range r.Tools {
		outputs = append(outputs, tool.Name)
	}

	fmt.Println("떨어진 아이템:", strings.Join(outputs, ", "))
}

func printBoxInfo(r room.Room) {
	if r.ItemBox == nil {
		return
	}
	fmt.Println("상자가 놓여있습니다.")
}

func printMonsterInfo(r room.Room) {
	if r.Monster == nil {
		return
	}
	fmt.Printf("%s이(가) 돌아다니고 있습니다.\n", r.Monster.Name)
}

func printPath(p player.Player) {
	directions := findPossibleToMove(p)
	tmpStrs := make([]string, len(directions))
	for i := range directions {
		tmpStrs[i] = string(directions[i])
	}
	fmt.Println("이동가능한 경로:", strings.Join(tmpStrs, ", "))
}

func findPossibleToMove(p player.Player) []direct.Direction {
	directions := []direct.Direction{direct.East, direct.West, direct.South, direct.North}
	result := make([]direct.Direction, 0, len(directions))
	for _, d := range directions {
		if CanIGo(p, d) {
			result = append(result, d)
		}
	}
	return result
}

func GetCommand() string {
	fmt.Print(">> ")
	return util.GetUserInput()
}

func HandleCommand(p *player.Player, text string) {
	switch text {
	case string(direct.East), "ㄷ", "e", "6":
		moveTo(p, direct.East)
	case string(direct.West), "ㅅ", "t", "4":
		moveTo(p, direct.West)
	case string(direct.South), "ㄴ", "s", "2":
		moveTo(p, direct.South)
	case string(direct.North), "ㅂ", "q", "8":
		moveTo(p, direct.North)
	case "소지", "소지품", "인벤", "인벤토리", "tw", "twv", "dq", "thwl", "thwlvna":
		player.PrintInventory(*p)
	case "장비", "wkdql":
		player.PrintEquipments(*p)
	case "정보", "wjdqh":
		player.PrintStatus(*p)
	default:
		handleLongCommand(p, text)
	}
}

func GetBattleCommand() string {
	fmt.Print("!! ")
	return util.GetUserInput()
}

func HandleBattleCommand(p *player.Player, text string) {
	r := currentRoom(*p)
	m := r.Monster
	switch text {
	case "공격", "쳐", "ㅊ", "c":
		fmt.Printf("당신은 %s을(를) 공격합니다.\n", m.Name)
		if r.Monster != nil && p.Equipment.RightHand.Name != "" {
			attack := p.Equipment.RightHand.Attack + p.Attack
			fmt.Printf("당신은 %s에게 %d의 피해를 입혔습니다.\n", m.Name, attack)
			m.CurrentHealth -= attack
			if m.CurrentHealth <= 0 {
				fmt.Printf("%s이(가) 쓰러졌습니다.\n", m.Name)
				itemName, itemCount := monster.DropItem(m)
				fmt.Printf("%s이(가) %d개 떨어졌습니다.\n", itemName, itemCount)
				for i := 0; i < itemCount; i++ {
					switch itemName {
					case "회복약":
						r.Tools = append(r.Tools, item.NewPotion())
					case "열쇠":
						r.Tools = append(r.Tools, item.NewKey())
					}
				}
				r.Monster = nil
				p.Mode = player.InNormal
				fmt.Println()
			}
		}
		if r.Monster != nil && p.Equipment.LeftHand.Name != "" {
			attack := p.Equipment.LeftHand.Attack + p.Attack
			fmt.Printf("당신은 %s에게 %d의 피해를 입혔습니다.\n", m.Name, attack)
			m.CurrentHealth -= attack
			if m.CurrentHealth <= 0 {
				fmt.Printf("%s이(가) 쓰러졌습니다.\n", m.Name)
				itemName, itemCount := monster.DropItem(m)
				fmt.Printf("%s이(가) %d개 떨어졌습니다.\n", itemName, itemCount)
				for i := 0; i < itemCount; i++ {
					switch itemName {
					case "회복약":
						r.Tools = append(r.Tools, item.NewPotion())
					case "열쇠":
						r.Tools = append(r.Tools, item.NewKey())
					}
				}
				r.Monster = nil
				p.Mode = player.InNormal
				fmt.Println()
			}
		}
		if r.Monster != nil && p.Equipment.RightHand.Name == "" && p.Equipment.LeftHand.Name == "" {
			attack := p.Attack
			fmt.Printf("당신은 %s에게 %d의 피해를 입혔습니다.\n", m.Name, attack)
			m.CurrentHealth -= attack
			if m.CurrentHealth <= 0 {
				fmt.Printf("%s이(가) 쓰러졌습니다.\n", m.Name)
				itemName, itemCount := monster.DropItem(m)
				fmt.Printf("%s이(가) %d개 떨어졌습니다.\n", itemName, itemCount)
				for i := 0; i < itemCount; i++ {
					switch itemName {
					case "회복약":
						r.Tools = append(r.Tools, item.NewPotion())
					case "열쇠":
						r.Tools = append(r.Tools, item.NewKey())
					}
				}
				r.Monster = nil
				p.Mode = player.InNormal
				fmt.Println()
			}
		}

		if r.Monster != nil {
			monster.AttackPlayer(m, p)
		}
	case "도망":
		if rand.Int()%2 == 1 {
			fmt.Printf("당신은 %s에게서 성공적으로 도망쳤습니다.\n", m.Name)
			directions := findPossibleToMove(*p)
			to := directions[rand.Int()%len(directions)]
			moveTo(p, to)
			return
		} else {
			fmt.Println("도망칠 수 없습니다!")
			monster.AttackPlayer(m, p)
			return
		}
	case "회복약 사용":
		tool, ok := player.FindTool(*p, "회복약")
		if !ok {
			fmt.Println("당신에게 그런 도구는 없습니다!")
			return
		}

		fmt.Println("당신은 회복약을 사용했습니다.")
		fmt.Printf("체력이 %d 회복되었습니다.\n", tool.HealthRecovery)
		p.CurrentHealth += tool.HealthRecovery
		if p.CurrentHealth > p.MaxHealth {
			p.CurrentHealth = p.MaxHealth
		}
		player.RemoveToolFromInventory(p, tool)

		monster.AttackPlayer(m, p)
	default:
		fmt.Println(msg.ErrWrongInput)
	}
}

func moveTo(p *player.Player, d direct.Direction) {
	if !CanIGo(*p, d) {
		fmt.Println("이동할 수 없는 곳입니다.")
		return
	}
	x, y := player.CurrentPosition(*p)
	fX, fY := direct.GetFutureXY(d, x, y)
	player.SetPosition(p, fX, fY)
}

func handleLongCommand(p *player.Player, text string) {
	tokens := strings.Split(text, " ")
	switch len(tokens) {
	case 2:
		handleTwoWordsCommand(p, tokens)
	case 3:
		handleThreeWordsCommand(p, tokens)
	default:
		fmt.Println(msg.ErrWrongInput)
	}
}

func handleTwoWordsCommand(p *player.Player, tokens []string) {
	r := currentRoom(*p)

	target := tokens[0]
	command := tokens[1]

	if r.Door != nil {
		if r.Door.Name == target {
			handleDoorCommand(p, r.Door, command)
			return
		}
	}

	for _, weapon := range r.Weapons {
		if weapon.Name == target {
			handleWeaponCommand(p, *weapon, command)
			return
		}
	}

	for _, armor := range r.Armors {
		if armor.Name == target {
			handleArmorCommand(p, *armor, command)
			return
		}
	}

	for _, tool := range r.Tools {
		if tool.Name == target {
			handleItemCommand(p, *tool, command)
			return
		}
	}

	if r.ItemBox != nil {
		if target == "상자" {
			handleBoxCommand(r, command)
			return
		}
	}

	for r.Monster != nil {
		if target == r.Monster.Name {
			handleMonsterCommand(r, p, command)
			return
		}
	}

	for _, weapon := range p.Inventory.Weapons {
		if target == weapon.Name {
			handleWeaponCommand(p, weapon, command)
			return
		}
	}

	for _, armor := range p.Inventory.Armors {
		if target == armor.Name {
			handleArmorCommand(p, armor, command)
			return
		}
	}

	for _, tool := range p.Inventory.Tools {
		if target == tool.Name {
			handleItemCommand(p, tool, command)
			return
		}
	}

	fmt.Println(msg.ErrWrongInput)
}

func handleMonsterCommand(r *room.Room, p *player.Player, command string) {
	switch command {
	case "공격", "쳐":
		p.Mode = player.InBattle
		fmt.Printf("당신은 %s와(과) 전투를 시작합니다.\n", r.Monster.Name)
	}
}

func handleBoxCommand(r *room.Room, command string) {
	switch command {
	case "열", "열다", "연다":
		fmt.Println("상자를 열었습니다.")
		weapon, armor, tools := item.OpenBox(r.ItemBox)
		r.ItemBox = nil
		if weapon != nil {
			r.Weapons = append(r.Weapons, weapon)
		}
		if armor != nil {
			r.Armors = append(r.Armors, armor)
		}
		if len(tools) != 0 {
			r.Tools = append(r.Tools, tools...)
		}
	default:
		fmt.Println(msg.ErrCannot)
	}
}

func handleDoorCommand(p *player.Player, d *door.Door, command string) {
	switch command {
	case "보", "보다", "본다":
		printDoorStatus(*d)
	case "열", "열다", "연다":
		switch d.State {
		case door.Open:
			fmt.Printf("%s은 이미 열려있습니다.\n", d.Name)
			return
		case door.Closed:
			door.OpenDoor(d)
			nextDoor := findNextDoor(p, *d)
			door.OpenDoor(nextDoor)
			fmt.Printf("%s을(를) 열었습니다.\n", d.Name)
		default:
			fmt.Println(msg.ErrCannot)
			return
		}
	case "닫", "닫다", "닫는다":
		switch d.State {
		case door.Closed:
			fmt.Printf("%s은 이미 닫혀있습니다.\n", d.Name)
			return
		case door.Open:
			door.CloseDoor(d)
			nextDoor := findNextDoor(p, *d)
			door.CloseDoor(nextDoor)
			fmt.Printf("%s을(를) 닫았습니다.\n", d.Name)
		default:
			fmt.Println(msg.ErrCannot)
			return
		}
		if d.State == door.Closed {
			fmt.Printf("%s은 이미 닫혀있습니다.\n", d.Name)
			return
		}
	default:
		fmt.Println(msg.ErrWrongInput)
	}
}

func handleWeaponCommand(p *player.Player, weapon item.Weapon, command string) {
	r := currentRoom(*p)
	switch command {
	case "줍", "주워":
		player.AddWeaponToInventory(p, weapon)
		room.RemoveWeapon(r, weapon)
	case "착", "착용":
		player.EquipWeapon(p, weapon)
	case "벗", "벗어":
		player.UnEquipWeapon(p, weapon.Name)
	default:
		fmt.Println(msg.ErrWrongInput)
	}
}

func handleArmorCommand(p *player.Player, armor item.Armor, command string) {
	r := currentRoom(*p)
	switch command {
	case "줍", "주워":
		player.AddArmorToInventory(p, armor)
		room.RemoveArmor(r, armor)
	case "입", "입다":
		player.EquipArmor(p, armor)
	case "벗", "벗어":
		player.UnEquipArmor(p, armor.Name)
	default:
		fmt.Println(msg.ErrWrongInput)
	}
}

func handleItemCommand(p *player.Player, tool item.Tool, command string) {
	r := currentRoom(*p)
	switch command {
	case "줍", "주워":
		player.AddToolToInventory(p, tool)
		room.RemoveTool(r, tool)
	default:
		fmt.Println(msg.ErrWrongInput)
	}
}

func handleThreeWordsCommand(p *player.Player, tokens []string) {
	command := tokens[2]
	switch command {
	case "사용":
		tool, ok := player.FindInventory(*p, tokens[0])
		if !ok {
			fmt.Println(msg.ErrNotHave)
			return
		}
		r := currentRoom(*p)
		if r.Door.Name != tokens[1] {
			fmt.Println(msg.ErrNotFound)
			return
		}
		useTool(p, tool, r.Door)
	default:
		fmt.Println(msg.ErrWrongInput)
	}
}

func currentRoom(p player.Player) *room.Room {
	x, y := player.CurrentPosition(p)
	return rooms[x][y]
}

func useTool(p *player.Player, tool item.Tool, d *door.Door) {
	switch tool.Type {
	case item.Hammer:
		if d.Type == door.Glass {
			door.CrashDoor(d)
			nextDoor := findNextDoor(p, *d)
			door.CrashDoor(nextDoor)
			fmt.Printf("%s로 %s을 부쉈습니다.\n", tool.Name, d.Name)
			player.RemoveToolFromInventory(p, tool)
			return
		}
	case item.Key:
		if d.Type == door.Locked {
			door.UnlockDoor(d)
			nextDoor := findNextDoor(p, *d)
			door.UnlockDoor(nextDoor)
			fmt.Printf("%s로 %s을 열었습니다.\n", tool.Name, d.Name)
			player.RemoveToolFromInventory(p, tool)
			return
		}
	}
	fmt.Println(msg.ErrCannot)
}

func findNextDoor(p *player.Player, door door.Door) *door.Door {
	x, y := player.CurrentPosition(*p)
	fX, fY := direct.GetFutureXY(door.Direction, x, y)
	nextDoor := rooms[fX][fY].Door
	if nextDoor != nil && nextDoor.Name == door.Name {
		return nextDoor
	}
	log.Panicln("cannot find next door")
	return nil
}
