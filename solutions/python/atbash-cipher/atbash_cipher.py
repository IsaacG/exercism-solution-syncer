import string

mapping = dict(zip(string.ascii_lowercase, string.ascii_lowercase[::-1]))
mapping.update({i: i for i in string.digits})


def encode(text):
    text = decode(text)
    return " ".join("".join(text[i:i + 5]) for i in range(0, len(text), 5))


def decode(text):
    text = text.lower()
    text = (t for t in text if t.isalnum())
    text = [mapping[t] for t in text]
    return "".join(text)
