package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mvstermind/cyber-toolkit/web-enumeration/URLicorn/ui"
)

var EXITMESSAGES []string

func init() {
	EXITMESSAGES = append(EXITMESSAGES, "bye", "exit")
}

func main() {
	ui.PrintUnicornAscii()

	var userInput string
	for {
		if userInput == EXITMESSAGES[0] || userInput == EXITMESSAGES[1] {
			fmt.Println("bye!!")
			break
		}

		userInput = getUserInput()
		fmt.Println(userInput)
	}
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	text, _ := reader.ReadString('\n')
	trimmed := strings.TrimSpace(text)

	return trimmed
}
