class SumOfMultiples
  private

  attr_accessor :factors

  def initialize(*factors)
    self.factors = factors
  end

  public

  def to(limit)
    Range.new(1, limit - 1).select { |number| factors.any? { |factor| (number % factor).zero? } }.sum
  end
end
