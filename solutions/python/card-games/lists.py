def get_rounds(number):
    return list(range(number, number + 3))


def concatenate_rounds(rounds_1, rounds_2):
    return rounds_1 + rounds_2


def list_contains_round(rounds, number):
    return number in rounds


def card_average(hand):
    return sum(hand) / len(hand)


def approx_average_is_average(hand):
    a = (hand[0] + hand[-1]) / 2
    b = hand[int(len(hand) / 2)]
    return card_average(hand) in {a, b}


def average_even_is_average_odd(hand):
    return card_average(hand[0::2]) == card_average(hand[1::2])


def maybe_double_last(hand):
    if hand[-1] == 11:
        hand[-1] *= 2
    return hand
