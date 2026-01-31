class EliudsEggs
  def self.egg_count(number)
    count = 0
    while number > 0
      if number & 1 == 1
        count += 1
      end
      number >>= 1
    end
    return count
  end
end
