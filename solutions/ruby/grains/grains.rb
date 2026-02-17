module Grains
  def self.square(square)
    raise ArgumentError, 'Invalid input' unless (1..64).include?(square)

    2**(square - 1)
  end

  def self.total
    2**64 - 1
  end
end
