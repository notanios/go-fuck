package main

import (
	"strings"

	"github.com/golang-collections/collections/stack"
)

func sanityze(s string) string {
	legalSet := "<>,.+-[]"

	filter := func(r rune) rune {
		if strings.Contains(legalSet, string(r)) {
			return r
		}

		return -1
	}

	return strings.Map(filter, s)
}

func validate(s string) bool {
	vStack := stack.New()

	for _, c := range s {
		char := string(c)
		if char == "[" {
			vStack.Push(c)
		} else if char == "]" {
			if vStack.Peek() != nil && string(vStack.Peek().(rune)) == "[" {
				vStack.Pop()
			} else {
				return false
			}
		}
	}

	if vStack.Len() > 0 {
		return false
	}

	return true
}

func validateSimple(s string) bool {
	counter := 0

	for _, c := range s {
		char := string(c)
		if char == "[" {
			counter++
		} else if char == "]" {
			counter--
		}
	}

	if counter > 0 {
		return false
	}

	return true
}
