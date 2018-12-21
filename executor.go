package main

import "fmt"

func execute(nodes []Node) {
	for _, node := range nodes {
		// say(node.value)
		switch node.value {
		case incrementCommand:
			tTape.incrementValue()
		case decrementCommand:
			tTape.decrementValue()
		case goNextCommand:
			tTape.goToNextCell()
		case goPrevCommand:
			tTape.goToPreviousCell()
		case loopCommand:
			for tTape.active.Value.(byte) != byte(0) {
				execute(node.subnodes)
			}
		case writeCommand:
			tTape.printActiveValue()
		case readCommand:
		}
	}
}

func say(c Command) {
	commands := []string{"increment", "decrement", "next", "prev", "read", "write", "loop"}
	fmt.Printf("%s ", commands[c])
}
