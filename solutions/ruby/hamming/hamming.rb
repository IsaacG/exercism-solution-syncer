module Hamming
  def self.compute(strand_a, strand_b)
    raise ArgumentError, 'Different lengths' if strand_a.length != strand_b.length

    strand_a.chars.zip(strand_b.chars).count { |a, b| a != b }
  end
end
