module Luhn
  def self.valid?(identifier)
    identifier = identifier.delete(' ').reverse
    return false if identifier.length == 1
    return false if identifier.match?(/[^0-9]/)

    sum = identifier.chars.map(&:to_i).each_with_index.sum do |digit, i|
      if i.odd?
        digit *= 2
        digit -= 9 if digit > 9
      end
      digit
    end

    (sum % 10).zero?
  end
end
