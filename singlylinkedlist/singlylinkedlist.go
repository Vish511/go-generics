package singlylinkedlist

import "fmt"

// SinglyLinkedList is a generic singly linked list data structure
type SinglyLinkedList[T any] struct {
	Head *Node[T] // The first node of the list
	Tail *Node[T] // The last node of the list
	Len  int      // The length of the list
}

// Node represents a single element in the linked list
type Node[T any] struct {
	Value T        // The value of the node
	Next  *Node[T] // The pointer to the next node
}

// NewSinglyLinkedList initializes and returns a new empty singly linked list
func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

// NewNode creates and returns a new node with the given value
func NewNode[T any](value T) *Node[T] {
	return &Node[T]{
		Value: value,
	}
}

// Push adds a new node with the given value to the end of the list
func (sll *SinglyLinkedList[T]) Push(value T) {
	var newNode = NewNode(value)
	if sll.Head == nil {
		// If the list is empty, set both head and tail to the new node
		sll.Head = newNode
		sll.Tail = newNode
		sll.Len++
		return
	}
	// Append the new node to the end of the list and update the tail
	sll.Tail.Next = newNode
	sll.Tail = newNode
	sll.Len++
}

// Pop removes and returns the last node from the list
func (sll *SinglyLinkedList[T]) Pop() *Node[T] {
	if sll.Len == 0 {
		// If the list is empty, return nil
		return nil
	}
	if sll.Len == 1 {
		// If there's only one node, remove it and reset head and tail
		var poppedNode = sll.Head
		sll.Head = nil
		sll.Tail = nil
		sll.Len = 0
		return poppedNode
	}
	// Otherwise, remove the tail node and return it
	var oldTail = sll.Tail
	var nodeBeforeTail = sll.getNodeBeforeTail()
	nodeBeforeTail.Next = nil
	return oldTail
}

// Unshift adds a new node with the given value to the beginning of the list
func (sll *SinglyLinkedList[T]) Unshift(value T) {
	if sll.Len == 0 {
		// If the list is empty, simply push the value
		sll.Push(value)
		return
	}
	// Create a new node and make it the new head
	var newNode = NewNode(value)
	var currHead = sll.Head
	newNode.Next = currHead
	sll.Head = newNode
	sll.Len++
}

// Shift removes and returns the first node from the list
func (sll *SinglyLinkedList[T]) Shift() *Node[T] {
	if sll.Len < 2 {
		// If there's only one or no nodes, pop the element
		return sll.Pop()
	}
	// Otherwise, remove the head node and return it
	var oldHead = sll.Head
	var newHead = oldHead.Next
	sll.Head = newHead
	sll.Len--
	return oldHead
}

// Get returns the node at the given index
func (sll *SinglyLinkedList[T]) Get(idx int) *Node[T] {
	if idx < 0 || idx >= sll.Len {
		// Return nil if the index is out of bounds
		return nil
	}
	var currIdx = 0
	var curr = sll.Head
	for curr != nil {
		if currIdx == idx {
			// If the current index matches the given index, return the node
			return curr
		}
		curr = curr.Next
		currIdx++
	}
	return nil
}

// Set updates the value of the node at the given index
func (sll *SinglyLinkedList[T]) Set(idx int, value T) bool {
	var foundNode = sll.Get(idx)
	if foundNode != nil {
		// If the node exists, update its value
		foundNode.Value = value
		return true
	}
	return false
}

// Insert adds a new node with the given value at the specified index
func (sll *SinglyLinkedList[T]) Insert(idx int, value T) bool {
	if idx < 0 || idx > sll.Len {
		// Return false if the index is out of bounds
		return false
	}
	if idx == 0 {
		// If inserting at the beginning, use Unshift
		sll.Unshift(value)
		return true
	}
	if idx == sll.Len {
		// If inserting at the end, use Push
		sll.Push(value)
		return true
	}
	// Insert the new node at the given index
	var newNode = NewNode(value)
	var prev = sll.Get(idx - 1)
	var prevNext = prev.Next
	prev.Next = newNode
	newNode.Next = prevNext
	sll.Len++
	return true
}

// Remove removes and returns the node at the given index
func (sll *SinglyLinkedList[T]) Remove(idx int, value T) *Node[T] {
	if idx < 0 || idx >= sll.Len {
		// Return nil if the index is out of bounds
		return nil
	}
	if idx == 0 {
		// If removing from the beginning, use Shift
		return sll.Shift()
	}
	if idx == sll.Len-1 {
		// If removing from the end, use Pop
		return sll.Pop()
	}
	// Otherwise, remove the node from the middle
	var prev = sll.Get(idx - 1)
	var removedNode = prev.Next
	prev.Next = removedNode.Next
	sll.Len--
	return removedNode
}

// Reverse reverses the order of nodes in the list
func (sll *SinglyLinkedList[T]) Reverse() {
	var curr = sll.Head
	var prev *Node[T]
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	// Swap head and tail after reversal
	sll.Head, sll.Tail = sll.Tail, sll.Head
}

// getNodeBeforeTail returns the node before the tail node
func (sll *SinglyLinkedList[T]) getNodeBeforeTail() *Node[T] {
	if sll.Len <= 1 {
		// If there's one or no nodes, there's no node before the tail
		return nil
	}
	if sll.Len == 2 {
		// If there are exactly two nodes, the head is the node before the tail
		return sll.Head
	}
	// Traverse to find the node before the tail
	var curr = sll.Head
	var nodeBeforeTail *Node[T]
	for curr.Next != nil {
		nodeBeforeTail = curr
		curr = curr.Next
	}
	return nodeBeforeTail
}

// Traverse prints the elements of the list
func (sll *SinglyLinkedList[T]) Traverse() {
	var curr = sll.Head
	for curr != nil {
		// Print the current node's value
		fmt.Printf("%v", curr.Value)
		if curr.Next != nil {
			// Print " -> " if there is a next node
			fmt.Print(" -> ")
		}
		curr = curr.Next
	}
	// Print the end of the list
	fmt.Println(" -> nil")
}
