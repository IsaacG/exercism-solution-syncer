class Tournament
  private

  TEMPLATE = "%<name>-30s | %<played>2s | %<win>2s | %<draw>2s | %<loss>2s | %<points>2s\n".freeze

  attr_accessor :teams

  def initialize(input)
    self.teams = Hash.new { |hash, team| hash[team] = { name: team, win: 0, loss: 0, draw: 0 } }
    input.strip.each_line(chomp: true) { |line| record_match(*line.split(';')) }
    add_summaries
  end

  def add_summaries
    teams.each_value do |v|
      v[:played] = v.values_at(:win, :loss, :draw).sum
      v[:points] = 3 * v[:win] + 1 * v[:draw]
    end
  end

  def record_win(team_a, team_b)
    teams[team_a][:win] += 1
    teams[team_b][:loss] += 1
  end

  def record_draw(team_a, team_b)
    teams[team_a][:draw] += 1
    teams[team_b][:draw] += 1
  end

  def record_match(team_a, team_b, result)
    if result == 'draw'
      record_draw(team_a, team_b)
    else
      team_a, team_b = team_b, team_a if result == 'loss'
      record_win(team_a, team_b)
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
