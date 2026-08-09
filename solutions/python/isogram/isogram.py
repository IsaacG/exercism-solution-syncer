def is_isogram(word):
    seen = set[str]()
    for char in word.lower():
        if not char.isalpha():
            continue
        if char in seen:
            return False
        seen.add(char)
    return True
