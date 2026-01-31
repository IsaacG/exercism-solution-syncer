require 'date'

class Meetup
  def initialize(month, year)
    @month = month
    @year = year
  end

  def day(dow, week)
    wdays = [:sunday, :monday, :tuesday, :wednesday, :thursday, :friday, :saturday]

    if week == :last
      d = Date.new(@year, @month, 1).next_month.prev_day
      shift = -1
    else
      offset = {first: 1, second: 8, third: 15, fourth: 22, fifth: 29, teenth: 13}[week]
      d = Date.new(@year, @month, offset)
      shift = +1
    end

    while wdays[d.wday] != dow
      d += shift
    end
    return d
  end
end

