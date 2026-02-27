module PerfectNumber
  def self.classify(number)
    raise ArgumentError, 'Classification is only possible for positive integers.' unless number.positive?

    sum = 0
    Range.new(1, number - 1).each { |i| sum += i if (number % i).zero? }
    return 'perfect' if sum == number
    return 'abundant' if sum > number

    'deficient'
  end
end
