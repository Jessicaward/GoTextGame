package main

import (
	"TextAdventureGame/GoTextGame/Node"
	"TextAdventureGame/GoTextGame/UI"
	"os"
)

//Main game loop.
func main() {
	userInput := UI.DisplayMenu()

	if userInput == "1" || userInput == "Play" || userInput == "play" {
		Play()
	} else {
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