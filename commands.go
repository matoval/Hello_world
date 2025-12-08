package main

import (
	"fmt"
	"strings"
)

func (g *Game) ProcessCommand(input string) bool {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return false
	}
	
	command := strings.ToLower(parts[0])
	args := parts[1:]
	
	switch command {
	case "help":
		g.showHelp()
	case "ls":
		g.listItems(args)
	case "cd":
		g.changeRoom(args)
	case "cat":
		g.examineItem(args)
	case "cp":
		g.copyItem(args)
	case "unrar":
		g.openChest(args)
	case "inventory", "inv":
		g.showInventory()
	case "exit", "quit":
		fmt.Println("Thanks for playing!")
		return true
	default:
		fmt.Printf("Command not recognized: %s\n", command)
		fmt.Println("Type 'help' for available commands.")
	}
	
	return false
}

func (g *Game) showHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  ls [room]          - List items in current room or specified room")
	fmt.Println("  cd <room>          - Move to another room")
	fmt.Println("  cat <item>         - Examine an item")
	fmt.Println("  cp <item>          - Copy an item (add to inventory)")
	fmt.Println("  unrar <chest> <key> - Open a chest with a specific key")
	fmt.Println("  inventory          - Show your inventory")
	fmt.Println("  help               - Show this help message")
	fmt.Println("  exit               - Quit the game")
}

func (g *Game) listItems(args []string) {
	if len(args) > 0 {
		// List items in specified room
		roomName := strings.Join(args, " ")
		if room, exists := g.Rooms[roomName]; exists {
			if !room.Visited && room != g.CurrentRoom {
				fmt.Println("You haven't been to that room yet.")
				return
			}
			g.printRoomItems(room)
			return
		}
		// Try with first arg only
		if room, exists := g.Rooms[args[0]]; exists {
			if !room.Visited && room != g.CurrentRoom {
				fmt.Println("You haven't been to that room yet.")
				return
			}
			g.printRoomItems(room)
			return
		}
		fmt.Printf("Room not found: %s\n", roomName)
		return
	}
	
	// List items in current room
	g.printRoomItems(g.CurrentRoom)
}

func (g *Game) printRoomItems(room *Room) {
	fmt.Printf("Items in %s:\n", room.Name)
	if len(room.Items) == 0 {
		fmt.Println("  (nothing of interest)")
		return
	}
	
	for itemName := range room.Items {
		fmt.Printf("  %s\n", itemName)
	}
	
	// Show connections
	fmt.Printf("\nConnected rooms:\n")
	if len(room.Connections) == 0 {
		fmt.Println("  (no connections)")
		return
	}
	
	for roomName := range room.Connections {
		fmt.Printf("  %s\n", roomName)
	}
}

func (g *Game) changeRoom(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: cd <room>")
		return
	}
	
	roomName := strings.Join(args, " ")
	
	// Try exact match first
	if room, exists := g.CurrentRoom.Connections[roomName]; exists {
		g.CurrentRoom = room
		room.Visited = true
		fmt.Printf("Moved to %s\n", room.Name)
		fmt.Printf("Description: %s\n", room.Description)
		return
	}
	
	// Try with first argument only
	if room, exists := g.CurrentRoom.Connections[args[0]]; exists {
		g.CurrentRoom = room
		room.Visited = true
		fmt.Printf("Moved to %s\n", room.Name)
		fmt.Printf("Description: %s\n", room.Description)
		return
	}
	
	fmt.Printf("Cannot move to %s from here.\n", roomName)
	fmt.Print("Available rooms: ")
	roomNames := make([]string, 0, len(g.CurrentRoom.Connections))
	for name := range g.CurrentRoom.Connections {
		roomNames = append(roomNames, name)
	}
	fmt.Println(strings.Join(roomNames, ", "))
}

func (g *Game) examineItem(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: cat <item>")
		return
	}
	
	itemName := strings.Join(args, " ")
	
	// Check current room
	if item, exists := g.CurrentRoom.Items[itemName]; exists {
		fmt.Printf("%s: %s\n", item.Name, item.Description)
		if item.Content != "" {
			fmt.Printf("Content: %s\n", item.Content)
		}
		return
	}
	
	// Check inventory
	if item, exists := g.Inventory[itemName]; exists {
		fmt.Printf("%s: %s\n", item.Name, item.Description)
		if item.Content != "" {
			fmt.Printf("Content: %s\n", item.Content)
		}
		return
	}
	
	fmt.Printf("Item not found: %s\n", itemName)
}

func (g *Game) copyItem(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: cp <item>")
		return
	}
	
	itemName := strings.Join(args, " ")
	
	// Check if item exists in current room
	if item, exists := g.CurrentRoom.Items[itemName]; exists {
		// Only copy keys to inventory (to maintain game logic)
		if item.Type == "key" || item.Type == "normal" {
			g.Inventory[itemName] = item
			fmt.Printf("Copied %s to inventory.\n", itemName)
			return
		} else if item.Type == "flag" {
			// Special case for flag - game completion
			g.Inventory[itemName] = item
			fmt.Println("Congratulations! You found the flag and successfully escaped the derelict spaceship!")
			fmt.Println("You win!")
			return
		} else {
			fmt.Printf("You cannot copy %s. Use 'unrar' to open it.\n", itemName)
			return
		}
	}
	
	fmt.Printf("Item not found in current room: %s\n", itemName)
}

func (g *Game) openChest(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage: unrar <chest> <key>")
		return
	}
	
	chestName := args[0]
	keyName := strings.Join(args[1:], " ")
	
	// Check if chest exists in current room
	chest, chestExists := g.CurrentRoom.Items[chestName]
	if !chestExists {
		fmt.Printf("Chest not found in current room: %s\n", chestName)
		return
	}
	
	if chest.Type != "chest" {
		fmt.Printf("%s is not a chest.\n", chestName)
		return
	}
	
	// Check if key is in inventory
	key, keyExists := g.Inventory[keyName]
	if !keyExists {
		fmt.Printf("Key not found in inventory: %s\n", keyName)
		return
	}
	
	if key.Type != "key" {
		fmt.Printf("%s is not a key.\n", keyName)
		return
	}
	
	// Check if key matches chest
	if chest.KeyID != key.KeyID {
		fmt.Printf("The %s doesn't fit the %s.\n", keyName, chestName)
		return
	}
	
	// Successfully opened chest
	fmt.Printf("You use the %s to open the %s.\n", keyName, chestName)
	fmt.Printf("Inside you find: %s\n", chest.Content)
	
	// Add content to room as a new item
	contentItem := &Item{
		Name:        fmt.Sprintf("%s contents", chestName),
		Description: fmt.Sprintf("Contents retrieved from the %s", chestName),
		Type:        "normal",
		Content:     chest.Content,
	}
	g.CurrentRoom.Items[contentItem.Name] = contentItem
	
	// Remove chest from room
	delete(g.CurrentRoom.Items, chestName)
}

func (g *Game) showInventory() {
	fmt.Println("Inventory:")
	if len(g.Inventory) == 0 {
		fmt.Println("  (empty)")
		return
	}
	
	for itemName := range g.Inventory {
		fmt.Printf("  %s\n", itemName)
	}
}