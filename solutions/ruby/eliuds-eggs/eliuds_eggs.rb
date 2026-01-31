class EliudsEggs
  def self.egg_count(number)
    count = 0
    while number.positive?
      count += 1 if number & 1 == 1
      number >>= 1
    end
    count
  end
end
