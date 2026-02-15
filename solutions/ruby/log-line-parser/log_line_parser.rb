class LogLineParser

  private

  attr_writer :log_level, :message

  def initialize(line)
    level, message = line.split(': ')
    self.log_level = level.downcase.delete_prefix('[').delete_suffix(']')
    self.message = message.strip
  end

  public

  attr_reader :log_level, :message

  def reformat
    "#{message} (#{log_level})"
  end

end
