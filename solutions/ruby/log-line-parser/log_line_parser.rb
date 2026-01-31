class LogLineParser
  def initialize(line)
    @line = line
  end

  def message
    return @line.split(': ')[1].strip
  end

  def log_level
    level = @line.split(': ')[0].downcase
    return level[1, level.length - 2]
  end

  def reformat
    return "#{message} (#{log_level})"
  end
end
