package main

import (
	"container/list"
	"fmt"
)

var tTape *Tape

// Tape is a representation of Turing Machine tape
type Tape struct {
	cells  *list.List
	active *list.Element
	len    int
}

// NewTape returns a new empty Tape instance, please use TapeInit instead
func NewTape() *Tape {
	listVar := list.New()
	byteValue := byte(0)
	listVar.PushFront(byteValue)
	tapeVar := Tape{cells: listVar, active: listVar.Front(), len: 1}

	return &tapeVar
}

// TapeInit is a global tape init method
func TapeInit() {
	tTape = NewTape()
}

func (this *Tape) incrementValue() {
	if cellValue, ok := this.active.Value.(byte); ok {
		cellValue++
		this.active.Value = cellValue
	} else {
		fmt.Println(cellValue, ok)
	}
}

func (this *Tape) decrementValue() {
	if cellValue, ok := this.active.Value.(byte); ok {
		cellValue--
		this.active.Value = cellValue
	} else {
		fmt.Println(cellValue, ok)
	}
}

// func (this *tape) read() {

// }

func (this *Tape) printActiveValue() {
	fmt.Printf("%c", this.active.Value)
}

func (this *Tape) printAll() {
	for e := this.cells.Front(); e != nil; e = e.Next() {
		fmt.Printf("<%d, %c> ", e.Value, e.Value)
	}
	fmt.Println()
}

func (this *Tape) goToNextCell() {
	nextCell := this.active.Next()

	if nextCell == nil {
		this.cells.PushBack(byte(0))
		nextCell = this.active.Next()
	}

	this.active = nextCell
}

func (this *Tape) goToPreviousCell() {
	prevCell := this.active.Prev()

	if prevCell == nil {
		this.cells.PushFront(byte(0))
		prevCell = this.active.Prev()
	}

	this.active = prevCell
}
