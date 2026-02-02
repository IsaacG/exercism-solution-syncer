class ResistorColorTrio
  private

  COLORS = %w[black brown red orange yellow green blue violet grey white].freeze
  UNITS = %w[ohms kiloohms megaohms gigaohms].freeze

  attr_writer :label

  def initialize(colors)
    tens, ones, power = colors
    base = COLORS.index(tens) * 10 + COLORS.index(ones)
    resistance = base * 10**COLORS.index(power)

    shifts = 0
    while !resistance.zero? && (resistance % 1000).zero?
      shifts += 1
      resistance /= 1000
    end
    self.label = "Resistor value: #{resistance} #{UNITS[shifts]}"
  end

  public

  attr_reader :label
end
