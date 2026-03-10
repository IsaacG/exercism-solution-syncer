class Tournament
  TEMPLATE = "%<name>-30s | %<played>2s | %<win>2s | %<draw>2s | %<loss>2s | %<points>2s\n".freeze
  RESULTS = {
    'draw' => %i[draw draw],
    'win' => %i[win loss],
    'loss' => %i[loss win]
  }.freeze

  private_constant :TEMPLATE, :RESULTS

  private

  attr_accessor :teams

  def initialize(input)
    self.teams = Hash.new { |hash, team| hash[team] = { name: team, win: 0, loss: 0, draw: 0 } }
    input
      .strip
      .lines
      .map { |l| l.strip.split(';') }
      .map { |team_a, team_b, result| [team_a, team_b].zip(RESULTS[result]) }
      .flatten(1)
      .each { |team, outcome| teams[team][outcome] += 1 }
    add_summaries
  end

  def add_summaries
    teams.each_value do |v|
      v[:played] = v.values_at(:win, :loss, :draw).sum
      v[:points] = 3 * v[:win] + 1 * v[:draw]
    end
  end

  public

  def to_s
    teams
      .each_value
      .sort_by { |team| [-team[:points], team[:name]] }
      .map { |team| format(TEMPLATE, team) }
      .unshift(format(TEMPLATE, { name: 'Team', played: 'MP', win: 'W', loss: 'L', draw: 'D', points: 'P' }))
      .join
  end

  def self.tally(input)
    Tournament.new(input).to_s
  end
end
