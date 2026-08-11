"""Simple cipher."""

import itertools
import random
import string

ORD_A = ord("a")


def add(char_a: str, char_b: str, reverse: bool) -> str:
    """Add two chars together. Or subtract, if reverse."""
    val = ord(char_a) - ORD_A
    dist = ord(char_b) - ORD_A
    if reverse:
        dist = -dist
    res = (val + dist) % 26
    return chr(res + ORD_A)


class Cipher:
    """Simple cipher."""

    def __init__(self, key: str | None = None):
        self.key = key or "".join(random.choices(string.ascii_lowercase, k=100))

    def encode(self, text: str, reverse: bool = False) -> str:
        """Return an encoded string using the key."""
        return "".join(add(char, key, reverse) for char, key in zip(text, itertools.cycle(self.key)))

    def decode(self, text: str) -> str:
        """Return a decoded string using the key."""
        return self.encode(text, reverse=True)
