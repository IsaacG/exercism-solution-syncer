class Cipher
  private

  attr_accessor :shifts
  attr_writer :key

  LETTERS = ('a'..'z').to_a.freeze
  ORD_A = 'a'.ord

  def initialize(key = nil)
    key = (0..100).map { (rand(LETTERS.length) + ORD_A).chr }.join if key.nil?

    raise ArgumentError, 'Missing key' if key == ''
    raise ArgumentError, 'Invalid key' unless key.chars.to_a.all? { |letter| LETTERS.include?(letter) }

    self.key = key
    self.shifts = key.chars.map { |letter| LETTERS.index(letter) }
  end

  def encode_with(text, shifts)
    text.chars.zip(shifts.cycle).map { |letter, shift| ((letter.ord - ORD_A + shift) % 26 + ORD_A).chr }.join
  end

  public

  attr_reader :key

  def encode(text)= encode_with(text, shifts)
  def decode(text)= encode_with(text, shifts.map(&:-@))
end
