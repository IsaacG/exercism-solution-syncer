class Microwave
  def initialize(input)
    @sec = (input / 100 * 60) + (input % 100)
  end

  def timer
    format('%02<hour>d:%02<minute>d', { hour: @sec / 60, minute: @sec % 60 })
  end
end
