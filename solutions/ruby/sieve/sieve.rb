class Sieve
  private

  attr_writer :primes

  def initialize(limit)
    primes = []
    sieve = Array.new(limit + 1, true)
    (2..limit).each do |number|
      next unless sieve[number]

      primes.append(number)
      (number..limit).step(number) { |mark| sieve[mark] = false }
    end
    self.primes = primes
  end

  public

  attr_reader :primes
end
