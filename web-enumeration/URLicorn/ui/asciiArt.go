package ui

import (
	"bufio"
	"fmt"
	"os"
)

func PrintUnicornAscii() {
	file, err := os.Open("ui/ascii_art.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

}
