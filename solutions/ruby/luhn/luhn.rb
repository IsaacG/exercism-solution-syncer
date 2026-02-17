module Luhn
  def self.valid?(number)
    number = number.delete(' ').reverse
    return false if number.length == 1
    return false if number.match?(/[^0-9]/)

    sum = number.chars.map(&:to_i).each_with_index.sum do |num, index|
      unless (index % 2).zero?
        num *= 2
        num -= 9 if num > 9
      end
      num
    end

    (sum % 10).zero?
  end
end
