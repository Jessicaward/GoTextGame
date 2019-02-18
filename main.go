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
		case "1":
		case "Play":
		case "play":
			Play()
		default:
			Quit()
	}
}

func Play(){
	var node = Node.GetNode("Start", Node.LoadFile("GoTextGame/Content/FirstGame.jess"))
	UI.DisplayNodeInformation(node)
}

func Quit(){
	os.Exit(3)
}