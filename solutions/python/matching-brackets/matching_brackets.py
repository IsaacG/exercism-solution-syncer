MATCHING = {")": "(", "]": "[", "}": "{"}


def is_paired(input_string):
    unmatched_context = []
    for character in input_string:
        if character in MATCHING.values():
            unmatched_context.append(character)
        elif character in MATCHING:
            if not unmatched_context:
                return False
            if unmatched_context[-1] == MATCHING[character]:
                unmatched_context.pop()
            else:
                return False
    return not unmatched_context
