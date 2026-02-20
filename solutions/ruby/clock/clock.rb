class Clock
  private

  attr_writer :minutes

  def initialize(hour: 0, minute: 0)
    self.minutes = (hour * 60 + minute) % (24 * 60)
  end

  protected

  attr_reader :minutes

  public

  def to_s
    format('%02<hours>d:%02<minutes>d', { hours: minutes / 60, minutes: minutes % 60 })
  end

  def ==(other)
    other.respond_to?(:minutes, true) && minutes == other.minutes
  end

  def +(other)
    Clock.new(minute: minutes + other.minutes)
  end

  def -(other)
    Clock.new(minute: minutes - other.minutes)
  end
end
