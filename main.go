package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	game := NewGame()
	game.Initialize()
	
	fmt.Println("Welcome to the Spaceship Text Adventure!")
	fmt.Println("You are aboard a derelict spaceship. Find the flag to escape.")
	fmt.Println("Use Linux commands to navigate and interact with your environment.")
	fmt.Println("Type 'help' for available commands.")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Printf("[%s]$ ", game.CurrentRoom.Name)
		if !scanner.Scan() {
			break
		}
		
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}
		
		if game.ProcessCommand(input) {
			break // Game ended
		}
	}
}