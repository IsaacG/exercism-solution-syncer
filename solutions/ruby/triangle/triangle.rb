class Triangle
  private

  attr_accessor :sides, :count

  def initialize(sides)
    self.sides = sides
    self.count = Set.new(sides).length
  end

  def valid?
    sides.min.positive? && 2 * sides.max < sides.sum
  end

  public

  def equilateral?
    valid? && count == 1
  end

  def isosceles?
    valid? && !scalene?
  end

  def scalene?
    valid? && count == 3
  end
end
