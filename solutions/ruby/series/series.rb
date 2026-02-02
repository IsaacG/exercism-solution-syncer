class Series
  private

  attr_accessor :string

  def initialize(string)
    raise ArgumentError if string.empty?

    self.string = string
  end

  public

  def slices(size)
    raise ArgumentError if size <= 0 || size > string.length

    (0..(string.length - size)).map { |start| string[start, size] }
  end
end
