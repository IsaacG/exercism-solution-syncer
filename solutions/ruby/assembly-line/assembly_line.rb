class AssemblyLine
  def initialize(speed)
    @speed = speed
  end

  def production_rate_per_hour
    if 1 <= @speed && @speed <= 4
      rate = 1.00
    elsif 5 <= @speed && @speed <= 8
      rate = 0.90
    elsif @speed == 9
      rate = 0.80
    else
      rate = 0.77
    end
    return 221 * @speed * rate
  end

  def working_items_per_minute
    return (production_rate_per_hour / 60).floor
  end
end
