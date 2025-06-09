"""Parse resistor colors."""

import enum


class Colors(enum.IntEnum):
    """Color mapping."""
    black = 0
    brown = 1
    red = 2
    orange = 3
    yellow = 4
    green = 5
    blue = 6
    violet = 7
    grey = 8
    white = 9


def label(colors: list[str]) -> str:
    """Return the value of a resistor color."""
    val = 0
    for color in colors[:2]:
        val = val * 10 + Colors[color].value
    val *= 10 ** Colors[colors[2]].value
    if val % 1000 == 0:
        return f"{val // 1000} kiloohms"
    return f"{val} ohms"
