def abbreviate(words):
    out = ""
    take = True
    for i in words:
        if i in "-_ ":
            take = True
        elif take:
            take = False
            out += i
    return out.upper()
