module Chess
  RANKS = (1..8).freeze
  FILES = ('A'..'H').freeze

  def self.valid_square?(rank, file)
    RANKS.include?(rank) and FILES.include?(file)
  end

  def self.nickname(first_name, last_name)
    first_name.upcase[0...2] + last_name.upcase[-2..]
  end

  def self.move_message(first_name, last_name, square)
    unless valid_square?(square[1].to_i, square[0])
      return "#{nickname(first_name, last_name)} attempted to move to #{square}, but that is not a valid square"
    end

    "#{nickname(first_name, last_name)} moved to #{square}"
  end
end
