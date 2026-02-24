module Tournament
  TEMPLATE = "%<name>-30s | %<played>2s | %<win>2s | %<draw>2s | %<loss>2s | %<points>2s\n".freeze

  def self.add_summaries(teams)
    teams.each_value do |v|
      v[:played] = %i[win loss draw].sum { |i| v[i] }
      v[:points] = 3 * v[:win] + 1 * v[:draw]
    end
  end

  def self.record_match(teams, team_a, team_b, draw)
    if draw
      teams[team_a][:draw] += 1
      teams[team_b][:draw] += 1
    else
      teams[team_a][:win] += 1
      teams[team_b][:loss] += 1
    end
  end

  def self.process_input(input)
    teams = Hash.new { |hash, team| hash[team] = { name: team, win: 0, loss: 0, draw: 0 } }

    input.strip.each_line(chomp: true) do |line|
      a, b, result = line.split(';')
      a, b = b, a if result == 'loss'
      record_match(teams, a, b, result == 'draw')
    end
    add_summaries(teams)
  end

  def self.tally(input)
    process_input(input)
      .each_value
      .sort_by { |team| [-team[:points], team[:name]] }
      .map { |team| format(TEMPLATE, team) }
      .unshift(format(TEMPLATE, { name: 'Team', played: 'MP', win: 'W', loss: 'L', draw: 'D', points: 'P' }))
      .join
  end
end
