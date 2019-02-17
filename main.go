package main

import (
	"TextAdventureGame/GoTextGame/Node"
	"TextAdventureGame/GoTextGame/UI"
)

//Main game loop.
func main() {
	var node = Node.GetNode("Start", Node.LoadFile("GoTextGame/Content/FirstGame.jess"))
	UI.DisplayNodeInformation(node)
}

//todo: write menu function