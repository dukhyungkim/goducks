package util

import (
	"bufio"
	"os"
	"strings"
)

func GetUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
