class AssemblyLine
  def initialize(speed)
    @speed = speed
  end

  def production_rate_per_hour
    rates = [[4, 1.00], [8, 0.90], [9, 0.80]]
    rate = (rates.find { |speed_rate| @speed <= speed_rate[0] } || [nil, 0.77])[1]
    221 * @speed * rate
  end

  def working_items_per_minute
    (production_rate_per_hour / 60).floor
  end
end
