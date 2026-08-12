"""Solve an alphametics puzzle."""

import itertools
import string


def solve(puzzle: str) -> dict[str, int] | None:
    """Return a puzzle solution."""
    left, right = puzzle.split(" == ")
    parts = [p.strip() for p in left.split("+")]
    words = parts + [right]
    letters = {letter for word in words for letter in word}
    # Try all permutations of digits for the letters.
    for vals in itertools.permutations(string.digits, r=len(letters)):
        mapping = dict(zip(letters, vals))
        translation = str.maketrans(mapping)
        translated = [w.translate(translation) for w in words]
        # All numbers cannot start with a 0.
        if any(w[0].startswith("0") for w in translated):
            continue
        if sum(int(w) for w in translated[:-1]) == int(translated[-1]):
            return {k: int(v) for k, v in mapping.items()}
    return None
