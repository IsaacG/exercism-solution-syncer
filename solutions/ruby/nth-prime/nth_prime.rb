class PrimeGenerator
  private

  attr_accessor :primes

  def initialize
    self.primes = [2, 3]
  end

  public

  def nth(num)
    raise ArgumentError if num < 1

    candidate = primes[-1]
    while primes.length < num
      candidate += 2
      primes.append(candidate) if primes.all? { |prime| (candidate % prime).positive? }
    end
    primes[num - 1]
  end
end

GENERATOR = PrimeGenerator.new

module Prime
  def self.nth(num)
    GENERATOR.nth(num)
  end
end
