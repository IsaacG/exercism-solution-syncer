"""Convert an int to a Roman numeral."""


# The 10's, 5's and 1's position chars for 1, 10, 100, 1000.
DIGIT_CHARS = ["XVI", "CLX", "MDC", "??M"]


def digit(number: int, position: int) -> tuple[str, int]:
    """Compute the "digit" numeral and remainder, eg 100's numeral."""
    scale = 10 ** position
    chars = DIGIT_CHARS[position]
    parts = {
        9: chars[2] + chars[0],
        5: chars[1],
        4: chars[2] + chars[1],
        1: chars[2]
    }

    out = ""
    for num, numerals in parts.items():
        num *= scale
        while number >= num:
            out += numerals
            number -= num
    return out, number


def roman(number: int) -> str:
    """Return the Roman numeral for a number."""
    out = ""

    for position in range(len(DIGIT_CHARS) - 1, -1, -1):
        part, number = digit(number, position)
        out += part

    return out
