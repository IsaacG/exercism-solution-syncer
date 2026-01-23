def line_up(name, number):
    """Format a message.

    Numbers ending in 1 (unless ending in 11) → "st"
    Numbers ending in 2 (unless ending in 12) → "nd"
    Numbers ending in 3 (unless ending in 13) → "rd"
    All other numbers → "th"
    """
    follows_rule = (number % 100) // 10 != 1
    match number % 10:
        case 1 if follows_rule:
            suffix = "st"
        case 2 if follows_rule:
            suffix = "nd"
        case 3 if follows_rule:
            suffix = "rd"
        case _:
            suffix = "th"
    return f"{name}, you are the {number}{suffix} customer we serve today. Thank you!"
