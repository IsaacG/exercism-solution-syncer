"""Solve a cage for Killer Sudoku."""
import itertools


def combinations(target: int, size: int, exclude: list[int]) -> list[list[int]]:
    """Return possible combinations for a cage."""
    exclude_set = set(exclude)
    return [
        list(combo)
        for combo in itertools.combinations(range(1, 10), size)
        if sum(combo) == target and exclude_set.isdisjoint(combo)
    ]
