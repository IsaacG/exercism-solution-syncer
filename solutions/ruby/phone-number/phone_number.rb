module PhoneNumber
  def self.clean(number)
    number = number.delete('-+(). ')
    number = number.delete_prefix('1') if number.length == 11

    return nil if number.match?(/[^0-9]/) ||
                  number.length != 10 ||
                  %w[0 1].include?(number[0]) ||
                  %w[0 1].include?(number[3])

    number
  end
end
