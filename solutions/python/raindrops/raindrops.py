SOUNDS = [(3, "Pling"), (5, "Plang"),  (7, "Plong")]


def convert(number):
    return "".join(s for f, s in SOUNDS if number % f == 0) or str(number)
