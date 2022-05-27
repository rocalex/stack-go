package main

import (
	"fmt"
)

const MAX int = 10

var size int = 0

type Stack struct {
	items [MAX]int
	top   int
}

func main() {
	var s Stack

	createEmptyStack(&s)

	push(&s, 1)
	push(&s, 2)
	push(&s, 3)
	push(&s, 4)

	printStack(&s)

	pop(&s)

	fmt.Print("\nAfter popping out\n")
	printStack(&s)
}

func createEmptyStack(s *Stack) {
	s.top = -1
}

func isFull(s *Stack) bool {
	if s.top == MAX-1 {
		return true
	} else {
		return false
	}
}

func isEmpty(s *Stack) bool {
	if s.top == -1 {
		return true
	} else {
		return false
	}
}

func push(s *Stack, newItem int) {
	if isFull(s) {
		fmt.Println("STACK FULL")
	} else {
		s.top++
		s.items[s.top] = newItem
	}
	size++
}

func pop(s *Stack) {
	if isEmpty(s) {
		fmt.Print("\n STACK EMPTY \n")
	} else {
		fmt.Print("Item popped= ")
		fmt.Print(s.items[s.top])
		s.top--
	}
	size--
	fmt.Print("\n")
}

func printStack(s *Stack) {
	fmt.Print("Stack: ")
	for i := 0; i < size; i++ {
		fmt.Print(s.items[i])
		fmt.Print(" ")
	}
}
