class Robot
  private

  attr_writer :bearing, :coordinates
  attr_accessor :pos_x, :pos_y

  private_constant

  BEARINGS = %i[east north west south].freeze
  OFFSETS = {
    east: [1, 0],
    north: [0, 1],
    west: [-1, 0],
    south: [0, -1]
  }.freeze

  def initialize
    at(0, 0)
    orient(:east)
  end

  public

  attr_reader :bearing, :coordinates

  def at(pos_x, pos_y)
    self.coordinates = [pos_x, pos_y]
  end

  def orient(orientation)
    raise ArgumentError, 'Invalid bearing' unless BEARINGS.include?(orientation)

    self.bearing = orientation
  end

  def turn_right
    self.bearing = BEARINGS[(BEARINGS.index(bearing) + 3) % 4]
  end

  def turn_left
    self.bearing = BEARINGS[(BEARINGS.index(bearing) + 1) % 4]
  end

  def advance
    self.coordinates = coordinates.zip(OFFSETS[bearing]).map(&:sum)
  end
end

class Simulator
  private_constant

  COMMANDS = { 'R' => :turn_right, 'L' => :turn_left, 'A' => :advance }.freeze

  def place(robot, x: nil, y: nil, direction: nil)
    robot.at(x, y) unless x.nil? || y.nil?
    robot.orient(direction) unless direction.nil?
  end

  def instructions(commands)
    commands.chars.map { |char| COMMANDS[char] }
  end

  def evaluate(robot, commands)
    instructions(commands).each { |command| robot.send(command) }
  end
end
