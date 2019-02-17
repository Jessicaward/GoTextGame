package UI

import (
	"TextAdventureGame/GoTextGame/Model"
	"fmt"
)

//Summary of node is displayed.
//This also calls function to display all options.
func DisplayNodeInformation(node Model.Node){
	fmt.Println(node.Text)
	DisplayOptions(node.Options)
}

//Options are formatted and displayed to user.
func DisplayOptions(options []Model.Option){
	for i := 0; i < len(options); i++{
		fmt.Println("> " + options[i].Value)
	}
}