package UI

import (
	"TextAdventureGame/GoTextGame/Model"
	"bufio"
	"fmt"
	"os"
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

func DisplayMenu() string{
	fmt.Println("1. Play")
	fmt.Println("2. Quit")
	reader := bufio.NewScanner(os.Stdin)
	return reader.Text()
}