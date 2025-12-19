// Package linkedlist provides linkedlist implementation
package linkedlist

import "mod_name/utils/collections/linkednode"

/*
LinkedList this types colapse a Simple Linked List.

This can be used as a Simple Linked List, Queue or Stack.

This list accept any type, but just one specific type by time.

If you wish to use it as a List of structs, don't forgot the *
to declare it as a pointer.

Using as a Queue:

	Use the method AddBack() to add a element as the last on queue.
	Use the method PopFirst() to pop the first element on queue.
	This will do the pattern FIFO (first in, first out).

Using as a Stack:

	Attention: this struct can be used as a Stack, but it is
	less economic than create your own Stack.

	Use the method AddFront() to add a element as the first on Stack.
	Use the method PopFirst() to pop the first element on Stack.
	It will do the pattern LIFO (last in, first ou)

Method list:
*/
type LinkedList[T any] struct {
	size uint
	head *linkednode.SimpleLinkedNode[T]
	tail *linkednode.SimpleLinkedNode[T]
}

// New a constructor for LinkedList
func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{
		size: 0,
		head: nil,
		tail: nil,
	}
}
