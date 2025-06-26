"""Flowerfield."""

ADJACENT = [complex(x, y) for x in (-1, 0, 1) for y in (-1, 0, 1) if x or y]


def annotate(garden: list[str]) -> list[str]:
    """Annotate a garden."""
    height = len(garden)
    width = len(garden[0]) if garden else 0

    if not all(len(row) == width for row in garden):
        raise ValueError("The board is invalid with current input.")

    flowers = set()
    for y, line in enumerate(garden):
        for x, val in enumerate(line):
            if val == "*":
                flowers.add(complex(x, y))
            elif val != " ":
                raise ValueError("The board is invalid with current input.")

    def cell_value(point: complex) -> str:
        """Return the value for one square."""
        if point in flowers:
            return "*"
        count = sum(point + offset in flowers for offset in ADJACENT)
        return str(count) if count else " "

    return [
        "".join(cell_value(complex(x, y)) for x in range(width))
        for y in range(height)
    ]
