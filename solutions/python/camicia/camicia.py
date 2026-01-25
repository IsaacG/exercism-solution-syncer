import collections


def simulate_game(player_a, player_b):
    cards = [["-" if i.isdecimal() else i for i in v] for v in [player_a, player_b]]
    hands = {a: collections.deque(reversed(b)) for a, b in zip([True, False], cards)}
    middle = []
    seen = set()
    p1 = True
    cards = 0
    tricks = 0
    count = {"A": 4, "K": 3, "Q": 2, "J": 1}
    status = "finished"

    # Play cards until one player holds all the cards.
    while all(hands.values()) or middle:
        # Check for a loop.
        fp = (tuple(hands[True]), tuple(hands[False]))
        if fp in seen:
            status = "loop"
            break
        seen.add(fp)

        # Check if the player is out of cards and cannot play.
        if not hands[p1]:
            p1 = not p1
            hands[p1].extendleft(middle)
            tricks += 1
            middle.clear()
            break

        # Add a card to the middle pile.
        card = hands[p1].pop()
        middle.append(card)
        cards += 1
        p1 = not p1

        # Resolve a penalty.
        if card != "-":
            penalty = count[card]
            while penalty and hands[p1]:
                card = hands[p1].pop()
                middle.append(card)
                cards += 1
                penalty -= 1
                if card != "-":
                    p1 = not p1
                    penalty = count.get(card, 0)

            # Collect the pile.
            p1 = not p1
            hands[p1].extendleft(middle)
            tricks += 1
            middle.clear()

    return {"status": status, "cards": cards, "tricks": tricks}
