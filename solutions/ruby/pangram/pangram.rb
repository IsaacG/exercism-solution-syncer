module Pangram
  def self.pangram?(sentence)
    Set.new(sentence.downcase.chars) >= Set.new('a'..'z')
  end
end
