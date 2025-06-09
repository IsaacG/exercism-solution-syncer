def abbreviate(words):
  words = words.replace('-', ' ').replace('_', ' ').split()
  return ''.join(w[0] for w in words).upper()

def abbreviate2(words):
  for char in '-_':
    words = words.replace(char, ' ')
  out = ''
  for word in words.split():
    out += word[0]
  return out.upper()


# vim:ts=2:sw=2:expandtab
