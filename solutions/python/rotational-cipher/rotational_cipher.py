"""Rotational cipher."""

import string

LEN_ALPHA = len(string.ascii_uppercase)
CHAR_RANGES = (
    (string.ascii_lowercase, ord(string.ascii_lowercase[0])),
    (string.ascii_uppercase, ord(string.ascii_uppercase[0])),
)


def rotate(text: str, key: int) -> str:
    """Rotate a string."""
    return ''.join(_add(t, key) for t in text)


def _add(char: str, key: int) -> str:
    """Rotate one character."""
    for chars, start in CHAR_RANGES:
        if char in chars:
            return _normalize(ord(char) + key, start)
    else:
        # Non-alpha characters are unchanged.
        return char


def _normalize(char: str, start_ord: int) -> str:
    """Normalize a char to fall within an alphabet."""
    return chr((char - start_ord) % LEN_ALPHA + start_ord)
