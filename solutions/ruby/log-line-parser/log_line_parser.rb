class LogLineParser
  private

  attr_accessor :level, :log_message

  def initialize(line)
    self.level, self.log_message = line.split(': ')
  end

  public

  def message
    log_message.strip
  end

  def log_level
    log_level = level.downcase
    log_level[1, log_level.length - 2]
  end

  def reformat
    "#{message} (#{log_level})"
  end
end
