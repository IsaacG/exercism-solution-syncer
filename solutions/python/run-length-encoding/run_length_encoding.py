import re

REP_RE = re.compile(r"(\d*)(\D)")


def decode(string):
    out = ""
    for match in REP_RE.finditer(string):
        count, char = match.groups()
        count = int(count) if count else 1
        out += char * count
    return out


def encode(string):
    out = ""
    count = 0
    chars = list(string)
    for prev, cur in zip([""] + chars, chars + [""]):
        if cur == prev:
            count += 1
        else:
            out += f"{count}{prev}" if count > 1 else prev
            count = 1
    return out
