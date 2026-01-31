class LogLineParser
  def initialize(line)
    @line = line
  end

  def message
    @line.split(': ')[1].strip
  end

  def log_level
    level = @line.split(': ')[0].downcase
    level[1, level.length - 2]
  end

  def reformat
    "#{message} (#{log_level})"
  end
end
