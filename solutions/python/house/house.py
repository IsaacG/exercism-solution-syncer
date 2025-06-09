"""Generate House verses."""

PREFIX = "This is the "
SUFFIX = "Jack built."
VERB_AND_SUBJECTS = (
    ("lay in", "house"),
    ("ate", "malt"),
    ("killed", "rat"),
    ("worried", "cat"),
    ("tossed", "dog"),
    ("milked", "cow with the crumpled horn"),
    ("kissed", "maiden all forlorn"),
    ("married", "man all tattered and torn"),
    ("crowed in the morn that woke", "priest all shaven and shorn"),
    ("kept", "rooster"),
    ("belonged to", "farmer sowing his corn"),
    ("", "horse and the hound and the horn"),
)


def recite(start_verse: int, end_verse: int) -> list[str]:
    """Return a number of verses."""
    return [verse(i) for i in range(start_verse - 1, end_verse)]


def verse(verse_num: int) -> str:
    """Return one verse of the song."""
    parts = [VERB_AND_SUBJECTS[verse_num][1]]
    parts += [
        " the ".join(VERB_AND_SUBJECTS[i])
        for i in range(verse_num - 1, -1, -1)
    ]
    parts += [SUFFIX]
    return PREFIX + " that ".join(parts)
