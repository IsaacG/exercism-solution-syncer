class Microwave
  def initialize(input)
    @sec = (input / 100 * 60) + (input % 100)
  end

  def timer
    "%02d:%02d" % [@sec / 60, @sec % 60]
  end
end
