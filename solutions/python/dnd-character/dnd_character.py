import random

def modifier(n):
  return (n - 10) // 2


class Character:

  def __init__(self):
    random.seed()
    self.strength = self.ability()
    self.dexterity = self.ability()
    self.constitution = self.ability()
    self.intelligence = self.ability()
    self.wisdom = self.ability()
    self.charisma = self.ability()
    self.hitpoints = 10 + modifier(self.constitution)
        

  def ability(self):
    dice = sorted([random.randint(1, 6) for _ in range(4)], reverse=True)
    return sum(dice[0:3])




# vim:ts=2:sw=2:expandtab
