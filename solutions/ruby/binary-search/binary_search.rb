class BinarySearch
  private

  attr_accessor :numbers

  def initialize(numbers)
    self.numbers = numbers
  end

  public

  def search_for(number)
    lower = 0
    upper = numbers.length
    while lower < upper
      pivot = (upper + lower) / 2
      found = numbers[pivot]
      if found == number
        return pivot
      elsif found > number
        upper = pivot
      else
        lower = pivot + 1
      end
    end
  end
end
