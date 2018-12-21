package main

// Command denotes type of node
type Command int

const (
	incrementCommand Command = iota
	decrementCommand
	goNextCommand
	goPrevCommand
	readCommand
	writeCommand
	loopCommand
)

// Node is a semantical entity of the code instruction or block "[]"
type Node struct {
	value    Command
	subnodes []Node
}

// Parse will parse BrainFuck string and transform it into realization specific AST
func parse(commands string) []Node {
	nodes := []Node{}

	for index := 0; index < len(commands); index++ {
		char := commands[index]
		switch char {
		case '.':
			nodes = append(nodes, Node{value: writeCommand, subnodes: nil})
		case ',':
			nodes = append(nodes, Node{value: readCommand, subnodes: nil})
		case '-':
			nodes = append(nodes, Node{value: decrementCommand, subnodes: nil})
		case '+':
			nodes = append(nodes, Node{value: incrementCommand, subnodes: nil})
		case '<':
			nodes = append(nodes, Node{value: goPrevCommand, subnodes: nil})
		case '>':
			nodes = append(nodes, Node{value: goNextCommand, subnodes: nil})
		case '[':
			{
				headIndex := index
				index++
				for counter := 1; counter != 0; index++ {
					char := commands[index]
					switch char {
					case '[':
						counter++
					case ']':
						counter--
					}
				}
				blockSlice := commands[headIndex+1 : index]
				nodes = append(nodes, Node{value: loopCommand, subnodes: parse(blockSlice)})
				index--
			}
		}
	}

	return nodes
}
