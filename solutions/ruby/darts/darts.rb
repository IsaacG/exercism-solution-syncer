class Darts
  include Math

  private

  attr_accessor :distance

  SCORES = {
    0..1 => 10,
    1..5 => 5,
    5..10 => 1,
    10.. => 0
  }.freeze

  def initialize(pos_x, pos_y)
    self.distance = Math.sqrt(pos_x * pos_x + pos_y * pos_y)
  end

  public

  def score
    SCORES.each { |range, points| return points if range.include?(distance) }
  end
end
