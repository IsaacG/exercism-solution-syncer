import string

def response(hey):
  hey = hey.strip()

  if not hey:  # Silence
    return 'Fine. Be that way!'
  if hey.endswith('?'):
    return "Calm down, I know what I'm doing!" if yelling(hey) else 'Sure.'
  else:
    return 'Whoa, chill out!' if yelling(hey) else 'Whatever.'


def yelling(hey):
  return (any(c in string.ascii_uppercase for c in hey)
      and not any(c in string.ascii_lowercase for c in hey))

    


# vim:ts=2:sw=2:expandtab
