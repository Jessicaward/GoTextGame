package Node

import (
	"TextAdventureGame/GoTextGame/Model"
	"bufio"
	"os"
	"strings"
)

/*The whole idea behind this parser is to be able to write a simple file format, which can be parsed by this parser.
The file will need to contain a number of things:
	-Key for the current node.
	-Text to be displayed for the user.
	-Options that the user can select
		-Option text
		-Node 'key' for the option
Example of the file format:
----
@Start [You're in a forest, what do you do?]
>Scream [Scream]
>RunAway [Run away]

@Scream [You screamed, a bear heard you, and then ate you.]
>Start [Retry?]

@RunAway [You ran north, and found a village, everyone is so nice!]
>NextNodeHere [Do something]
----

Great Stack Exchange post for the node/option idea: https://gamedev.stackexchange.com/questions/144873/optimizing-data-structure-for-my-text-adventure */

//This returns the current node, along with all the options.
func GetNode(nodeKey string, openedFile string) Model.Node{
	var node Model.Node
	nodeLine := strings.Split(openedFile, "@" + nodeKey)[1]
	optionLines := strings.Split(nodeLine, ">")

	node.Key = nodeKey
	node.Text = strings.Replace(strings.Replace(optionLines[0], "]", "", -1), "[", "", -1) //todo: this is code gore, for the love of god please remove it.
	node.Options = GetOptionsForNode(nodeKey, optionLines[1:])

	return node
}

//This returns all the options from the node, specified by node key.
func GetOptionsForNode(nodeKey string, optionLines []string) []Model.Option{
	var options []Model.Option

	for i := 0; i < len(optionLines); i++{

	}

	return options
}

//File is loaded in (by path specified as argument)
//Returns file as a string (with all new lines removed).
func LoadFile(filePath string) string{
	data := ""
	file, error := os.Open(filePath)
	CheckError(error)
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		data += scanner.Text()
	}

	return string(data[:])
}

//If error exists, exception is raised, all hell breaks loose, shit goes down.
func CheckError(e error){
	if e != nil{
		//aaaaaaaa panic
		panic(e)
	}
}