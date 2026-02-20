module Luhn
  def self.valid?(identifier)
    identifier = identifier.delete(' ').reverse
    return false if identifier.length == 1 || identifier.match?(/[^0-9]/)

    sum = identifier.chars.each_with_index.sum do |digit, i|
      digit = digit.to_i
      if i.odd?
        digit *= 2
        digit -= 9 if digit > 9
      end
      digit
    end

    (sum % 10).zero?
  end
end
