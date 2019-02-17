package UI

import (
	"TextAdventureGame/GoTextGame/Model"
	"fmt"
)

func DisplayNodeInformation(node Model.Node){
	DisplayOptions(node.Options)
}

func DisplayOptions(options []Model.Option){
	for i := 0; i < len(options); i++{
		fmt.Println(options[i].Value)
	}
}