package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/mvstermind/cyber-toolkit/web-enumeration/URLicorn/ui"
)

var ExitMessages []string
var Commands []string

func init() {
	ExitMessages = append(ExitMessages, "bye", "exit")
	Commands = append(Commands,
		"help",
		"target",
		"wordlist",
		"run",
		"dir",
		"sub",
		"save",
	)
}

type Target struct {
	addr     string
	wordlist string
	options  [2]string
	run      bool
	saveFile string
}

func NewTarget(name string) Target {
	return Target{
		addr:     name,
		wordlist: "",
		options:  [2]string{},
		run:      false,
		saveFile: "",
	}
}

func main() {
	ui.PrintUnicornAscii()

	var userInput string
	for {
		if userInput == ExitMessages[0] || userInput == ExitMessages[1] {
			fmt.Println("bye!!")
			break
		}

		userInput = getUserInput()
		err := checkUserChoice(userInput)
		if err != nil {
			fmt.Printf("Error!!\n%s", err)
		}
	}
}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	text, _ := reader.ReadString('\n')
	trimmed := strings.TrimSpace(text)

	return trimmed
}

func checkUserChoice(input string) error {

	inputSlice := strings.Split(input, " ")
	userCommand := strings.ToLower(inputSlice[0])
	v := 0

	for i := 0; i < len(Commands); i++ {
		if userCommand == Commands[i] {
			v++
		}
	}

	if v != 1 {
		return errors.New("unexpected input")
	}

	target := NewTarget(input)

	switch input {
	// "help"
	case Commands[0]:
		target.displayHelp()

	// "target"
	case Commands[1]:
		target.setTarget(inputSlice[1])

	}

	return nil
}

func (t *Target) displayHelp() {

	fmt.Println(`
Options:
    target - Specify target to scan
    set wordlist [dir/sub]- Set wordlist to use for scanning (separate for dir and sub)
	help - Display help message
	run - Execute command with given settings
    save - Save output (specify output filename)

Scann options:
    uri - Scan directories and files in website
    sub - Scan subdomains

        `)
}

func (t *Target) setTarget(victimAddr string) {
	t.addr = victimAddr
}
