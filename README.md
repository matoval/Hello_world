# Spaceship Text Adventure

A text-based adventure game built in Go that simulates Linux command-line navigation through a derelict spaceship.

## Game Objective

You are aboard a derelict spaceship. Your goal is to find the hidden flag to escape. Explore the ship, collect keys, unlock chests, and solve the mystery to find your way out.

## How to Play

1. Navigate between rooms using Linux-style commands
2. Find keys hidden throughout the ship
3. Use keys to unlock chests and discover useful items
4. Find the hidden flag to win the game

## Commands

- `ls [room]` - List items in current room or specified room
- `cd <room>` - Move to another room
- `cat <item>` - Examine an item
- `cp <item>` - Copy an item (add to inventory)
- `unrar <chest> <key>` - Open a chest with a specific key
- `inventory` - Show your inventory
- `help` - Show available commands
- `exit` - Quit the game

## Game Features

- 10 interconnected spaceship rooms
- Key/chest puzzle mechanics
- Hidden flag to find
- Inventory system
- Linux-style commands for navigation and interaction

## How to Run

1. Make sure you have Go installed
2. Run `go run .` in the project directory
3. Start exploring!

Good luck, and watch out for what lurks in the dark corners of the ship...