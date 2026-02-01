module CollatzConjecture
  def self.steps(number)
    raise ArgumentError, 'Invalid input' unless number >= 1

    count = 0
    while number > 1
      if number.even?
        number /= 2
      else
        number = number * 3 + 1
      end
      count += 1
    end
    count
  end
end
