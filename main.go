package main

import (
	"fmt"

	"github.com/Vish511/go-generics/singlylinkedlist"
)

func main() {
	// Example 1: Using SinglyLinkedList with int type
	fmt.Println("Example 1: Singly Linked List with int")
	intList := singlylinkedlist.NewSinglyLinkedList[int]()
	intList.Push(10)
	intList.Push(20)
	intList.Push(30)
	intList.Traverse()

	// Removing from the end
	intList.Pop()
	intList.Traverse()

	// Adding at the beginning
	intList.Unshift(5)
	intList.Traverse()

	// Shift (removing from the beginning)
	intList.Shift()
	intList.Traverse()

	// Inserting at a specific position
	intList.Insert(1, 15)
	intList.Traverse()

	// Reversing the list
	intList.Reverse()
	intList.Traverse()

	// Example 2: Using SinglyLinkedList with string type
	fmt.Println("\nExample 2: Singly Linked List with string")
	stringList := singlylinkedlist.NewSinglyLinkedList[string]()
	stringList.Push("A")
	stringList.Push("B")
	stringList.Push("C")
	stringList.Traverse()

	// Removing from the end
	stringList.Pop()
	stringList.Traverse()

	// Adding at the beginning
	stringList.Unshift("Z")
	stringList.Traverse()

	// Shift (removing from the beginning)
	stringList.Shift()
	stringList.Traverse()

	// Inserting at a specific position
	stringList.Insert(1, "Y")
	stringList.Traverse()

	// Reversing the list
	stringList.Reverse()
	stringList.Traverse()
}
