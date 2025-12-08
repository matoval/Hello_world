package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Item struct {
	Name        string
	Description string
	Type        string // "key", "chest", "flag", "normal"
	KeyID       int    // For matching keys to chests
	Content     string // For chests/notes
}

type Room struct {
	Name        string
	Description string
	Connections map[string]*Room
	Items       map[string]*Item
	Visited     bool
}

type Game struct {
	CurrentRoom *Room
	Inventory   map[string]*Item
	Rooms       map[string]*Room
}

func NewGame() *Game {
	rand.Seed(time.Now().UnixNano())
	
	return &Game{
		Inventory: make(map[string]*Item),
		Rooms:     make(map[string]*Room),
	}
}

func (g *Game) Initialize() {
	g.createRooms()
	g.createItems()
	g.placeItems()
	g.CurrentRoom = g.Rooms["Bridge"]
}

func (g *Game) createRooms() {
	roomNames := []string{
		"Bridge", "Engine Room", "Cargo Bay", "Medical Bay", 
		"Armory", "Crew Quarters", "Storage", "Observatory", 
		"Communications", "Airlock",
	}
	
	// Create rooms
	for _, name := range roomNames {
		g.Rooms[name] = &Room{
			Name:        name,
			Description: g.getRoomDescription(name),
			Connections: make(map[string]*Room),
			Items:       make(map[string]*Item),
			Visited:     false,
		}
	}
	
	// Create connections between rooms in a random pattern
	connections := []struct {
		from, to string
	}{
		{"Bridge", "Communications"},
		{"Bridge", "Observatory"},
		{"Bridge", "Crew Quarters"},
		{"Engine Room", "Cargo Bay"},
		{"Engine Room", "Storage"},
		{"Cargo Bay", "Medical Bay"},
		{"Cargo Bay", "Armory"},
		{"Medical Bay", "Crew Quarters"},
		{"Armory", "Storage"},
		{"Armory", "Airlock"},
		{"Crew Quarters", "Communications"},
		{"Storage", "Observatory"},
		{"Observatory", "Communications"},
		{"Communications", "Airlock"},
		{"Airlock", "Cargo Bay"},
	}
	
	// Add bidirectional connections
	for _, conn := range connections {
		g.Rooms[conn.from].Connections[conn.to] = g.Rooms[conn.to]
		g.Rooms[conn.to].Connections[conn.from] = g.Rooms[conn.from]
	}
}

func (g *Game) getRoomDescription(name string) string {
	descriptions := map[string]string{
		"Bridge":          "The command center of the ship. Control panels flicker with emergency lighting.",
		"Engine Room":     "Massive reactors hum with residual energy. Pipes and conduits line the walls.",
		"Cargo Bay":       "Crates and containers are stacked haphazardly. The air smells of metal and oil.",
		"Medical Bay":     "Sterile environment with medical equipment. Emergency lights create eerie shadows.",
		"Armory":          "Weapon racks line the walls. Ammo containers are scattered about.",
		"Crew Quarters":   "Personal belongings float in zero gravity. Some signs of a struggle are evident.",
		"Storage":         "Various supplies and equipment are stored here. It's dimly lit and cluttered.",
		"Observatory":     "Large windows show the vastness of space. Star charts cover the walls.",
		"Communications":  "Radio equipment and long-range transmitters. Control panels show incoming signals.",
		"Airlock":         "The gateway to the void of space. Heavy doors with security protocols.",
	}
	
	return descriptions[name]
}

func (g *Game) createItems() {
	// Create keys (5 keys total)
	keys := []struct {
		name, desc string
		id         int
	}{
		{"Alpha Key", "A small key with an alpha symbol etched on it.", 1},
		{"Beta Key", "A metallic key with a beta marking.", 2},
		{"Gamma Key", "A crystalline key pulsing with energy. Marked gamma.", 3},
		{"Delta Key", "A heavy key with delta inscribed on its surface.", 4},
		{"Epsilon Key", "A tiny key with epsilon written in alien script.", 5},
	}
	
	// Create chests that require specific keys
	chests := []struct {
		name, desc string
		keyID      int
		content    string
	}{
		{"Armory Chest", "A heavily reinforced chest requiring a specific key.", 1, "Military-grade plasma rifle"},
		{"Storage Chest", "A sealed container with an electronic lock.", 2, "Emergency beacon device"},
		{"Medical Chest", "A bio-secured container for medical supplies.", 3, "Antidote for alien pathogen"},
		{"Observatory Chest", "A crystalline container that resonates softly.", 4, "Ancient star map"},
		{"Engineering Chest", "A heavy chest with multiple locking mechanisms.", 5, "Quantum fusion core"},
	}
	
	// Create additional items
	notes := []struct {
		name, desc, content string
	}{
		{"Engine Log", "A partially damaged engineering log", "Engine failure caused by quantum instability. Key system malfunctioning."},
		{"Captain's Note", "Personal note from the captain", "The flag is hidden in the room most connected to our journey."},
		{"Scientist's Journal", "Research notes from the science team", "Pattern analysis shows the observatory connects to everything."},
	}
	
	// Add items to rooms
	// Keys in various rooms
	for i, keyData := range keys {
		roomName := []string{"Engine Room", "Cargo Bay", "Medical Bay", "Armory", "Storage"}[i]
		g.Rooms[roomName].Items[keyData.name] = &Item{
			Name:        keyData.name,
			Description: keyData.desc,
			Type:        "key",
			KeyID:       keyData.id,
		}
	}
	
	// Chests in various rooms
	for i, chestData := range chests {
		roomName := []string{"Armory", "Storage", "Medical Bay", "Observatory", "Engine Room"}[i]
		g.Rooms[roomName].Items[chestData.name] = &Item{
			Name:        chestData.name,
			Description: chestData.desc,
			Type:        "chest",
			KeyID:       chestData.keyID,
			Content:     chestData.content,
		}
	}
	
	// Notes in various rooms
	for i, noteData := range notes {
		roomName := []string{"Engine Room", "Bridge", "Observatory"}[i]
		g.Rooms[roomName].Items[noteData.name] = &Item{
			Name:        noteData.name,
			Description: noteData.desc,
			Type:        "normal",
			Content:     noteData.content,
		}
	}
	
	// Place the flag in a strategically chosen room (Communications)
	g.Rooms["Communications"].Items["Flag"] = &Item{
		Name:        "Flag",
		Description: "The escape flag! This marks your successful escape from the derelict ship.",
		Type:        "flag",
	}
}

func (g *Game) placeItems() {
	// Items were placed during creation, but this allows for dynamic placement logic
	// This function can be expanded if needed
}