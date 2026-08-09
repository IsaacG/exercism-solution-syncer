import random

def modifier(n):
    return (n - 10) // 2


TRAITS = ('strength', 'dexterity', 'constitution', 'intelligence', 'wisdom', 'charisma')


class Character:

    def __init__(self):
        for trait in TRAITS:
            setattr(self, trait, self.ability())

        self.hitpoints = 10 + modifier(self.constitution)

    def ability(self):
        dice = sorted([random.randint(1, 6) for _ in range(4)], reverse=True)
        return sum(dice[:3])
