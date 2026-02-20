module PrimeFactors
  module_function

  def check_prime(found, candidate)
    limit = Integer.sqrt(candidate)
    found.each do |prime|
      return true if prime > limit
      return false if (candidate % prime).zero?
    end
    true
  end

  def primes
    yield 2
    found = []
    (3..).step(2).each do |candidate|
      if check_prime(found, candidate)
        found.append(candidate)
        yield candidate
      end
    end
  end

  def of(number)
    result = []
    primes do |prime|
      return result if number == 1

      while (number % prime).zero?
        number /= prime
        result.append(prime)
      end
    end
  end
end
