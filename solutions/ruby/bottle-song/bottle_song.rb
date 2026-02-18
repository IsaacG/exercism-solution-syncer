module BottleSong
  NUMBER = %w[no one two three four five six seven eight nine ten].freeze

  def self.bottles(count)
    subject = count == 1 ? 'bottle' : 'bottles'
    "#{NUMBER[count]} green #{subject}"
  end

  private_constant :NUMBER
  private_class_method :bottles

  def self.recite(start, count)
    Range.new(0, count - 1).map do |step|
      remaining = start - step - 1
      [
        "#{bottles(start - step).capitalize} hanging on the wall",
        "#{bottles(start - step).capitalize} hanging on the wall",
        'And if one green bottle should accidentally fall',
        "There'll be #{bottles(remaining)} hanging on the wall."
      ].join(",\n")
    end.join("\n\n") + "\n"
  end
end
