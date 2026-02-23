module Tournament
  TEMPLATE = "%<name>-30s | %<played>2s | %<win>2s | %<draw>2s | %<loss>2s | %<points>2s\n".freeze

  def self.compute(input)
    teams = Hash.new { |hash, team| hash[team] = { name: team, win: 0, loss: 0, draw: 0 } }

    input.strip.each_line(chomp: true) do |line|
      a, b, result = line.split(';')
      a, b = b, a if result == 'loss'
      if result == 'draw'
        teams[a][:draw] += 1
        teams[b][:draw] += 1
      else
        teams[a][:win] += 1
        teams[b][:loss] += 1
      end
    end

    teams.each_value do |v|
      v[:played] = %i[win loss draw].sum { |i| v[i] }
      v[:points] = 3 * v[:win] + 1 * v[:draw]
    end

    teams
  end

  def self.tally(input)
    teams = compute(input)
    points = teams.transform_values { |team| 3 * team[:win] + 1 * team[:draw] }

    teams
      .each_value
      .sort_by { |team| [-points[team[:name]], team[:name]] }
      # This line fails with,
      # TypeError: no implicit conversion of Symbol into Integer
      # .sort_by { |team| [-team[:name][:points], team[:name]] }
      .map { |team| format(TEMPLATE, team) }
      .unshift(format(TEMPLATE, { name: 'Team', played: 'MP', win: 'W', loss: 'L', draw: 'D', points: 'P' }))
      .join
  end
end
