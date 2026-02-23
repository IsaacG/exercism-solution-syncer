module RotationalCipher
  LETTERS = ('a'..'z').to_a

  def self.rotate(string, distance)
    rotated = LETTERS.rotate(distance).join('')
    from = LETTERS.join('') + LETTERS.join('').upcase
    to = rotated + rotated.upcase
    string.tr(from, to)
  end
end
