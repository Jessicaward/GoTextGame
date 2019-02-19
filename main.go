package main

import (
	"TextAdventureGame/GoTextGame/Node"
	"TextAdventureGame/GoTextGame/UI"
	"os"
)

//Main game loop.
func main() {
	userInput := UI.DisplayMenu()

	switch userInput{
		case "2":
		case "Quit":
		case "quit":
			Quit()
	default:
			Play()
	}
}

func Play(){
	var node = Node.GetNode("Start", Node.LoadFile("GoTextGame/Content/FirstGame.jess"))
	UI.DisplayNodeInformation(node)
}

func Quit(){
	os.Exit(3)
}