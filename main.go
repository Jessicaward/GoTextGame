package main

import (
	"TextAdventureGame/GoTextGame/Node"
	"TextAdventureGame/GoTextGame/UI"
)

func main() {
	var node = Node.GetNode("Start")

	UI.DisplayNodeInformation(node)
}