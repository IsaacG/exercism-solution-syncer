class Matrix
  private

  attr_accessor :matrix

  def initialize(matrix)
    self.matrix = matrix.lines.map { |line| line.strip.split.map(&:to_i) }
  end

  public

  def row(i)= matrix[i - 1]

  def column(i)= matrix.map { |row| row[i - 1] }
end
