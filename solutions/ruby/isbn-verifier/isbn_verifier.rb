module IsbnVerifier
  NUMBERS = ('0'..'9').to_a + ['10']

  def self.valid?(data)
    values = data.delete('-').chars
    values[-1] = '10' if values[-1] == 'X'
    return false unless values.size == 10 && values.all? { |v| NUMBERS.include?(v) }

    (values.each_index.map { |i| values[i].to_i * (10 - i) }.sum % 11).zero?
  end
end
