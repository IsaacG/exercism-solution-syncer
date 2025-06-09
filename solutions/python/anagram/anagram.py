from collections import defaultdict


def letter_count(word):
  count = defaultdict(int)
  for c in word.lower():
    count[c] += 1
  return count


def find_anagrams(word, candidates):
  wcount = letter_count(word)

  return [c for c in candidates
          if c.lower() != word.lower() and wcount == letter_count(c)]


# vim:ts=2:sw=2:expandtab
