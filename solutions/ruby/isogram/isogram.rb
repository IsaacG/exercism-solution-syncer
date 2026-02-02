module Isogram
  def self.isogram?(word)
    seen = Set.new
    word.downcase.delete(' -').chars.each do |letter|
      return false if seen.include?(letter)

      seen.add(letter)
    end
  end
end
