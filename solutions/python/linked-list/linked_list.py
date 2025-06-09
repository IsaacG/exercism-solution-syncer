"""Doubly linked list."""

from __future__ import annotations


class Node:
    """List node."""
    def __init__(self, value, succeeding=None, previous=None):
        """Create a new node, inserting into a list as needed."""

        # If given one or the other, compute the missing one.
        if succeeding and not previous:
            previous = succeeding.prev
        if previous and not succeeding:
            succeeding = previous.next

        self.value = value
        self.next = succeeding
        self.prev = previous

        # Adjust the linking on the prev and next nodes.
        if self.next:
            self.next.prev = self
        if self.prev:
            self.prev.next = self

    def remove(self):
        """Remove a node from the list and return its value."""
        # Relink prev and next to each other.
        if self.next:
            self.next.prev = self.prev
        if self.prev:
            self.prev.next = self.next
        return self.value


class LinkedList:
    """A linked list."""

    def __init__(self):
        """Set up a linked list with an empty head and tail node."""
        self.head = Node(None)
        self.last = Node(None, previous=self.head)
        self.length = 0

    def push(self, value):
        """Push a node to the end of the list."""
        Node(value, succeeding=self.last)
        self.length += 1

    def unshift(self, value):
        """Insert a node at the start of the list."""
        Node(value, previous=self.head)
        self.length += 1

    def pop(self):
        """Pop a node from the end of the list."""
        assert self.length, "List is empty."
        self.length -= 1
        return self.last.prev.remove()

    def shift(self):
        """Remove a node from the start of the list."""
        assert self.length, "List is empty."
        self.length -= 1
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
