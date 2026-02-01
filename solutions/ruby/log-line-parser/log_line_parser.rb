class LogLineParser
  def initialize(line)
    @_level, @_msg = line.split(': ')
  end

  attr_accessor :_level, :_msg

  def message
    @_msg.strip
  end

  def log_level
    level = @_level.downcase
    level[1, level.length - 2]
  end

  def reformat
    "#{message} (#{log_level})"
  end
end
