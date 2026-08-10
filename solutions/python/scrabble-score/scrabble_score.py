VALUES = {
    'AEIOULNRST': 1,
    'DG': 2,
    'BCMP': 3,
    'FHVWY': 4,
    'K': 5,
    'JX': 8,
    'QZ': 10,
}

# Expand the keys in VALUES into letters.
LETTERS = {l: v for c, v in VALUES.items() for l in c}


def score(word):
    return sum(LETTERS[w] for w in word.upper())
