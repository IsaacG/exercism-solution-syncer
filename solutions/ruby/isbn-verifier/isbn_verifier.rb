module IsbnVerifier
  NUMBERS = (0..10).map(&:to_s)

  def self.valid?(data)
    values = data.delete('-').chars
    values[-1] = '10' if values[-1] == 'X'
    return false unless values.size == 10 && values.all? { |v| NUMBERS.include?(v) }

    (values.map.with_index { |value, i| value.to_i * (10 - i) }.sum % 11).zero?
  end
end
