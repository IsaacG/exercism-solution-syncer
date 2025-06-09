"""Rotational cipher."""

import string

def rotate(text: str, key: int) -> str:
    """Rotate a string."""
    return ''.join(_add(t, key) for t in text)

def _add(char: str, key: int) -> str:
    """Rotate one character."""
    if char in string.ascii_uppercase:
        return chr((ord(char) + key - ord('A')) % 26 + ord('A'))
    elif char in string.ascii_lowercase:
        return chr((ord(char) + key - ord('a')) % 26 + ord('a'))
    else:
        # Non-alpha characters are unchanged.
        return char
