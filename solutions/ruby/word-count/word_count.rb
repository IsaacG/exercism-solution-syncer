class Phrase
  private

  attr_writer :word_count

  def initialize(phrase)
    self.word_count = Hash.new(0)
    phrase
      .tr(',_', '  ')
      .downcase
      .split
      .each do |word|
        word = word.match(/[[:alnum:]](.*[[:alnum:]])?/).match(0)
        word_count[word] += 1
      end
  end

  public

  attr_reader :word_count
end
