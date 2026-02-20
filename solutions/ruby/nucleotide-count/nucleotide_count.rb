class Counter
  private

  attr_writer :histogram

  def initialize(dna)
    count = %w[A C G T].map { |i| [i, 0] }.to_h
    dna.chars.each do |char|
      raise ArgumentError, 'Invalid input' unless count.include?(char)

      count[char] += 1
    end
    self.histogram = count
  end

  public

  attr_reader :histogram
end

module Nucleotide
  def self.from_dna(dna)
    Counter.new(dna)
  end
end
