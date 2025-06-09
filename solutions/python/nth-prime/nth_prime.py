"""Compute primes."""

from typing import Generator

def primes() -> Generator[int, None, None]:
    """Prime number generator."""
    yield 2
    found: set[int] = set()
    cur = 1
    while True:
        cur += 2
        if any(cur % factor == 0 for factor in found):
            continue
        found.add(cur)
        yield cur

def prime(number: int) -> int:
    """Return the n'th prime."""
    if number <= 0:
        raise ValueError("there_is_no_zeroth_prime")
    gen = primes()
    return [next(gen) for _ in range(number)][-1]
