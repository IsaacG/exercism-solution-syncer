def response(hey):
    hey = hey.strip()
    if not hey:
        return "Fine. Be that way!"
    if hey.endswith("?"):
        if hey.isupper():
            return "Calm down, I know what I'm doing!"
        return "Sure."
    if hey.isupper():
        return "Whoa, chill out!"
    return "Whatever."
