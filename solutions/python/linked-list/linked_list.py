"""Doubly linked list."""

from __future__ import annotations
import collections.abc


class Node[T]:
    """List node."""
    def __init__(self, value: T):
        """Create a new node."""
        self.value = value
        self.prev: Node | None = None
        self.next: Node | None = None

    def add_prev(self, value: T) -> Node:
        """Add a node prior to self."""
        new_node = Node(value)
        assert self.prev is not None
        self.link(self.prev, new_node, self)
        return new_node

    def add_next(self, value: T) -> Node:
        """Add a node after self."""
        new_node = Node(value)
        assert self.next is not None
        self.link(self, new_node, self.next)
        return new_node

    @staticmethod
    def link(node_a: Node, node_b: Node, node_c: Node) -> None:
        """Set up node linkage between three nodes in a sequence."""
        node_a.next, node_b.next = node_b, node_c
        node_b.prev, node_c.prev = node_a, node_b

    def remove(self) -> T:
        """Remove a node from the list and return its value."""
        assert self.prev is not None
        assert self.next is not None
        self.next.prev = self.prev
        self.prev.next = self.next
        return self.value


class LinkedList[T]:
    """A linked list."""

    def __init__(self):
        """Set up a linked list with an empty head and tail node."""
        # Create the head and last nodes.
        self.head = Node(None)
        self.last = Node(None)
        # Link the end nodes.
        self.head.next, self.last.prev = self.last, self.head
        self.length = 0

    def push(self, value: T) -> None:
        """Push a node to the end of the list."""
        self.last.add_prev(value)
        self.length += 1

    def unshift(self, value: T) -> None:
        """Insert a node at the start of the list."""
        self.head.add_next(value)
        self.length += 1

    def delete(self, value) -> None:
        """Remove the first node with a matching value."""
        cur = self.head.next
        while cur != self.last and cur.value != value:
            cur = cur.next
        if cur.value != value:
            raise ValueError("Value not found")
        cur.remove()
        self.length -= 1

    def remove(self, node: Node) -> T:
        """Remove a node from the list and return its value."""
        if self.length <= 0:
            raise IndexError("List is empty")
        val = node.remove()
        self.length -= 1
        return val

    def pop(self) -> T:
        """Pop a node from the end of the list."""
        return self.remove(self.last.prev)

    def shift(self) -> T:
        """Remove a node from the start of the list."""
        return self.remove(self.head.next)

    def __iter__(self) -> collections.abc.Iterator[T]:
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
