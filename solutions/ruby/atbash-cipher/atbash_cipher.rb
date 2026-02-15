module Atbash
  LETTERS = ('a'..'z').to_a.freeze
  DIGITS = ('0'..'9').to_a.freeze
  ALNUM = LETTERS + DIGITS

  def self.convert(data)
    data.downcase.tr(LETTERS.join, LETTERS.reverse.join).chars.select { |i| ALNUM.include?(i) }
  end

  def self.decode(data)
    convert(data).join
  end

  def self.encode(data)
    convert(data).each_slice(5).map(&:join).join(' ')
  end
end
