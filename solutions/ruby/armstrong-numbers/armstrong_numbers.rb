module ArmstrongNumbers
  def self.include?(number)
    remaining = number
    digits = []
    until remaining.zero?
      digits.append(remaining % 10)
      remaining /= 10
    end
    digits.sum { |digit| digit**digits.length } == number
  end
end
