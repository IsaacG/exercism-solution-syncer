class Squares
  private

  attr_accessor :number

  def initialize(number)
    self.number = number
  end

  public

  def square_of_sum()= (1..number).sum**2

  def sum_of_squares()= (1..number).sum { |n| n * n }

  def difference()= square_of_sum - sum_of_squares
end
