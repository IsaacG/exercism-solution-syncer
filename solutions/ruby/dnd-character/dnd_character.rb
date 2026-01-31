class DndCharacter
  attr_accessor :strength, :dexterity, :constitution, :intelligence, :wisdom, :charisma, :hitpoints

  def self.modifier(attrib)
    (attrib - 10) / 2
  end

  def roll_attrib
    (1..4).to_a.map { Random.rand(1..6) }.sort[1, 3].sum
  end

  def initialize
    self.strength = roll_attrib
    self.dexterity = roll_attrib
    self.constitution = roll_attrib
    self.intelligence = roll_attrib
    self.wisdom = roll_attrib
    self.charisma = roll_attrib
    self.hitpoints = 10 + DndCharacter.modifier(constitution)
  end
end
