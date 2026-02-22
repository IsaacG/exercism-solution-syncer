class Anagram
  private

  attr_accessor :word, :word_count

  def initialize(word)
    self.word = word.downcase
    self.word_count = count(word)
  end

  def count(word)
    count = Hash.new(0)
    word.downcase.chars { |c| count[c] += 1 }
    count
  end

  public

  def match(candidates)
    candidates.select { |candidate| candidate.downcase != word && count(candidate) == word_count }
  end
end
