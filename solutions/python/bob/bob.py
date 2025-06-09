import string

def response(hey):
  hey = hey.strip()

  if not hey:  # Silence
    return 'Fine. Be that way!'
  if hey.endswith('?'):
    return "Calm down, I know what I'm doing!" if hey.isupper() else 'Sure.'
  else:
    return 'Whoa, chill out!' if hey.isupper() else 'Whatever.'


# vim:ts=2:sw=2:expandtab
