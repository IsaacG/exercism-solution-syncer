class DndCharacter
  attr_accessor :strength
  attr_accessor :dexterity
  attr_accessor :constitution
  attr_accessor :intelligence
  attr_accessor :wisdom
  attr_accessor :charisma
  attr_accessor :hitpoints

  def self.modifier(v)
    return (v - 10) / 2
  end

  def roll_attrib
    (1..4).to_a.map{ Random.rand(1..6) }.sort[1,3].sum
  end

  def initialize
    self.strength = roll_attrib
    self.dexterity = roll_attrib
    self.constitution = roll_attrib
    self.intelligence = roll_attrib
    self.wisdom = roll_attrib
    self.charisma = roll_attrib
    self.hitpoints = 10 + DndCharacter.modifier(self.constitution)
  end
end
