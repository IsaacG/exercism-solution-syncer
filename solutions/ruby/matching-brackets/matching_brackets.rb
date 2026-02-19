module Brackets
  PAIRS = '{}()[]<>'.chars.each_slice(2).to_h
  OPEN = PAIRS.keys
  CLOSE = PAIRS.values

  def self.paired?(data)
    stack = []
    data.chars.each do |char|
      if OPEN.include?(char)
        stack.append(char)
      elsif CLOSE.include?(char)
        return false unless PAIRS[stack.pop] == char
      end
    end
    stack.empty?
  end
end
