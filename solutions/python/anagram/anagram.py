from collections import defaultdict

def find_anagrams(word, candidates):
  matches = []

  wcount = defaultdict(int)
  for c in word.lower():
    wcount[c] += 1

  for cand in candidates:
    if cand.lower() == word.lower():
      continue

    ccount = defaultdict(int)
    for c in cand.lower():
      ccount[c] += 1
    if ccount == wcount:
      matches.append(cand)

  return matches


# vim:ts=2:sw=2:expandtab
