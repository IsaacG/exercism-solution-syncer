module Prime
  def self.nth(count)
    raise ArgumentError if count < 1

    found = [2]
    candidate = 1
    while found.length < count
      candidate += 2
      is_prime = true
      limit = Integer.sqrt(candidate)

      found.each do |prime|
        break if prime > limit

        if (candidate % prime).zero?
          is_prime = false
          break
        end
      end
      found.append(candidate) if is_prime
    end
    found[count - 1]
  end
end
