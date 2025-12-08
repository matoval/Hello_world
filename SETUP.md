# Setup Instructions

## Prerequisites

- Go 1.19 or later installed on your system

## Running the Game

Option 1: Run directly with Go
```
go run .
```

Option 2: Build and run
```
go build -o spaceship-game .
./spaceship-game
```

On Windows:
```
go build -o spaceship-game.exe .
spaceship-game.exe
```

## Playing the Game

When you start the game, you'll be in the Bridge of a derelict spaceship. Use Linux-style commands to explore:

1. Type `ls` to see items in your current room
2. Use `cd <room>` to move to another room
3. Use `cat <item>` to examine items
4. Use `cp <key>` to add keys to your inventory
5. Use `unrar <chest> <key>` to open chests
6. Find the flag to win!

Type `help` anytime to see available commands.