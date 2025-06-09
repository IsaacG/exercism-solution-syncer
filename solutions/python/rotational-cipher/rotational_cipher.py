"""Rotational cipher."""

import string

ORD_LOWER = ord(string.ascii_lowercase[0])
ORD_UPPER = ord(string.ascii_uppercase[0])
LEN_ALPHA = len(string.ascii_uppercase)


def rotate(text: str, key: int) -> str:
    """Rotate a string."""
    return ''.join(_add(t, key) for t in text)


def _add(char: str, key: int) -> str:
    """Rotate one character."""
    if char in string.ascii_uppercase:
        return _normalize(ord(char) + key, ORD_UPPER)
    elif char in string.ascii_lowercase:
        return _normalize(ord(char) + key, ORD_LOWER)
    else:
        # Non-alpha characters are unchanged.
        return char


def _normalize(char: str, start_ord: int) -> str:
    """Normalize a char to fall within an alphabet."""
    return chr((char - start_ord) % LEN_ALPHA + start_ord)
