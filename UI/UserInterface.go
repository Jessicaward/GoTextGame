package UI

import (
	"TextAdventureGame/GoTextGame/Model"
	"bufio"
	"fmt"
	"log"
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
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		input += scanner.Text()
	}

	if(scanner.Err() != nil){
		//not sure if I should be logging a fatal error here?
		//idk how to otherwise handle it
		//todo: look into error handling in go
		log.Fatal(scanner.Err())
	}

	return input
}