require 'date'

class Meetup
  WDAYS = Date::DAYNAMES.map { |name| name.downcase.to_sym }

  def initialize(month, year)
    @month = month
    @year = year
  end

  def day(dow, week)
    if week == :last
      d = Date.new(@year, @month, 1).next_month.prev_day
      shift = -1
    else
      offset = { first: 1, second: 8, third: 15, fourth: 22, fifth: 29, teenth: 13 }[week]
      d = Date.new(@year, @month, offset)
      shift = +1
    end

    d += shift while WDAYS[d.wday] != dow
    d
  end
end
