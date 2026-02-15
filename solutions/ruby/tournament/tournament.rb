module Tournament
  TEMPLATE = "%<name>-30s | %<played>2s | %<win>2s | %<draw>2s | %<loss>2s | %<points>2s\n".freeze

  def self.initialize_teams(input)
    teams = {}
    input
      .lines(chomp: true)
      .map { |line| line.split(';')[0, 2] }
      .flatten
      .uniq
      .each { |team| teams[team] = { name: team, win: 0, loss: 0, draw: 0 } }
    teams
  end

  def self.compute(input)
    teams = initialize_teams(input)

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

    teams
  end

  def self.tally(input)
    teams = compute(input)
    points = teams.transform_values { |team| 3 * team[:win] + 1 * team[:draw] }

    teams
      .each_value
      .sort_by { |team| [-points[team[:name]], team[:name]] }
      .map do |team|
        format(
          TEMPLATE,
          {
            name: team[:name],
            played: %i[win loss draw].sum { |i| team[i] },
            win: team[:win],
            loss: team[:loss],
            draw: team[:draw],
            points: points[team[:name]]
          }
        )
      end
      .unshift(format(TEMPLATE, { name: 'Team', played: 'MP', win: 'W', loss: 'L', draw: 'D', points: 'P' }))
      .join
  end
end
