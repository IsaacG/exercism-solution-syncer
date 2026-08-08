import collections


def find_anagrams(word: str, candidates: list[str]) -> list[str]:
    word_count = collections.Counter(word.lower())
    return [
        c for c in candidates
        if c.lower() != word.lower() and word_count == collections.Counter(c.lower())
    ]
