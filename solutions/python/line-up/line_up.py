def line_up(name, number):
    """Format a message.

    Numbers ending in 1 (unless ending in 11) → "st"
    Numbers ending in 2 (unless ending in 12) → "nd"
    Numbers ending in 3 (unless ending in 13) → "rd"
    All other numbers → "th"
    """
    suffix = "th" if (number % 100) // 10 == 1 else {1: "st", 2: "nd", 3: "rd"}.get(number % 10, "th")
    return f"{name}, you are the {number}{suffix} customer we serve today. Thank you!"
