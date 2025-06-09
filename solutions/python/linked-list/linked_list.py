"""Doubly linked list."""

from __future__ import annotations


class Node:
    """List node."""
    def __init__(self, value):
        """Create a new node."""
        self.value = value
        self.prev = None
        self.next = None

    def add_prev(self, value):
        """Add a node prior to self."""
        new_node = Node(value)
        self.link(self.prev, new_node, self)
        return new_node

    def add_next(self, value):
        """Add a node after self."""
        new_node = Node(value)
        self.link(self, new_node, self.next)
        return new_node

    @staticmethod
    def link(node_a, node_b, node_c):
        """Set up node linkage between three nodes in a sequence."""
        node_a.next, node_b.next = node_b, node_c
        node_b.prev, node_c.prev = node_a, node_b

    def remove(self):
        """Remove a node from the list and return its value."""
        self.next.prev = self.prev
        self.prev.next = self.next
        return self.value


def increment(func):
    """LinkedList increment decorator to bump the length."""

    def wrapper(self, *args):
        func(self, *args)
        self.length += 1

    return wrapper


def decrement(func):
    """LinkedList decrement decorator to check and reduce the length."""

    def wrapper(self, *args):
        if self.length <= 0:
            raise IndexError("Cannot remove from an empty list.")
        self.length -= 1
        return func(self, *args)

    return wrapper


class LinkedList:
    """A linked list."""

    def __init__(self):
        """Set up a linked list with an empty head and tail node."""
        # Create the head and last nodes.
        self.head = Node(None)
        self.last = Node(None)
        # Link the end nodes.
        self.head.next, self.last.prev = self.last, self.head
        self.length = 0

    @increment
    def push(self, value):
        """Push a node to the end of the list."""
        self.last.add_prev(value)

    @increment
    def unshift(self, value):
        """Insert a node at the start of the list."""
        self.head.add_next(value)

    @decrement
    def pop(self):
        """Pop a node from the end of the list."""
        return self.last.prev.remove()

    @decrement
    def shift(self):
        """Remove a node from the start of the list."""
        return self.head.next.remove()

    def __iter__(self):
        """Iterate through the list."""
        cur = self.head.next
        while cur != self.last:
            yield cur.value
            cur = cur.next

    def __str__(self) -> str:
        """Return a string form."""
        return " -> ".join(repr(i) for i in self)

    def __len__(self) -> int:
        """Return the list length."""
        return self.length
