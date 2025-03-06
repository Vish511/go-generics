# Singly Linked List in Go using Generics

This is a Go implementation of a **Singly Linked List** data structure using **generics**. The code demonstrates the flexibility of generics in Go, allowing the linked list to store values of any data type. You can perform common linked list operations like **Push**, **Pop**, **Unshift**, **Shift**, **Insert**, **Remove**, **Reverse**, and more.

## Features
- **Generic Linked List**: Works with any data type using Go generics (`T any`).
- **Basic Operations**: Includes methods for adding, removing, and manipulating nodes.
- **Traversal**: Ability to traverse and print the list.
- **Reverse**: Reverse the order of the nodes in the list.

## Table of Contents
- [Installation](#installation)
- [Usage](#usage)
- [Methods](#methods)
- [Example](#example)

## Installation

To use the singly linked list in your Go project, clone this repository and include it in your Go module.

```bash
git clone https://github.com/Vish511/go-generics.git
```

In your Go project, import the package:

```go
import "path/to/singlylinkedlist"
```

## Usage

Below is an example of how to create a new linked list and perform different operations:

```go
package main

import (
	"fmt"
	"singlylinkedlist"
)

func main() {
	// Example: Singly Linked List with int type
	fmt.Println("Example 1: Singly Linked List with int")
	intList := singlylinkedlist.NewSinglyLinkedList[int]()
	intList.Push(10)
	intList.Push(20)
	intList.Push(30)
	intList.Traverse() // Output: 10 -> 20 -> 30 -> nil

	// Perform various operations
	intList.Pop() // Remove last node
	intList.Traverse() // Output: 10 -> 20 -> nil

	intList.Unshift(5) // Add to the beginning
	intList.Traverse() // Output: 5 -> 10 -> 20 -> nil

	intList.Shift() // Remove the first node
	intList.Traverse() // Output: 10 -> 20 -> nil

	// Insert at specific position
	intList.Insert(1, 15)
	intList.Traverse() // Output: 10 -> 15 -> 20 -> nil

	// Reverse the list
	intList.Reverse()
	intList.Traverse() // Output: 20 -> 15 -> 10 -> nil

	// Example: Singly Linked List with string type
	fmt.Println("\nExample 2: Singly Linked List with string")
	stringList := singlylinkedlist.NewSinglyLinkedList[string]()
	stringList.Push("A")
	stringList.Push("B")
	stringList.Push("C")
	stringList.Traverse() // Output: A -> B -> C -> nil

	// Perform operations on the string list
	stringList.Pop() // Remove last node
	stringList.Traverse() // Output: A -> B -> nil

	stringList.Unshift("Z") // Add to the beginning
	stringList.Traverse() // Output: Z -> A -> B -> nil

	stringList.Shift() // Remove the first node
	stringList.Traverse() // Output: A -> B -> nil

	// Insert at specific position
	stringList.Insert(1, "Y")
	stringList.Traverse() // Output: A -> Y -> B -> nil

	// Reverse the list
	stringList.Reverse()
	stringList.Traverse() // Output: B -> Y -> A -> nil
}
```

## Methods

The linked list supports the following operations:

- **Push(value T)**: Adds a new node with the specified value to the end of the list.
- **Pop()**: Removes and returns the last node from the list.
- **Unshift(value T)**: Adds a new node with the specified value to the beginning of the list.
- **Shift()**: Removes and returns the first node from the list.
- **Get(idx int)**: Returns the node at the given index.
- **Set(idx int, value T)**: Updates the value of the node at the specified index.
- **Insert(idx int, value T)**: Inserts a new node at the given index.
- **Remove(idx int, value T)**: Removes and returns the node at the given index.
- **Reverse()**: Reverses the order of nodes in the list.
- **Traverse()**: Prints all nodes in the list.

## Example

Example output of the code running on an `int` and `string` linked list:

```
Example 1: Singly Linked List with int
10 -> 20 -> 30 -> nil
10 -> 20 -> nil
5 -> 10 -> 20 -> nil
10 -> 20 -> nil
10 -> 15 -> 20 -> nil
20 -> 15 -> 10 -> nil

Example 2: Singly Linked List with string
A -> B -> C -> nil
A -> B -> nil
Z -> A -> B -> nil
A -> B -> nil
A -> Y -> B -> nil
B -> Y -> A -> nil
```

---
```
