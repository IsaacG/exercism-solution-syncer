import datetime
import enum


class Weekday(enum.IntEnum):
  Monday = 0
  Tuesday = 1
  Wednesday = 2
  Thursday = 3
  Friday = 4
  Saturday = 5
  Sunday = 6


class MeetupDayException(Exception):
  pass


OFFSET = {'1st': 1, '2nd': 8, '3rd': 15, '4th': 22, '5th': 29, 'teenth': 13}


def meetup(year, month, week, day_of_week):
  
  if week in OFFSET:
    try:
      d = datetime.date(year, month, OFFSET[week])
    except ValueError:
      raise MeetupDayException('Invalid')
  elif week == 'last':
    month += 1
    if month > 12:
      year += month // 12
      month %= 12
    d = datetime.date(year, month, 1) - datetime.timedelta(7)
  else:
    raise MeetupDayException('Invalid')

  day_of_week = Weekday[day_of_week].value
  dow_shift = datetime.timedelta((day_of_week - d.weekday()) % 7)
  return d + dow_shift


# vim:ts=2:sw=2:expandtab
