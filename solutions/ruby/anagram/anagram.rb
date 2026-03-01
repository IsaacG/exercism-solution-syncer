class Anagram
  private

  attr_accessor :word, :word_count

  def initialize(word)
    self.word = word.downcase
    self.word_count = count(word)
  end

  def count(word)
    word.downcase.chars.tally
  end

  public

  def match(candidates)
    candidates.select { |candidate| candidate.downcase != word && count(candidate) == word_count }
  end
end
