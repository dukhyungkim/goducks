package main

import "fmt"

func main() {
	InitGame()

	for {
		PrintCurrentStatus()
		text := GetUserInput()
		HandleUserInput(text)
		fmt.Println()
		if IsGoal() {
			fmt.Println("축하합니다.")
			break
		}
	}
}
