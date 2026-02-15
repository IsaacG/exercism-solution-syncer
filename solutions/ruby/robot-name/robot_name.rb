class Robot
  private

  @@names_used = Set.new
  LETTERS = ('A'..'Z').to_a.freeze
  DIGITS = ('0'..'9').to_a.freeze

  attr_writer :name

  def initialize
    reset
  end

  def gen_name
    (0...2).map { LETTERS[rand(LETTERS.length)] }.join +
      (0...3).map { DIGITS[rand(DIGITS.length)] }.join
  end

  public

  attr_reader :name

  def reset
    loop do
      self.name = gen_name
      break unless @@names_used.include?(name)
    end
    @@names_used.add(name)
  end

  def self.forget
    @@names_used = Set.new
  end

end
