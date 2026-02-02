module Acronym
  def self.abbreviate(string)
    string.tr('-_', '  ').split.map { |word| word[0] }.join.upcase
  end
end
